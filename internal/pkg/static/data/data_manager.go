package data

import (
	"ohdada/g2gserver/internal/pkg/config"
	"ohdada/g2gserver/internal/pkg/constant/envname"
	"ohdada/g2gserver/internal/pkg/pb"
)

var (
	defaultManagerInstance *Manager
)

// Manager .
type Manager struct {
	serverInfoMap       map[pb.ServerType]map[string]*pb.ServerInfo
	platformProviderMap map[string]*pb.PlatformProvider
}

// DefaultManager .
func DefaultManager() *Manager {
	if defaultManagerInstance == nil {
		defaultManagerInstance = &Manager{
			serverInfoMap:       make(map[pb.ServerType]map[string]*pb.ServerInfo),
			platformProviderMap: make(map[string]*pb.PlatformProvider),
		}

		authApiUriBase := config.GetString(envname.AuthApiUriBase)
		authSecret := config.GetString(envname.AuthSecret)

		defaultManagerInstance.platformProviderMap["default"] = &pb.PlatformProvider{
			FactoryName: "default",
			Name:        "default",
			ApiUrlBase:  authApiUriBase,
			Auth: &pb.Auth{
				Type:   pb.AuthType_AT_JWT,
				Secret: authSecret,
			},
		}
	}
	return defaultManagerInstance
}

// LoadFromStore .
func (manager *Manager) LoadFromStore() {

}

func (manager *Manager) GetPlatformProviderMap() map[string]*pb.PlatformProvider {
	return manager.platformProviderMap
}
