package gateway

import (
	"ohdada/g2gserver/internal/pkg/common"
)

// PlayerCCHandlerCloseRoutine .
type PlayerCCHandlerCloseRoutine struct {
	player *Player
}

// Handle .
func (handler *PlayerCCHandlerCloseRoutine) Handle(command *common.ChannelCommand) {
	player := handler.player

	player.isReceiveCloseRoutine = true

}
