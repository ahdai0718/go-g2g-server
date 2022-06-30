package network

// WebsocketConnectionMessageHandler .
type WebsocketConnectionMessageHandler interface {
	Handle(message []byte)
}
