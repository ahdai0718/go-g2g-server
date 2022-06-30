package player

import "ohdada/g2gserver/internal/pkg/network"

// Config .
type Config struct {
	WC          *network.WebsocketConnection
	IsConnected bool
}
