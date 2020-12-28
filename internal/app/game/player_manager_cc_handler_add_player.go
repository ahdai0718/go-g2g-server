package game

import (
	"ohdada/g2gserver/internal/pkg/common"
	"ohdada/g2gserver/internal/pkg/core/player"
)

// PlayerManagerCCHandlerAddPlayer .
type PlayerManagerCCHandlerAddPlayer struct {
	manager *PlayerManager
}

// Handle .
func (handler *PlayerManagerCCHandlerAddPlayer) Handle(command *common.ChannelCommand) {
	manager := handler.manager
	player := command.Data[0].(player.Player)

	if player == nil {
		if command.Channel != nil {
			command.Channel <- false
		}
		return
	}

	playerExists, isExists := manager.playerMapByID[player.ID()]

	if isExists {
		delete(manager.playerMapByID, playerExists.ID())
		delete(manager.playerMapByWC, playerExists.WebsocketConnection())
		playerExists.CloseConnection()
	}

	manager.playerMapByID[player.ID()] = player
	manager.playerMapByWC[player.WebsocketConnection()] = player

	go player.Observe()

	player.HasExists(playerExists)
	player.OnConnect()

	if command.Channel != nil {
		command.Channel <- true
	}
}
