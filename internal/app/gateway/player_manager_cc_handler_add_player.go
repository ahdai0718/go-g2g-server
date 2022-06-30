package gateway

import (
	"ohdada/g2gserver/internal/pkg/common"
	pkgPlayer "ohdada/g2gserver/internal/pkg/core/player"
)

// PlayerManagerCCHandlerAddPlayer .
type PlayerManagerCCHandlerAddPlayer struct {
	manager *PlayerManager
}

// Handle .
func (handler *PlayerManagerCCHandlerAddPlayer) Handle(command *common.ChannelCommand) {
	manager := handler.manager
	player := command.Data[0].(pkgPlayer.Player)

	if player == nil {
		if command.Channel != nil {
			command.Channel <- false
		}
		return
	}

	value, isExists := manager.playerMapByID.Load(player.ID())

	if isExists {
		playerExists := value.(pkgPlayer.Player)
		manager.playerMapByID.Delete(playerExists.ID())
		manager.playerMapByWC.Delete(playerExists.WebsocketConnection())
		playerExists.CloseConnection()
	}

	manager.playerMapByID.Store(player.ID(), player)
	manager.playerMapByWC.Store(player.WebsocketConnection(), player)

	go player.Observe()

	if isExists {
		playerExists := value.(pkgPlayer.Player)
		player.HasExists(playerExists)
	}

	player.OnConnect()

	if command.Channel != nil {
		command.Channel <- true
	}
}
