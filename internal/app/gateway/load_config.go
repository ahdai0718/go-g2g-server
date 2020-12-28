package gateway

import (
	"ohdada/g2gserver/internal/pkg/config"
	"ohdada/g2gserver/internal/pkg/constant/flagname"
	"ohdada/g2gserver/internal/pkg/core/serverinfo"
	"ohdada/g2gserver/internal/pkg/pb"
)

func loadConfig() {
	loadMisc()
	loadServerInfo()
	loadStoreConnection()

}

func loadServerInfo() {
	serverinfo.DefaultManager().SetLocalhost(&pb.ServerInfo{
		Name:                        config.GetString(flagname.ServerName),
		Host:                        config.GetString(flagname.ServerHost),
		Port:                        int64(config.GetInt(flagname.ServerPort)),
		Group:                       config.GetString(flagname.ServerGroup),
		WebsocketProtocol:           config.GetString(flagname.ServerWebSocketProtocol),
		WebsocketRoutePath:          config.GetString(flagname.ServerWebSocketRoutePath),
		HostForClient:               config.GetString(flagname.ServerHostForClient),
		PortForClient:               int64(config.GetInt(flagname.ServerPortForClient)),
		WebsocketProtocolForClient:  config.GetString(flagname.ServerWebSocketProtocolForClient),
		WebsocketRoutePathForClient: config.GetString(flagname.ServerWebSocketRoutePathForClient),
	})
}

func loadStoreConnection() {
	storeConnection.Driver = pb.StoreDriver_SD_DATABASE
	storeConnection.Host = config.GetString(flagname.DatabaseServerHost)
	storeConnection.Port = int64(config.GetInt(flagname.DatabaseServerPort))
	storeConnection.Schema = config.GetString(flagname.DatabaseServerSchema)
	storeConnection.User = config.GetString(flagname.DatabaseServerUser)
	storeConnection.Password = config.GetString(flagname.DatabaseServerPassword)
	storeConnection.MaxConnection = int64(config.GetInt(flagname.DatabaseServerMaxConnection))
}

func loadMisc() {
	runMode = config.GetString(flagname.RunMode)
	idleTimeoutSecond = config.GetInt(flagname.IdleTimeoutSecond)
}
