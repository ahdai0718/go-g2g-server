package player

import (
	"ohdada/g2gserver/internal/pkg/common"
	"ohdada/g2gserver/internal/pkg/network"
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/platform"
	"time"
)

// Base .
type Base struct {
	lastTimeReceiveMessage      time.Time
	cchandlerSimpleFactory      common.ChannelCommandHandlerFactory
	requestHandlerSimpleFactory RequestHandlerSimpleFactory
	commandChannel              chan *common.ChannelCommand
	ticker                      *time.Ticker
	wc                          *network.WebsocketConnection
	pb.Player
	isReceiveClose bool
}

// ID .
func (base *Base) ID() string {
	return base.Player.Id
}

// DisplayName .
func (base *Base) DisplayName() string {
	return base.Player.DisplayName
}

// CommandChannel .
func (base *Base) CommandChannel() chan *common.ChannelCommand {
	return base.commandChannel
}

// WebsocketConnection .
func (base *Base) WebsocketConnection() *network.WebsocketConnection {
	return base.wc
}

// PlatformProvider ...
func (base *Base) PlatformProvider() platform.Provider {
	return base.wc.PlatformProvider()
}

// ReceiveRequestChannel .
func (base *Base) ReceiveRequestChannel() chan *pb.Request {
	return base.wc.ReceiveRequestChannel()
}

// ConnectionCloseChannel .
func (base *Base) ConnectionCloseChannel() chan bool {
	return base.wc.CloseChannel()
}

// CChandlerSimpleFactory .
func (base *Base) CChandlerSimpleFactory() common.ChannelCommandHandlerFactory {
	return base.cchandlerSimpleFactory
}

// HandleChannelCommand .
func (base *Base) HandleChannelCommand(command *common.ChannelCommand) {
	handler := base.cchandlerSimpleFactory.Create(command, base)
	if handler != nil {
		handler.Handle(command)
	}
}

// HandleRequest .
func (base *Base) HandleRequest(request *pb.Request) {
	handler := base.requestHandlerSimpleFactory.CreateRequestHandler(request)
	if handler != nil {
		handler.Handle(request)
	}
}

// AddChannelCommand .
func (base *Base) AddChannelCommand(command *common.ChannelCommand) {
	base.commandChannel <- command
}

// CloseConnection .
func (base *Base) CloseConnection() {
	base.wc.Close()
}
