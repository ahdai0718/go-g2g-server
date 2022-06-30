package gateway

import (
	"ohdada/g2gserver/internal/pkg/common"
)

// PlayerChannelCommand
const (
	_ = iota
	PlayerCCSendInfo
	PlayerCCCloseRoutine
)

// Create .
func (player *Player) Create(command *common.ChannelCommand, subjectList ...interface{}) (handler common.ChannelCommandHandler) {

	switch command.Command {
	case PlayerCCSendInfo:
		handler = &PlayerCCHandlerSendInfo{player: player}

	case PlayerCCCloseRoutine:
		handler = &PlayerCCHandlerCloseRoutine{player: player}
	}

	return handler
}
