package game

import (
	"fmt"
	"ohdada/g2gserver/internal/pkg/constant"
	pkgCoreGame "ohdada/g2gserver/internal/pkg/core/game"
	"ohdada/g2gserver/internal/pkg/core/serverinfo"
	"ohdada/g2gserver/internal/pkg/network"
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/store"
)

var (
	maxServer         = 32
	maxPlayer         = 1024
	idleTimeoutSecond = 600

	runMode = constant.RunModeRelease

	defaultPlayerWebsocketConnectionEventHandler = network.NewWebsocketConnectionEventHandler(maxPlayer, defaultPlayerWebsocketConnectionManager())
	defaultServerWebsocketConnectionEventHandler = network.NewWebsocketConnectionEventHandler(maxServer, defaultServerWebsocketConnectionManager())

	serverWebsocketConnectionMap = make(map[string]*network.WebsocketConnection)

	storeConnection = &pb.StoreConnection{}
	coreGame        pkgCoreGame.Game
)

// Run .
func Run() {
	loadConfig()
	initManager()
	initMQ()
	initCoreGame()
	startRoutine()
	runCoreGame()
}

// ServerInfo .
func ServerInfo() *pb.ServerInfo {
	return serverinfo.DefaultManager().Localhost()
}

// RunMode .
func RunMode() string {
	return runMode
}

func initManager() {
	if err := store.DefaultManager().Init(storeConnection); err != nil {
		panic(err)
	}

	if err := DefaultBroadcastManager().init(); err != nil {
		panic(err)
	}

	if err := DefaultPlayerManager().init(); err != nil {
		panic(err)
	}
}

func initCoreGame() {
	coreGame = pkgCoreGame.DefaultSimpleFactory().Create(ServerInfo().GameType)

	if coreGame == nil {
		panic(fmt.Errorf("game type [%d] is not exists", ServerInfo().GameType))
	}

	coreGame.SetStoreConnection(storeConnection)

	if err := coreGame.Init(); err != nil {
		panic(err)
	}

	coreGame.SetPlayerChannelRemove(DefaultPlayerManager().PlayerChannelRemove())
	coreGame.SetServerBroadcastChannel(DefaultBroadcastManager().RequestChannel())
	coreGame.SetGatewayBroadcastChannel(DefaultBroadcastManager().RequestToGatewayChannel())
}

func runCoreGame() {
	if coreGame == nil {
		panic(fmt.Errorf("game type [%d] is not exists", ServerInfo().GameType))
	}

	if err := coreGame.Run(); err != nil {
		panic(err)
	}
}

func startRoutine() {
	go defaultPlayerWebsocketConnectionEventHandler.Observe()
	go defaultServerWebsocketConnectionEventHandler.Observe()

	go defaultServerWebsocketConnectionManager().observe()

	go DefaultBroadcastManager().observe()
	go DefaultPlayerManager().observe()
}
