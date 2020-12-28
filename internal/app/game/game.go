package game

import (
	"ohdada/g2gserver/internal/pkg/config"
	"ohdada/g2gserver/internal/pkg/constant"
	"ohdada/g2gserver/internal/pkg/constant/flagname"
	"ohdada/g2gserver/internal/pkg/convertor/currency"
	pkgCoreGame "ohdada/g2gserver/internal/pkg/core/game"
	"ohdada/g2gserver/internal/pkg/core/serverinfo"
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/network"
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/static/data"
	"ohdada/g2gserver/internal/pkg/store"
	"fmt"

	"github.com/jinzhu/copier"
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
	initCoreGame()
	startRoutine()

	if config.GetBool(flagname.UseDatabaseServerInfo) {
		setServerInfoFromStaticData()
	}

	glog.SetCurrentServerInfo(ServerInfo())

	staticServerInfoMap := data.DefaultManager().GetServerInfoMap()

	if staticFluentdServerInfoMap, ok := staticServerInfoMap[pb.ServerType_ST_FLUENTD]; ok {

		staticFluentdServerInfoList := make([]*pb.ServerInfo, 0)

		if len(staticFluentdServerInfoMap) > 0 {
			for _, staticFluentdServerInfo := range staticFluentdServerInfoMap {
				glog.SetLogServerInfo(staticFluentdServerInfo)
				staticFluentdServerInfoList = append(staticFluentdServerInfoList, staticFluentdServerInfo)
			}
		}
		coreGame.SetLogServerInfo(staticFluentdServerInfoList)
	}

	defaultServerWebsocketConnectionManager().syncOnLaunch()

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

	data.DefaultManager().LoadFromStore()

	currency.DefaultConvertor().SetCurrencyMap(data.DefaultManager().GetCurrencyMap())

	if err := DefaultPlayerManager().init(); err != nil {
		panic(err)
	}
}

func initCoreGame() {
	coreGame = pkgCoreGame.DefaultSimpleFactory().Create(ServerInfo().GameType)

	if coreGame == nil {
		panic(fmt.Errorf("Game type [%d] is not exists", ServerInfo().GameType))
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
		panic(fmt.Errorf("Game type [%d] is not exists", ServerInfo().GameType))
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

func setServerInfoFromStaticData() {

	serverInfo := serverinfo.DefaultManager().Localhost()

	staticServerInfo := data.DefaultManager().GetServerInfo(pb.ServerType_ST_GAME, serverInfo.Name)

	if staticServerInfo != nil {
		copier.Copy(serverInfo, staticServerInfo)
	}

	serverinfo.DefaultManager().SetLocalhost(serverInfo)
}
