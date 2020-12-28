package main

import (
	"ohdada/g2gserver/internal/app/bot"
	"ohdada/g2gserver/internal/pkg/config"
	"ohdada/g2gserver/internal/pkg/constant"
	"ohdada/g2gserver/internal/pkg/constant/flagname"
	"ohdada/g2gserver/internal/pkg/glog"
	"flag"
	"runtime"
)

var _ = flag.String(flagname.ServerHostForClient, "", "Gateway host (name|ip)")
var _ = flag.String(flagname.ServerPortForClient, "", "Gateway host port")
var _ = flag.String(flagname.ServerWebSocketProtocolForClient, "", "Gateway websocket protocol (ws|wss)")
var _ = flag.String(flagname.ServerWebSocketRoutePathForClient, "", "Gateway websocket route path")
var _ = flag.Int(flagname.MaxBot, 4, "Max bot")

func init() {
	flag.Parse()

	config.Init()
	config.PrintAllSettings()

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

	bot.Run()

	select {}
}
