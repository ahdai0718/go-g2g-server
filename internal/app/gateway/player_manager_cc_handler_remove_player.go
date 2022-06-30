package gateway

import (
	"ohdada/g2gserver/internal/pkg/common"
	pkgPlayer "ohdada/g2gserver/internal/pkg/core/player"
)

// PlayerManagerCCHandlerRemovePlayer .
type PlayerManagerCCHandlerRemovePlayer struct {
	manager *PlayerManager
}

// Handle .
func (handler *PlayerManagerCCHandlerRemovePlayer) Handle(command *common.ChannelCommand) {
	manager := handler.manager

	player := command.Data[0].(pkgPlayer.Player)

	if player == nil {
		return
	}

	if value, isExists := manager.playerMapByWC.Load(player.WebsocketConnection()); isExists {
		playerExists := value.(pkgPlayer.Player)
		manager.playerMapByWC.Delete(playerExists.WebsocketConnection())
		playerExists.CloseConnection()
	}

	if value, isExists := manager.playerMapByID.Load(player.ID()); isExists {
		playerExists := value.(pkgPlayer.Player)
		if player == playerExists {
			manager.playerMapByID.Delete(playerExists.ID())
		}
	}

	player.CloseRoutine()
}
