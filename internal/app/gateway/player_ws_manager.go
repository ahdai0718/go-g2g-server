package gateway

import (
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/network"
)

var (
	playerWebsocketConnectionManager *PlayerWebsocketConnectionManager
)

func defaultPlayerWebsocketConnectionManager() *PlayerWebsocketConnectionManager {
	if playerWebsocketConnectionManager == nil {
		playerWebsocketConnectionManager = &PlayerWebsocketConnectionManager{}
		playerWebsocketConnectionManager.Init()
	}
	return playerWebsocketConnectionManager
}

// PlayerWebsocketConnectionManager .
type PlayerWebsocketConnectionManager struct {
	network.WebsocketConnectionManager
}

// OnConnect .
func (manager *PlayerWebsocketConnectionManager) OnConnect(wc *network.WebsocketConnection) {
	glog.Infoln("PlayerWebsocketConnectionManager OnConnect", wc)

	player := newPlayer(wc)

	DefaultPlayerManager().addPlayerInAsync(player)
}
