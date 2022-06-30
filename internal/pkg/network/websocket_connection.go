package network

import (
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/platform"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

// WebsocketMessage .
type WebsocketMessage struct {
	WebsocketConnection *WebsocketConnection
	Request             *pb.Request
}

// WebsocketConnection .
type WebsocketConnection struct {
	platformProvider        platform.Provider
	connection              *websocket.Conn
	receiveRequestChannel   chan *pb.Request
	closeChannel            chan bool
	closeSendRoutineChannel chan bool
	receiveCloseChannel     chan bool
	sendMessageChannel      chan []byte
	id                      string
	readTimeoutDuration     time.Duration
	writeTimeoutDuration    time.Duration
	isReceiveClose          bool
}

// ID .
func (wc *WebsocketConnection) ID() string {
	return wc.id
}

// PlatformProvider .
func (wc *WebsocketConnection) PlatformProvider() platform.Provider {
	return wc.platformProvider
}

// SetPlatformProvider .
func (wc *WebsocketConnection) SetPlatformProvider(platformProvider platform.Provider) {
	wc.platformProvider = platformProvider
}

// ReceiveRequestChannel .
func (wc *WebsocketConnection) ReceiveRequestChannel() chan *pb.Request {
	return wc.receiveRequestChannel
}

// CloseChannel .
func (wc *WebsocketConnection) CloseChannel() chan bool {
	return wc.closeChannel
}

// ObserveReadMessage .
func (wc *WebsocketConnection) ObserveReadMessage() {
	for {
		if wc.readTimeoutDuration > 0 {
			wc.connection.SetReadDeadline(time.Now().Add(wc.readTimeoutDuration))
		}
		_, message, readMessageError := wc.connection.ReadMessage()

		if readMessageError != nil {

			if websocket.IsUnexpectedCloseError(readMessageError, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				glog.Errorln("WebsocketConnection read message error", readMessageError)
			}

			if websocket.IsCloseError(readMessageError, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				glog.Warningln("WebsocketConnection close connection", readMessageError)
			}

			wc.Close()
			wc.closeChannel <- true
			return
		}

		request := &pb.Request{}
		err := proto.Unmarshal(message, request)
		if err == nil {
			wc.receiveRequestChannel <- request
		}
	}
}

// ObserveWriteMessage .
func (wc *WebsocketConnection) ObserveWriteMessage() {
	for {
		select {
		case message := <-wc.sendMessageChannel:

			if wc.writeTimeoutDuration > 0 {
				wc.connection.SetWriteDeadline(time.Now().Add(wc.writeTimeoutDuration))
			}
			writeMessageError := wc.connection.WriteMessage(websocket.BinaryMessage, message)

			if writeMessageError != nil {
				glog.Error(writeMessageError)
				return
			}

		case <-wc.receiveCloseChannel:
			if !wc.isReceiveClose {
				wc.isReceiveClose = true
				wc.closeSendRoutineChannel <- true
			}

		case <-wc.closeSendRoutineChannel:
			wc.connection.Close()
			return
		}
	}
}

// SendMessage .
func (wc *WebsocketConnection) SendMessage(message []byte) {
	if wc.sendMessageChannel != nil {
		wc.sendMessageChannel <- message
	}
}

// SendRequest .
func (wc *WebsocketConnection) SendRequest(command int32, data []byte) {

	request := &pb.Request{
		Command:   command,
		Data:      data,
		Timestamp: time.Now().Unix(),
		Error: &pb.Error{
			Code: int32(pb.ErrorCode_EC_NONE),
		},
	}

	message, err := proto.Marshal(request)

	if err != nil {
		glog.Error(err)
		return
	}

	wc.SendMessage(message)
}

// SendRequestError .
func (wc *WebsocketConnection) SendRequestError(command int32, data []byte, pbError *pb.Error) {

	request := &pb.Request{
		Command:   command,
		Data:      data,
		Timestamp: time.Now().Unix(),
		Error:     pbError,
	}

	message, err := proto.Marshal(request)

	if err != nil {
		glog.Error(err)
		return
	}

	wc.SendMessage(message)
}

// SendError .
func (wc *WebsocketConnection) SendError(code int32, action pb.ErrorAction) {
	errorMessage := &pb.Error{
		Code:   code,
		Action: action,
	}

	errorData, err := proto.Marshal(errorMessage)

	if err != nil {
		glog.Error(err)
		return
	}

	errorRequest := &pb.Request{
		Command:   int32(pb.RequestCommand_RC_ERROR),
		Data:      errorData,
		Timestamp: time.Now().Unix(),
	}

	message, err := proto.Marshal(errorRequest)

	if err != nil {
		glog.Error(err)
		return
	}

	wc.SendMessage(message)
}

// Close .
func (wc *WebsocketConnection) Close() {
	go func() {
		time.Sleep(time.Second)
		wc.receiveCloseChannel <- true
	}()
}
