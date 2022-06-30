package gateway

import (
	"net/http"
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/network"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	serverUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// ServerWebsocketHandler .
func ServerWebsocketHandler(context *gin.Context) {

	serverName := context.Request.Header.Get("CC-Server-Name")

	websocketConnection, connectionError := serverUpgrader.Upgrade(context.Writer, context.Request, nil)

	if _, ok := connectionError.(websocket.HandshakeError); ok {
		glog.Error(connectionError)
		http.Error(context.Writer, "Not a websocket handshake", http.StatusBadRequest)
		return
	} else if connectionError != nil {
		glog.Error(connectionError)
		http.Error(context.Writer, "Websocket connect error.", http.StatusInternalServerError)
		return
	}

	wc := network.NewWebsocketConnectionBuilder().
		SetID(serverName).
		SetConnection(websocketConnection).
		Build()

	// defaultServerWebsocketConnectionEventHandler.OnConnect(wc)

	go wc.ObserveReadMessage()
	go wc.ObserveWriteMessage()
}
