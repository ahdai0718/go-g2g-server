package main

import (
	"fmt"
	"ohdada/g2gserver/internal/app/gateway"
	"ohdada/g2gserver/internal/pkg/config"
	"ohdada/g2gserver/internal/pkg/constant"
	"ohdada/g2gserver/internal/pkg/constant/envname"
	"ohdada/g2gserver/internal/pkg/glog"
	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func init() {
	glog.Init()

	config.Init()
	config.PrintAllSettings()

	gin.SetMode(config.GetString(envname.GinMode))

	if config.GetString(envname.RunMode) == constant.RunModeDev {
		runtime.SetBlockProfileRate(1)
	}
}

func main() {
	if config.GetString(envname.RunMode) != constant.RunModeDev {
		defer func() {
			if err := recover(); err != nil {
				glog.Error(err)
			}
		}()
	}

	gateway.Run()

	httpserver := gin.Default()

	serverCorsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	httpserver.Use(cors.New(serverCorsConfig))

	httpserver.GET("/server/ws", gateway.ServerWebsocketHandler)
	httpserver.GET("/ws", gateway.PlayerWebsocketHandler)

	httpserver.POST("/api/oauth/access_token", gateway.HandleOAuthAccessToken)
	httpserver.POST("/api/player/token/validate", gateway.CheckOAuthToken, gateway.CheckAccessToken, gateway.HandlePlayerTokenValidate)

	pprof.Register(httpserver)

	serverInfo := gateway.ServerInfo()
	glog.Infoln("ServerInfo", serverInfo)

	httpserver.Run(fmt.Sprintf("%s:%d", serverInfo.Host, serverInfo.Port))

}
