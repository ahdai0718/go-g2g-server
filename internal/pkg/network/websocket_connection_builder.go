package network

import (
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/platform"
	"time"

	"github.com/gorilla/websocket"
)

// NewWebsocketConnectionBuilder .
func NewWebsocketConnectionBuilder() *WebsocketConnectionBuilder {
	builder := &WebsocketConnectionBuilder{}
	builder.messageChannelSize = 512
	return builder
}

// WebsocketConnectionBuilder .
type WebsocketConnectionBuilder struct {
	platformProvider platform.Provider
	WebsocketConnection
	messageChannelSize int
}

// SetID .
func (builder *WebsocketConnectionBuilder) SetID(id string) *WebsocketConnectionBuilder {
	builder.id = id
	return builder
}

// SetConnection .
func (builder *WebsocketConnectionBuilder) SetConnection(connection *websocket.Conn) *WebsocketConnectionBuilder {
	builder.connection = connection
	return builder
}

// SetMessageChannelSize .
func (builder *WebsocketConnectionBuilder) SetMessageChannelSize(messageChannelSize int) *WebsocketConnectionBuilder {
	builder.messageChannelSize = messageChannelSize
	return builder
}

// SetReadTimeoutDuration .
func (builder *WebsocketConnectionBuilder) SetReadTimeoutDuration(readTimeoutDuration time.Duration) *WebsocketConnectionBuilder {
	builder.readTimeoutDuration = readTimeoutDuration
	return builder
}

// SetWriteTimeoutDuration .
func (builder *WebsocketConnectionBuilder) SetWriteTimeoutDuration(writeTimeoutDuration time.Duration) *WebsocketConnectionBuilder {
	builder.writeTimeoutDuration = writeTimeoutDuration
	return builder
}

// SetPlatformProvider .
func (builder *WebsocketConnectionBuilder) SetPlatformProvider(platformProvider platform.Provider) *WebsocketConnectionBuilder {
	builder.platformProvider = platformProvider
	return builder
}

// Build .
func (builder *WebsocketConnectionBuilder) Build() *WebsocketConnection {

	websocketConnection := &WebsocketConnection{
		id:                      builder.id,
		connection:              builder.connection,
		receiveRequestChannel:   make(chan *pb.Request, builder.messageChannelSize),
		sendMessageChannel:      make(chan []byte, builder.messageChannelSize),
		closeChannel:            make(chan bool, 1),
		closeSendRoutineChannel: make(chan bool, 1),
		isReceiveClose:          false,
		receiveCloseChannel:     make(chan bool, 1),
		readTimeoutDuration:     builder.readTimeoutDuration,
		writeTimeoutDuration:    builder.writeTimeoutDuration,
		platformProvider:        builder.platformProvider,
	}

	return websocketConnection
}
