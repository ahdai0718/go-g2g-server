package cache

import (
	"ohdada/g2gserver/internal/pkg/constant"

	"github.com/golang/glog"
)

var (
	// Cache
	serverHostList = []string{}
	serverPortList = []string{}
	client         Client
)

// Init .
func Init(config *Config) (err error) {

	client = Create(config.Type)

	err = client.Init(config.ServerConfigList)
	if err != nil {
		glog.Error(err)
		return err
	}

	err = client.Set("test", "test")
	if err != nil {
		glog.Error(err)
		return err
	}

	if config.RunMode == constant.RunModeDev {
		var value string
		client.Get("test", &value)
		glog.Infof("cache server get key test : %v", value)
	}

	return
}

// DefaultClient .
func DefaultClient() Client {
	if client == nil {
		panic("should init cache client first")
	}

	return client
}
