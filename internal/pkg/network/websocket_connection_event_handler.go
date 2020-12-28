package network

import (
	"ohdada/g2gserver/internal/pkg/glog"
)

// WebsocketConnectionEventHandler .
type WebsocketConnectionEventHandler struct {
	connectChannel chan *WebsocketConnection
	callback       WebsocketConnectionEventCallback
}

// NewWebsocketConnectionEventHandler .
func NewWebsocketConnectionEventHandler(maxWebsocketConnection int, callback WebsocketConnectionEventCallback) *WebsocketConnectionEventHandler {
	return &WebsocketConnectionEventHandler{
		connectChannel: make(chan *WebsocketConnection, maxWebsocketConnection),
		callback:       callback,
	}
}

// Observe .
func (handler *WebsocketConnectionEventHandler) Observe() {
	for {
		select {
		case wc := <-handler.connectChannel:
			glog.Info(wc)
			handler.callback.OnConnect(wc)
		}
	}
}

// OnConnect .
func (handler *WebsocketConnectionEventHandler) OnConnect(websocketConnection *WebsocketConnection) {
	if websocketConnection != nil {
		handler.connectChannel <- websocketConnection
	} else {
		glog.Warningln("WebsocketConnection is <nil>...")
	}
}
