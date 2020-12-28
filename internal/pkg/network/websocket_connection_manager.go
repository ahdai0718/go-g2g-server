package network

// WebsocketConnectionManager .
type WebsocketConnectionManager struct {
	websocketConnectionMap map[string]*WebsocketConnection
}

// Init .
func (manager *WebsocketConnectionManager) Init() {
	manager.websocketConnectionMap = make(map[string]*WebsocketConnection)
}

// IsWebsocketConnectionExists .
func (manager *WebsocketConnectionManager) IsWebsocketConnectionExists(wc *WebsocketConnection) bool {

	_, isExists := manager.websocketConnectionMap[wc.ID()]

	return isExists
}

// AddWebsocketConnection .
func (manager *WebsocketConnectionManager) AddWebsocketConnection(wc *WebsocketConnection) {
	manager.websocketConnectionMap[wc.ID()] = wc
}

// RemoveWebsocketConnection .
func (manager *WebsocketConnectionManager) RemoveWebsocketConnection(wc *WebsocketConnection) {
	delete(manager.websocketConnectionMap, wc.ID())
}

// WebsocketConnectionMap .
func (manager *WebsocketConnectionManager) WebsocketConnectionMap() map[string]*WebsocketConnection {
	return manager.websocketConnectionMap
}
