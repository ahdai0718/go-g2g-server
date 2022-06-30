package example

import (
	"ohdada/g2gserver/internal/pkg/core/game/example/player"
	pkgPlayer "ohdada/g2gserver/internal/pkg/core/player"
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/network"
	"ohdada/g2gserver/internal/pkg/pb"
)

// NewGame .
func NewGame() *Game {
	return &Game{}
}

// Game .
type Game struct {
}

// SetRunMode .
func (game *Game) SetRunMode(runMode string) {

}

// SetStoreConnection .
func (game *Game) SetStoreConnection(storeConnection *pb.StoreConnection) {

}

// SetPlayerChannelRemove .
func (game *Game) SetPlayerChannelRemove(playerChannelRemove chan pkgPlayer.Player) {

}

// SetServerBroadcastChannel .
func (game *Game) SetServerBroadcastChannel(serverBroadcastChannel chan []byte) {

}

// SetGatewayBroadcastChannel .
func (game *Game) SetGatewayBroadcastChannel(chan []byte) {}

// IsUnderMaintain .
func (game *Game) IsUnderMaintain() bool {
	return false
}

// Init .
func (game *Game) Init() (err error) {

	glog.Info("game example init")

	return
}

// SetLogServerInfo .
func (game *Game) SetLogServerInfo(logServerInfoList []*pb.ServerInfo) {

}

// Run .
func (game *Game) Run() (err error) {

	return
}

// NewPlayer .
func (game *Game) NewPlayer(wc *network.WebsocketConnection) pkgPlayer.Player {
	return player.NewPlayer(&player.Config{
		WC:          wc,
		IsConnected: true,
	})
}
