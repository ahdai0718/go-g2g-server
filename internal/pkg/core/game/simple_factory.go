package game

import (
	"ohdada/g2gserver/internal/pkg/core/game/example"
	"ohdada/g2gserver/internal/pkg/core/player"
	"ohdada/g2gserver/internal/pkg/network"
	"ohdada/g2gserver/internal/pkg/pb"

	pbGame "ohdada/g2gserver/internal/pkg/pb/game"
)

var (
	defaultSimpleFactory *SimpleFactory
)

// DefaultSimpleFactory .
func DefaultSimpleFactory() *SimpleFactory {
	if defaultSimpleFactory == nil {
		defaultSimpleFactory = &SimpleFactory{}
	}
	return defaultSimpleFactory
}

// Game .
type Game interface {
	Init() (err error)
	Run() (err error)
	SetRunMode(string)
	SetStoreConnection(*pb.StoreConnection)
	SetPlayerChannelRemove(chan player.Player)
	SetServerBroadcastChannel(chan []byte)
	SetGatewayBroadcastChannel(chan []byte)
	SetLogServerInfo([]*pb.ServerInfo)
	IsUnderMaintain() bool
	NewPlayer(*network.WebsocketConnection) player.Player
}

// SimpleFactory .
type SimpleFactory struct{}

// Create .
func (factory *SimpleFactory) Create(gameType pbGame.GameType) (game Game) {

	switch gameType {
	default:
		game = example.NewGame()
	}

	return
}
