package gateway

import (
	"ohdada/g2gserver/internal/pkg/network"
	"ohdada/g2gserver/internal/pkg/pb"
)

var (
	serverWSMessaageHandlerSimpleFactory *ServerWSMessaageHandlerSimpleFactory
)

func defaultServerWSMessaageHandlerSimpleFactory() *ServerWSMessaageHandlerSimpleFactory {
	if serverWSMessaageHandlerSimpleFactory == nil {
		serverWSMessaageHandlerSimpleFactory = &ServerWSMessaageHandlerSimpleFactory{}
	}
	return serverWSMessaageHandlerSimpleFactory
}

// ServerWSMessaageHandlerSimpleFactory .
type ServerWSMessaageHandlerSimpleFactory struct{}

func (factory *ServerWSMessaageHandlerSimpleFactory) create(command pb.RequestCommand) (handler network.WebsocketConnectionMessageHandler) {

	switch command {
	case pb.RequestCommand_RC_SERVER_INFO:
		handler = &ServerWSMessageHandlerServerInfo{}

	case pb.RequestCommand_RC_SERVER_BROADCAST_GATEWAY:
		handler = &ServerWSMessageHandlerBroadcast{}

	}

	return
}
