package game

import (
	"ohdada/g2gserver/internal/pkg/config"
	"ohdada/g2gserver/internal/pkg/constant/envname"
	"ohdada/g2gserver/internal/pkg/core/serverinfo"
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/pb/game"

	"github.com/google/uuid"
)

func loadConfig() {
	loadMisc()
	loadServerInfo()
	loadStoreConnection()
}

func loadServerInfo() {
	serverinfo.DefaultManager().SetLocalhost(&pb.ServerInfo{
		Id:                 uuid.New().String(),
		Name:               config.GetString(envname.ServerName),
		Host:               config.GetString(envname.ServerHost),
		Port:               int32(config.GetInt(envname.ServerPort)),
		Group:              config.GetString(envname.ServerGroup),
		WebsocketProtocol:  config.GetString(envname.ServerWebSocketProtocol),
		WebsocketRoutePath: config.GetString(envname.ServerWebSocketRoutePath),
		GameType:           game.GameType(config.GetInt(envname.GameType)),
	})
}

func loadStoreConnection() {
	storeConnection.Driver = pb.StoreDriver_SD_ORM
	storeConnection.Host = config.GetString(envname.DatabaseServerHost)
	storeConnection.Port = int64(config.GetInt(envname.DatabaseServerPort))
	storeConnection.Schema = config.GetString(envname.DatabaseServerSchema)
	storeConnection.SchemaGame = config.GetString(envname.DatabaseServerSchemaGame)
	storeConnection.User = config.GetString(envname.DatabaseServerUser)
	storeConnection.Password = config.GetString(envname.DatabaseServerPassword)
	storeConnection.MaxConnection = int64(config.GetInt(envname.DatabaseServerMaxConnection))
}

func loadMisc() {
	runMode = config.GetString(envname.RunMode)
	idleTimeoutSecond = config.GetInt(envname.IdleTimeoutSecond)
}
