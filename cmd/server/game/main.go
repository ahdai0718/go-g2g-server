package main

import (
	"ohdada/g2gserver/internal/app/game"
	"ohdada/g2gserver/internal/pkg/config"
	"ohdada/g2gserver/internal/pkg/constant"
	"ohdada/g2gserver/internal/pkg/constant/flagname"
	"ohdada/g2gserver/internal/pkg/glog"
	"flag"
	"fmt"
	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var _ = flag.String(flagname.UseDatabaseServerInfo, "false", "Use database server info")

var _ = flag.String(flagname.ServerName, "", "Server name for")
var _ = flag.String(flagname.ServerHost, "", "Server host (name|ip)")
var _ = flag.String(flagname.ServerPort, "", "Server host port")
var _ = flag.String(flagname.ServerGroup, "", "Game server group name")
var _ = flag.String(flagname.ServerWebSocketProtocol, "", "Game server websocket protocol (ws|wss)")
var _ = flag.String(flagname.ServerWebSocketRoutePath, "", "Game server websocket route path")
var _ = flag.String(flagname.IPAddress, "", "Server public ip address")

var _ = flag.String(flagname.ServerHostForClient, "", "Server host for client")
var _ = flag.String(flagname.ServerPortForClient, "", "Server host port for client")
var _ = flag.String(flagname.ServerWebSocketProtocolForClient, "", "Game server websocket protocol for client(ws|wss)")
var _ = flag.String(flagname.ServerWebSocketRoutePathForClient, "", "Game server websocket route path for client")

var _ = flag.String(flagname.DatabaseServerHost, "", "Database host")
var _ = flag.String(flagname.DatabaseServerPort, "", "Database port")
var _ = flag.String(flagname.DatabaseServerSchema, "", "Database schema")
var _ = flag.String(flagname.DatabaseServerSchemaGame, "", "Database schema game")
var _ = flag.String(flagname.DatabaseServerUser, "", "Database user")
var _ = flag.String(flagname.DatabaseServerPassword, "", "Database password")
var _ = flag.Int64(flagname.DatabaseServerMaxConnection, 16, "Database max connections")

var _ = flag.String(flagname.RunMode, "release", "Game server run mode (dev|debug|release|test)")
var _ = flag.String(flagname.GinMode, "release", "Gin http server run mode (debug|release|test)")

var _ = flag.String(flagname.PlatformProviderFactoryName, "default", "Platform provider factory name")
var _ = flag.String(flagname.IdleTimeoutSecond, "600", "Timeout seconds for idle")
var _ = flag.Int(flagname.GameType, 0, "Game type")

func init() {
	flag.Parse()

	config.Init()
	config.PrintAllSettings()

	gin.SetMode(config.GetString(flagname.GinMode))

	if config.GetString(flagname.RunMode) == constant.RunModeDev {
		runtime.SetBlockProfileRate(1)
	}
}

func main() {

	if config.GetString(flagname.RunMode) != constant.RunModeDev {
		defer func() {
			if err := recover(); err != nil {
				glog.Error(err)
			}
		}()
	}

	game.Run()

	var httpserver *gin.Engine

	if config.GetString(flagname.GinMode) == gin.ReleaseMode {
		httpserver = gin.New()
	} else {
		httpserver = gin.Default()
	}

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
