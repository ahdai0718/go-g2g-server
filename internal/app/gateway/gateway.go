package gateway

import (
	"ohdada/g2gserver/internal/pkg/constant"
	"ohdada/g2gserver/internal/pkg/core/serverinfo"
	"ohdada/g2gserver/internal/pkg/network"
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/store"
	"time"
)

var (
	maxServer         = 32
	maxPlayer         = 1024
	idleTimeoutSecond = 600

	defaultPlayerWebsocketConnectionEventHandler = network.NewWebsocketConnectionEventHandler(maxPlayer, defaultPlayerWebsocketConnectionManager())
	// defaultServerWebsocketConnectionEventHandler = network.NewWebsocketConnectionEventHandler(maxServer, defaultServerWebsocketConnectionManager())

	runMode = constant.RunModeRelease

	storeConnection = &pb.StoreConnection{}
	ticker          = time.NewTicker(time.Second)
)

// Run .
func Run() {
	loadConfig()
	initManager()
	initMQ()
	startRoutine()
}

// RunMode .
func RunMode() string {
	return runMode
}

// ServerInfo .
func ServerInfo() *pb.ServerInfo {
	return serverinfo.DefaultManager().Localhost()
}

func initManager() {
	if err := store.DefaultManager().Init(storeConnection); err != nil {
		panic(err)
	}

	if err := serverinfo.DefaultManager().Init(); err != nil {
		panic(err)
	}

	if err := DefaultPlayerManager().init(); err != nil {
		panic(err)
	}

}

func startRoutine() {
	go defaultPlayerWebsocketConnectionEventHandler.Observe()
	// go defaultServerWebsocketConnectionEventHandler.Observe()

	// go defaultServerWebsocketConnectionManager().observe()

	go DefaultPlayerManager().observe()

	go observe()

}

func observe() {
	for {
		select {
		case <-ticker.C:

		}
	}
}
