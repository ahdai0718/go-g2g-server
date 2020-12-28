package gateway

import (
	"ohdada/g2gserver/internal/pkg/common"
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/pb"
)

// PlayerManagerCCHandlerBroadcast .
type PlayerManagerCCHandlerBroadcast struct {
	manager *PlayerManager
}

// Handle .
func (handler *PlayerManagerCCHandlerBroadcast) Handle(command *common.ChannelCommand) {
	manager := handler.manager
	data := command.Data[0].([]byte)

	if data == nil {
		glog.Warning("Return because of nil")
		return
	}

	for _, player := range manager.playerMapByID {
		player.WebsocketConnection().SendRequest(int32(pb.RequestCommand_RC_SERVER_BROADCAST_GATEWAY), data)
	}
}
