package gateway

import (
	"ohdada/g2gserver/internal/pkg/common"
	"ohdada/g2gserver/internal/pkg/pb"

	"ohdada/g2gserver/internal/pkg/glog"

	"google.golang.org/protobuf/proto"
)

// PlayerCCHandlerSendInfo .
type PlayerCCHandlerSendInfo struct {
	player *Player
}

// Handle .
func (handler *PlayerCCHandlerSendInfo) Handle(command *common.ChannelCommand) {
	player := handler.player

	playerData, err := proto.Marshal(player)

	if err != nil {
		glog.Errorln(err, player)
		return
	}

	player.sendRequest(pb.RequestCommand_RC_PLAYER_INFO, playerData)

}
