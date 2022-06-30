package gateway

import (
	"ohdada/g2gserver/internal/pkg/common"
	"ohdada/g2gserver/internal/pkg/core/player"
	"sync"
	"time"
)

var (
	defaultPlayerManager = &PlayerManager{
		playerMapByID:  &sync.Map{},
		playerMapByWC:  &sync.Map{},
		commandChannel: make(chan *common.ChannelCommand, maxPlayer),
		ticker:         time.NewTicker(time.Second),
	}
)

// DefaultPlayerManager .
func DefaultPlayerManager() *PlayerManager {
	return defaultPlayerManager
}

// PlayerManager .
type PlayerManager struct {
	playerMapByID  *sync.Map
	playerMapByWC  *sync.Map
	commandChannel chan *common.ChannelCommand
	ticker         *time.Ticker
}

func (manager *PlayerManager) init() error {
	return nil
}

func (manager *PlayerManager) observe() {
	for {
		select {
		case command := <-manager.commandChannel:
			manager.handleChannelCommand(command)

		case <-manager.ticker.C:

		}
	}
}

func (manager *PlayerManager) addChannelCommand(command *common.ChannelCommand) {
	manager.commandChannel <- command
}

func (manager *PlayerManager) handleChannelCommand(command *common.ChannelCommand) {
	handler := manager.createCCHandler(command)
	if handler != nil {
		handler.Handle(command)
	}
}

func (manager *PlayerManager) addPlayerInAsync(player player.Player) {
	manager.addChannelCommand(common.NewChannelCommand(
		PlayerManagerCCAddPlayer,
		player,
	))
}

func (manager *PlayerManager) removePlayerInAsync(player player.Player) {
	manager.addChannelCommand(common.NewChannelCommand(
		PlayerManagerCCRemovePlayer,
		player,
	))
}

func (manager *PlayerManager) broadcastInAsync(data []byte) {
	manager.addChannelCommand(common.NewChannelCommand(
		PlayerManagerCCBroadcast,
		data,
	))
}
