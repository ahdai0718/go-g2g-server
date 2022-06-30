package network

// WebsocketConnectionEventCallback .
type WebsocketConnectionEventCallback interface {
	OnConnect(wc *WebsocketConnection)
}
