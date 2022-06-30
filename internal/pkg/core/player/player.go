package player

import (
	"ohdada/g2gserver/internal/pkg/common"
	"ohdada/g2gserver/internal/pkg/network"
)

// Player .
type Player interface {
	ID() string
	WebsocketConnection() *network.WebsocketConnection
	Observe()
	HandleChannelCommand(*common.ChannelCommand)
	AddChannelCommand(*common.ChannelCommand)
	CloseConnection()
	CloseRoutine()
	HasExists(Player)
	OnConnect()
}
