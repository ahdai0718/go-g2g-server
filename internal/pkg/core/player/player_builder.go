package player

import (
	"ohdada/g2gserver/internal/pkg/common"
	"ohdada/g2gserver/internal/pkg/network"
	"time"
)

// NewBuilder .
func NewBuilder() *Builder {
	return &Builder{}
}

// Builder .
type Builder struct {
	wc                          *network.WebsocketConnection
	ccHandlerBufferSize         int
	ccHandlerSimpleFactory      common.ChannelCommandHandlerFactory
	requestHandlerSimpleFactory RequestHandlerSimpleFactory
}

// SetWebsocketConnection .
func (builder *Builder) SetWebsocketConnection(wc *network.WebsocketConnection) *Builder {
	builder.wc = wc
	return builder
}

// SetCCHandlerBufferSize .
func (builder *Builder) SetCCHandlerBufferSize(ccHandlerBufferSize int) *Builder {
	builder.ccHandlerBufferSize = ccHandlerBufferSize
	return builder
}

// SetCCHandlerSimpleFactory .
func (builder *Builder) SetCCHandlerSimpleFactory(ccHandlerSimpleFactory common.ChannelCommandHandlerFactory) *Builder {
	builder.ccHandlerSimpleFactory = ccHandlerSimpleFactory
	return builder
}

// SetRequestHandlerSimpleFactory .
func (builder *Builder) SetRequestHandlerSimpleFactory(requestHandlerSimpleFactory RequestHandlerSimpleFactory) *Builder {
	builder.requestHandlerSimpleFactory = requestHandlerSimpleFactory
	return builder
}

// Build .
func (builder *Builder) Build() *Base {

	base := &Base{
		wc:                          builder.wc,
		commandChannel:              make(chan *common.ChannelCommand, builder.ccHandlerBufferSize),
		lastTimeReceiveMessage:      time.Now(),
		ticker:                      time.NewTicker(time.Second),
		isReceiveClose:              false,
		cchandlerSimpleFactory:      builder.ccHandlerSimpleFactory,
		requestHandlerSimpleFactory: builder.requestHandlerSimpleFactory,
	}

	return base
}
