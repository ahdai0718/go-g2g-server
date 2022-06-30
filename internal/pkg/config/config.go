package config

import (
	"flag"
	"ohdada/g2gserver/internal/pkg/constant/envname"
	"ohdada/g2gserver/internal/pkg/glog"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	config = &Config{}
)

// Config .
type Config struct {
	// AuthMap             map[string]*pb.Auth             `mapstructure:"auth"`
	// PlatformProviderMap map[string]*pb.PlatformProvider `mapstructure:"platform"`

	ServerName               string `mapstructure:"SERVER_NAME,omitempty,remain"`
	ServerHost               string `mapstructure:"SERVER_HOST,omitempty,remain"`
	ServerPort               string `mapstructure:"SERVER_PORT,omitempty,remain"`
	ServerGroup              string `mapstructure:"SERVER_GROUP,omitempty,remain"`
	ServerWebSocketProtocol  string `mapstructure:"SERVER_WS_PROTOCOL,omitempty,remain"`
	ServerWebSocketRoutePath string `mapstructure:"SERVER_WS_ROUTE_PATH,omitempty,remain"`

	ServerHostForClient               string `mapstructure:"SERVER_HOST_FOR_CLIENT,omitempty,remain"`
	ServerPortForClient               string `mapstructure:"SERVER_PORT_FOR_CLIENT,omitempty,remain"`
	ServerWebSocketProtocolForClient  string `mapstructure:"SERVER_WS_PROTOCOL_FOR_CLIENT,omitempty,remain"`
	ServerWebSocketRoutePathForClient string `mapstructure:"SERVER_WS_ROUTE_PATH_FOR_CLIENT,omitempty,remain"`

	IPAddress string `mapstructure:"IP_ADDRESS,omitempty,remain"`

	DatabaseServerHost          string `mapstructure:"DB_HOST,omitempty,remain"`
	DatabaseServerPort          string `mapstructure:"DB_PORT,omitempty,remain"`
	DatabaseServerSchema        string `mapstructure:"DB_SCHEMA,omitempty,remain"`
	DatabaseServerSchemaGame    string `mapstructure:"DB_SCHEMA_GAME,omitempty,remain"`
	DatabaseServerUser          string `mapstructure:"DB_USER,omitempty,remain"`
	DatabaseServerPassword      string `mapstructure:"DB_PASSWORD,omitempty,remain"`
	DatabaseServerMaxConnection string `mapstructure:"DB_MAX_CONNECTION,omitempty,remain"`

	RunMode string `mapstructure:"RUN_MODE,omitempty,remain"`
	GinMode string `mapstructure:"GIN_MODE,omitempty,remain"`

	PlatformProviderFactoryName string `mapstructure:"PLATFORM_PROVIDER_FACTORY_NAME,omitempty,remain"`
	IdleTimeoutSecond           string `mapstructure:"IDLE_TIMEOUT_SECOND,omitempty,remain"`
	GameType                    string `mapstructure:"GAME_TYPE,omitempty,remain"`
	MaxBot                      string `mapstructure:"MAX_BOT,omitempty,remain"`

	MQType string `mapstructure:"MQ_TYPE,omitempty,remain"`
	MQUris string `mapstructure:"MQ_URIS,omitempty,remain"`

	AuthApiUriBase string `mapstructure:"AUTH_API_URI_BASE,omitempty,remain"`
	AuthSecret     string `mapstructure:"AUTH_SECRET,omitempty,remain"`

	CacheServerHosts string `mapstructure:"CACHE_SERVER_HOSTS,omitempty,remain"`
	CacheServerPorts string `mapstructure:"CACHE_SERVER_PORTS,omitempty,remain"`
}

// Init .
func Init() {
	glog.Infoln("Config init()...")

	bindEnv()

	err := viper.Unmarshal(config)
	if err != nil {
		glog.Errorf("Fatal error config parse: %s \n", err)
	}

	glog.Infoln("config:", config)
}

// PrintAllSettings .
func PrintAllSettings() {
	glog.Infoln("viper.AllSettings():", viper.AllSettings())
}

// GetBool .
func GetBool(key string) bool {
	return viper.GetBool(key)
}

// GetInt .
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetString .
func GetString(key string) string {
	return viper.GetString(key)
}

// GetStringSlice .
func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

// GetConfig .
func GetConfig() *Config {
	if config == nil {
		glog.Error("Config is nil")
	}
	return config
}

func readConfigFile(configName string) {
	viper.SetConfigName(configName)
	viper.AddConfigPath("../../../configs")
	viper.AddConfigPath("../../configs")
	err := viper.MergeInConfig()
	if err != nil {
		glog.Warningf("Fatal error config file: %s \n", err)
	}
}

func readCommandLine() {
	glog.Infoln("flag.CommandLine", flag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func bindEnv() {
	viper.BindEnv(envname.ServerName)
	viper.BindEnv(envname.ServerHost)
	viper.BindEnv(envname.ServerPort)
	viper.BindEnv(envname.ServerGroup)
	viper.BindEnv(envname.ServerWebSocketProtocol)
	viper.BindEnv(envname.ServerWebSocketRoutePath)

	viper.BindEnv(envname.ServerHostForClient)
	viper.BindEnv(envname.ServerPortForClient)
	viper.BindEnv(envname.ServerWebSocketProtocolForClient)
	viper.BindEnv(envname.ServerWebSocketRoutePathForClient)

	viper.BindEnv(envname.IPAddress)

	viper.BindEnv(envname.DatabaseServerHost)
	viper.BindEnv(envname.DatabaseServerPort)
	viper.BindEnv(envname.DatabaseServerSchema)
	viper.BindEnv(envname.DatabaseServerSchemaGame)
	viper.BindEnv(envname.DatabaseServerUser)
	viper.BindEnv(envname.DatabaseServerPassword)
	viper.BindEnv(envname.DatabaseServerMaxConnection)

	viper.BindEnv(envname.RunMode)
	viper.BindEnv(envname.GinMode)

	viper.BindEnv(envname.PlatformProviderFactoryName)
	viper.BindEnv(envname.IdleTimeoutSecond)
	viper.BindEnv(envname.GameType)
	viper.BindEnv(envname.MaxBot)

	viper.BindEnv(envname.MQType)
	viper.BindEnv(envname.MQUris)

	viper.BindEnv(envname.AuthApiUriBase)
	viper.BindEnv(envname.AuthSecret)

	viper.BindEnv(envname.CacheServerHosts)
	viper.BindEnv(envname.CacheServerPorts)
}
