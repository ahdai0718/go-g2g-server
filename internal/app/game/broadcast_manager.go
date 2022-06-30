package game

import (
	"ohdada/g2gserver/internal/pkg/pb"
)

var (
	defaultBroadcastManager = &BroadcastManager{
		dataChannel:          make(chan []byte, maxPlayer),
		dataToGatewayChannel: make(chan []byte, maxPlayer),
	}
)

// DefaultBroadcastManager .
func DefaultBroadcastManager() *BroadcastManager {
	return defaultBroadcastManager
}

// BroadcastManager .
type BroadcastManager struct {
	dataChannel          chan []byte
	dataToGatewayChannel chan []byte
}

// RequestChannel .
func (manager *BroadcastManager) RequestChannel() chan []byte {
	return manager.dataChannel
}

// RequestToGatewayChannel .
func (manager *BroadcastManager) RequestToGatewayChannel() chan []byte {
	return manager.dataToGatewayChannel
}

func (manager *BroadcastManager) init() error {
	return nil
}

func (manager *BroadcastManager) observe() {
	for {
		select {
		case data := <-manager.dataChannel:
			manager.broadcast(data)

		case data := <-manager.dataToGatewayChannel:
			manager.sendBroadcastRequestToGateway(data)
		}
	}
}

func (manager *BroadcastManager) sendBroadcastRequestToGateway(data []byte) {
	defaultServerWebsocketConnectionManager().sendRequest(int32(pb.RequestCommand_RC_SERVER_BROADCAST_GATEWAY), data)
}

func (manager *BroadcastManager) broadcast(data []byte) {
	DefaultPlayerManager().broadcastInAsync(data)
}
