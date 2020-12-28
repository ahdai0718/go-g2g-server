package config

import (
	"ohdada/g2gserver/internal/pkg/glog"
	"flag"

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
}

// Init .
func Init() {
	glog.Infoln("Config init()...")

	readConfigFile("server")
	readConfigFile("game")
	readCommandLine()

	err := viper.Unmarshal(config)
	if err != nil {
		glog.Errorf("Fatal error config parse: %s \n", err)
	}

	glog.Infoln(config)
}

// PrintAllSettings .
func PrintAllSettings() {
	glog.Infoln(viper.AllSettings())
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
		glog.Errorf("Fatal error config file: %s \n", err)
	}
}

func readCommandLine() {
	glog.Infoln("flag.CommandLine", flag.CommandLine)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}
