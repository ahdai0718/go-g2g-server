package gateway

import (
	"ohdada/g2gserver/internal/pkg/config"
	"ohdada/g2gserver/internal/pkg/constant"
	"ohdada/g2gserver/internal/pkg/constant/flagname"
	"ohdada/g2gserver/internal/pkg/convertor/currency"
	"ohdada/g2gserver/internal/pkg/core/serverinfo"
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/network"
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/static/data"
	"ohdada/g2gserver/internal/pkg/store"
	"time"

	"github.com/jinzhu/copier"
)

var (
	maxServer         = 32
	maxPlayer         = 1024
	idleTimeoutSecond = 600

	defaultPlayerWebsocketConnectionEventHandler = network.NewWebsocketConnectionEventHandler(maxPlayer, defaultPlayerWebsocketConnectionManager())
	defaultServerWebsocketConnectionEventHandler = network.NewWebsocketConnectionEventHandler(maxServer, defaultServerWebsocketConnectionManager())

	runMode = constant.RunModeRelease

	storeConnection = &pb.StoreConnection{}
	ticker          = time.NewTicker(time.Second)
)

// Run .
func Run() {
	loadConfig()
	initManager()
	startRoutine()

	if config.GetBool(flagname.UseDatabaseServerInfo) {
		setServerInfoFromStaticData()
	}

	glog.SetCurrentServerInfo(ServerInfo())

	staticServerInfoMap := data.DefaultManager().GetServerInfoMap()

	if staticFluentdServerInfoList, ok := staticServerInfoMap[pb.ServerType_ST_FLUENTD]; ok {
		if len(staticFluentdServerInfoList) > 0 {
			for _, staticFluentdServerInfo := range staticFluentdServerInfoList {
				glog.SetLogServerInfo(staticFluentdServerInfo)
			}
		}
	}
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

	if err := serverinfo.DefaultManager().Init(); err != nil {
		panic(err)
	}

	if err := store.DefaultManager().Init(storeConnection); err != nil {
		panic(err)
	}
	data.DefaultManager().LoadFromStore()

	currency.DefaultConvertor().SetCurrencyMap(data.DefaultManager().GetCurrencyMap())

	if err := DefaultPlayerManager().init(); err != nil {
		panic(err)
	}

}

func startRoutine() {
	go defaultPlayerWebsocketConnectionEventHandler.Observe()
	go defaultServerWebsocketConnectionEventHandler.Observe()

	go defaultServerWebsocketConnectionManager().observe()

	go DefaultPlayerManager().observe()

	go observe()
}

func setServerInfoFromStaticData() {

	serverInfo := serverinfo.DefaultManager().Localhost()

	staticServerInfo := data.DefaultManager().GetServerInfo(pb.ServerType_ST_GATEWAY, serverInfo.Name)

	if staticServerInfo != nil {
		copier.Copy(serverInfo, staticServerInfo)
	}

	serverinfo.DefaultManager().SetLocalhost(serverInfo)
}

func observe() {
	for {
		select {
		case <-ticker.C:
			store.DefaultManager().Tick()
		}
	}
}
