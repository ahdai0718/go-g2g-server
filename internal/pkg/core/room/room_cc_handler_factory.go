package room

import "ohdada/g2gserver/internal/pkg/common"

// CCHandlerSimpleFactory .
type CCHandlerSimpleFactory interface {
	CreateCCHandler(command *common.ChannelCommand) common.ChannelCommandHandler
}
