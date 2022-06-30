package gateway

import (
	"ohdada/g2gserver/internal/pkg/common"
	"ohdada/g2gserver/internal/pkg/core/player"
	pkgPlayer "ohdada/g2gserver/internal/pkg/core/player"
	"ohdada/g2gserver/internal/pkg/core/serverinfo"
	"ohdada/g2gserver/internal/pkg/network"
	"ohdada/g2gserver/internal/pkg/pb"
	"time"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/proto"
)

const (
	maxPlayerCCHandlerBufferSize = 128
)

func newPlayer(wc *network.WebsocketConnection) *Player {

	player := &Player{
		closeRoutineChannel:   make(chan bool, 1),
		isReceiveCloseRoutine: false,
		ticker:                time.NewTicker(time.Second),
	}

	playerBase := pkgPlayer.NewBuilder().
		SetWebsocketConnection(wc).
		SetCCHandlerBufferSize(maxPlayerCCHandlerBufferSize).
		SetCCHandlerSimpleFactory(player).
		Build()

	player.Base = playerBase

	platformPlayer := player.PlatformProvider().PlatformPlayer()

	copier.Copy(player, platformPlayer)

	player.Id = platformPlayer.IdAtPlatform

	return player
}

// Player .
type Player struct {
	*player.Base
	closeRoutineChannel   chan bool
	ticker                *time.Ticker
	isReceiveCloseRoutine bool
}

// Observe .
func (player *Player) Observe() {

	for {
		select {
		case command := <-player.CommandChannel():
			player.HandleChannelCommand(command)

		case request := <-player.ReceiveRequestChannel():
			player.HandleRequest(request)

		case <-player.ConnectionCloseChannel():
			DefaultPlayerManager().removePlayerInAsync(player)

		case <-player.ticker.C:
			player.checkToCloseRoutine()

		case <-player.closeRoutineChannel:
			player.ticker.Stop()
			return
		}
	}
}

// HasExists .
func (player *Player) HasExists(playerExists player.Player) {

}

// OnConnect .
func (player *Player) OnConnect() {
	player.sendPlayerInfoInAsync()
	player.sendRunningServerInfoMap()
	player.sendServerInfo()
}

func (player *Player) sendRequest(command pb.RequestCommand, data []byte) {
	player.WebsocketConnection().SendRequest(int32(command), data)
}

// CloseRoutine .
func (player *Player) CloseRoutine() {
	player.AddChannelCommand(common.NewChannelCommand(
		PlayerCCCloseRoutine,
	))
}

func (player *Player) checkToCloseRoutine() {
	if player.isReceiveCloseRoutine {
		if len(player.CommandChannel()) == 0 {
			player.closeRoutineChannel <- true
		}
	}
}

func (player *Player) sendPlayerInfoInAsync() {
	player.AddChannelCommand(common.NewChannelCommand(
		PlayerCCSendInfo,
		player,
	))
}

func (player *Player) sendServerInfo() {
	data, _ := proto.Marshal(ServerInfo())
	player.sendRequest(pb.RequestCommand_RC_SERVER_INFO, data)
}

func (player *Player) sendRunningServerInfoMap() {
	data, _ := proto.Marshal(serverinfo.DefaultManager().CurrentRunningServerMap())
	player.sendRequest(pb.RequestCommand_RC_SERVER_INFO_MAP_BY_NAME, data)
}
