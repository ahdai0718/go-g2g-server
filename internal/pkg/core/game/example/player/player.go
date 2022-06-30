package player

import (
	"ohdada/g2gserver/internal/pkg/core/game/example/abstract"
	"ohdada/g2gserver/internal/pkg/core/game/example/constant"
	cchandler "ohdada/g2gserver/internal/pkg/core/game/example/player/cc_handler"
	pkgPlayer "ohdada/g2gserver/internal/pkg/core/player"
	"time"

	"github.com/jinzhu/copier"
)

// NewPlayer .
func NewPlayer(config *Config) *Player {

	player := &Player{
		closeRoutineChannel: make(chan bool, 1),
		isConnected:         config.IsConnected,
		ticker:              time.NewTicker(time.Second),
	}

	playerBase := pkgPlayer.NewBuilder().
		SetWebsocketConnection(config.WC).
		SetCCHandlerBufferSize(constant.PlayerMaxCCHandlerBufferSize).
		SetCCHandlerSimpleFactory(&cchandler.SimpleFactory{}).
		SetRequestHandlerSimpleFactory(player).
		Build()

	player.Base = playerBase

	platformPlayer := player.PlatformProvider().PlatformPlayer()

	copier.Copy(player, platformPlayer)

	player.Id = platformPlayer.IdAtPlatform

	return player
}

// Player .
type Player struct {
	lastRequestRoomJoin      time.Time
	lastRequestInfo          time.Time
	lastRequestRoomBroadcast time.Time
	room                     abstract.Room
	ticker                   *time.Ticker
	closeRoutineChannel      chan bool
	*pkgPlayer.Base
	creditToBet             int64
	continuousTickNoCommand int64
	isConnected             bool
	isReconnecting          bool
}
