package game

import (
	"ohdada/g2gserver/internal/pkg/common"
	"ohdada/g2gserver/internal/pkg/core/player"
)

// PlayerManagerCCHandlerRemovePlayer .
type PlayerManagerCCHandlerRemovePlayer struct {
	manager *PlayerManager
}

// Handle .
func (handler *PlayerManagerCCHandlerRemovePlayer) Handle(command *common.ChannelCommand) {
	manager := handler.manager

	player := command.Data[0].(player.Player)

	if player == nil {
		return
	}

	if playerExists, isExists := manager.playerMapByWC[player.WebsocketConnection()]; isExists {
		delete(manager.playerMapByWC, playerExists.WebsocketConnection())
		playerExists.CloseConnection()
	}

	if playerExists, isExists := manager.playerMapByID[player.ID()]; isExists {
		if player == playerExists {
			delete(manager.playerMapByID, playerExists.ID())
		}
	}
}
