package main

import (
	"ohdada/g2gserver/internal/app/bot"
	"ohdada/g2gserver/internal/pkg/config"
	"ohdada/g2gserver/internal/pkg/constant"
	"ohdada/g2gserver/internal/pkg/constant/envname"
	"ohdada/g2gserver/internal/pkg/glog"
	"runtime"
)

func init() {
	config.Init()
	config.PrintAllSettings()

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

	bot.Run()

	select {}
}
