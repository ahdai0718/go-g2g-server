package game

import (
	"fmt"
	"net/http"
	"net/url"
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/network"
	"ohdada/g2gserver/internal/pkg/pb"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var (
	serverWebsocketConnectionManager *ServerWebsocketConnectionManager
)

func defaultServerWebsocketConnectionManager() *ServerWebsocketConnectionManager {
	if serverWebsocketConnectionManager == nil {
		serverWebsocketConnectionManager = &ServerWebsocketConnectionManager{
			intervalSendServerInfo:        time.NewTicker(time.Minute),
			intervalSyncGatewayConnection: time.NewTicker(time.Minute),
			syncOnLaunchChannel:           make(chan interface{}, 1),
			addWebsocketConnectionChannel: make(chan *network.WebsocketConnection, 32),
		}
		serverWebsocketConnectionManager.Init()
	}
	return serverWebsocketConnectionManager
}

// ServerWebsocketConnectionManager .
type ServerWebsocketConnectionManager struct {
	network.WebsocketConnectionManager
	intervalSendServerInfo        *time.Ticker
	intervalSyncGatewayConnection *time.Ticker
	syncOnLaunchChannel           chan interface{}
	addWebsocketConnectionChannel chan *network.WebsocketConnection
}

// OnConnect .
func (manager *ServerWebsocketConnectionManager) OnConnect(wc *network.WebsocketConnection) {
	glog.Infoln("ServerWebsocketConnectionManager OnConnect", wc)
	manager.addWebsocketConnectionChannel <- wc
}

func (manager *ServerWebsocketConnectionManager) observeReceiveRequest(wc *network.WebsocketConnection) {
	for {
		select {
		case request := <-wc.ReceiveRequestChannel():
			command := pb.RequestCommand(request.Command)
			handler := defaultServerWSMessaageHandlerSimpleFactory().create(command)
			handler.Handle(request.Data)

		case <-wc.CloseChannel():
			manager.RemoveWebsocketConnection(wc)
		}
	}
}

func (manager *ServerWebsocketConnectionManager) observe() {
	for {
		select {
		case <-manager.intervalSendServerInfo.C:
			manager.sendServerInfo()

		case <-manager.intervalSyncGatewayConnection.C:

		case wc := <-manager.addWebsocketConnectionChannel:
			if !manager.IsWebsocketConnectionExists(wc) {
				manager.AddWebsocketConnection(wc)
				go manager.observeReceiveRequest(wc)
				manager.sendServerInfo()
			}

		case <-manager.syncOnLaunchChannel:

		}
	}
}

func (manager *ServerWebsocketConnectionManager) sendServerInfo() {

	data, err := proto.Marshal(ServerInfo())

	if err != nil {
		glog.Error(err)
	} else {
		for _, wc := range manager.WebsocketConnectionMap() {
			wc.SendRequest(int32(pb.RequestCommand_RC_SERVER_INFO), data)
		}
	}
}

func (manager *ServerWebsocketConnectionManager) sendRequest(command int32, data []byte) {
	for _, wc := range manager.WebsocketConnectionMap() {
		wc.SendRequest(command, data)
	}
}

func (manager *ServerWebsocketConnectionManager) syncOnLaunch() {
	manager.syncOnLaunchChannel <- true
}

func (manager *ServerWebsocketConnectionManager) connectToGeteway(serverInfoGateway *pb.ServerInfo) {

	if _, isExists := manager.WebsocketConnectionMap()[serverInfoGateway.Name]; isExists {
		return
	}

	url := url.URL{
		Scheme: serverInfoGateway.WebsocketProtocol,
		Host:   fmt.Sprintf("%s:%d", serverInfoGateway.Host, serverInfoGateway.Port),
		Path:   serverInfoGateway.WebsocketRoutePath,
	}

	header := http.Header{
		"CC-Server-Name": []string{ServerInfo().Name},
	}

	websocketConnection, _, err := websocket.DefaultDialer.Dial(url.String(), header)

	if err != nil {
		glog.Error(err)
	} else {
		wc := network.NewWebsocketConnectionBuilder().
			SetID(serverInfoGateway.Name).
			SetConnection(websocketConnection).
			Build()

		defaultServerWebsocketConnectionEventHandler.OnConnect(wc)

		go wc.ObserveReadMessage()
		go wc.ObserveWriteMessage()
	}
}
