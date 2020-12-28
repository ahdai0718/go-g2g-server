package game

import "ohdada/g2gserver/internal/pkg/common"

// PlayerManagerChannelCommand
const (
	_ = iota
	PlayerManagerCCAddPlayer
	PlayerManagerCCRemovePlayer
	PlayerManagerCCBroadcast
)

func (manager *PlayerManager) createCCHandler(command *common.ChannelCommand) (handler common.ChannelCommandHandler) {

	switch command.Command {
	case PlayerManagerCCAddPlayer:
		handler = &PlayerManagerCCHandlerAddPlayer{manager: manager}

	case PlayerManagerCCRemovePlayer:
		handler = &PlayerManagerCCHandlerRemovePlayer{manager: manager}

	case PlayerManagerCCBroadcast:
		handler = &PlayerManagerCCHandlerBroadcast{manager: manager}
	}

	return
}
