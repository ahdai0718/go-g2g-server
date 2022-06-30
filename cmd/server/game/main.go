package main

import (
	"fmt"
	"ohdada/g2gserver/internal/app/game"
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

	game.Run()

	httpserver := gin.Default()

	serverCorsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	httpserver.Use(cors.New(serverCorsConfig))

	httpserver.GET("/ws", game.PlayerWebsocketHandler)

	pprof.Register(httpserver)

	serverInfo := game.ServerInfo()
	glog.Infoln("ServerInfo", serverInfo)

	httpserver.Run(fmt.Sprintf("%s:%d", serverInfo.Host, serverInfo.Port))
}
