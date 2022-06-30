package game

import (
	"fmt"
	"net/http"
	"ohdada/g2gserver/internal/pkg/config"
	"ohdada/g2gserver/internal/pkg/constant/envname"
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/network"
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/platform"
	"ohdada/g2gserver/internal/pkg/static/data"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// PlayerWebsocketHandler .
func PlayerWebsocketHandler(context *gin.Context) {
	platformName := context.Query("platform")
	language := context.Query("language")

	pbPlatformProviderMap := data.DefaultManager().GetPlatformProviderMap()

	pbPlatformProvider, isExists := pbPlatformProviderMap[platformName]
	if !isExists {
		http.Error(context.Writer, fmt.Sprintf("platform name [%s] not exists.", platformName), http.StatusUnauthorized)
		return
	}

	platformProvider := platform.CreateProvider(pbPlatformProvider)

	if platformProvider == nil {
		http.Error(context.Writer, fmt.Sprintf("platform provider [%s] not exists.", platformName), http.StatusUnauthorized)
		return
	}

	platformProvider.Init(pbPlatformProvider)

	platformProvider.SetPublicIPAddress(ServerInfo().PublicIpAddress)
	platformProvider.SetRunMode(config.GetString(envname.RunMode))
	platformProvider.SetLanguage(language)

	platformPlayer, err := platformProvider.Auth(context.Request)

	if err != nil {
		glog.Error(err)
		http.Error(context.Writer, fmt.Sprintf("Auth error : [%s]", err.Error()), http.StatusUnauthorized)
		return
	}

	glog.Infof("PlayerID:%s", platformPlayer.IdAtPlatform)

	websocketConnection, connectionError := upgrader.Upgrade(context.Writer, context.Request, nil)

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
		SetID(platformPlayer.IdAtPlatform).
		SetConnection(websocketConnection).
		SetReadTimeoutDuration(time.Duration(idleTimeoutSecond) * time.Second).
		SetWriteTimeoutDuration(time.Duration(idleTimeoutSecond) * time.Second).
		SetPlatformProvider(platformProvider).
		Build()

	go wc.ObserveReadMessage()
	go wc.ObserveWriteMessage()

	if coreGame.IsUnderMaintain() {
		wc.SendError(int32(pb.ErrorCode_EC_LOGIN_UNDER_MAINTAIN), pb.ErrorAction_EA_CLOSE_CONNECTION)
		wc.Close()
		return
	}

	defaultPlayerWebsocketConnectionEventHandler.OnConnect(wc)
}
