package envname

const (
	UseDatabaseServerInfo = "USE_DATABASE_SERVER_INFO"

	ServerName               = "SERVER_NAME"
	ServerHost               = "SERVER_HOST"
	ServerPort               = "SERVER_PORT"
	ServerGroup              = "SERVER_GROUP"
	ServerWebSocketProtocol  = "SERVER_WS_PROTOCOL"
	ServerWebSocketRoutePath = "SERVER_WS_ROUTE_PATH"

	ServerHostForClient               = "SERVER_HOST_FOR_CLIENT"
	ServerPortForClient               = "SERVER_PORT_FOR_CLIENT"
	ServerWebSocketProtocolForClient  = "SERVER_WS_PROTOCOL_FOR_CLIENT"
	ServerWebSocketRoutePathForClient = "SERVER_WS_ROUTE_PATH_FOR_CLIENT"

	IPAddress = "IP_ADDRESS"

	DatabaseServerHost          = "DB_HOST"
	DatabaseServerPort          = "DB_PORT"
	DatabaseServerSchema        = "DB_SCHEMA"
	DatabaseServerSchemaGame    = "DB_SCHEMA_GAME"
	DatabaseServerUser          = "DB_USER"
	DatabaseServerPassword      = "DB_PASSWORD"
	DatabaseServerMaxConnection = "DB_MAX_CONNECTION"

	RunMode = "RUN_MODE"
	GinMode = "GIN_MODE"

	PlatformProviderFactoryName = "PLATFORM_PROVIDER_FACTORY_NAME"
	IdleTimeoutSecond           = "IDLE_TIMEOUT_SECOND"
	GameType                    = "GAME_TYPE"
	MaxBot                      = "MAX_BOT"

	MQType = "MQ_TYPE"
	MQUris = "MQ_URIS"

	AuthApiUriBase = "AUTH_API_URI_BASE"
	AuthSecret     = "AUTH_SECRET"

	CacheServerHosts = "CACHE_SERVER_HOSTS"
	CacheServerPorts = "CACHE_SERVER_PORTS"
)
