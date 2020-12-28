package data

import (
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/store"
)

var (
	defaultManagerInstance *Manager
)

// Manager .
type Manager struct {
	serverInfoMap       map[pb.ServerType]map[string]*pb.ServerInfo
	platformProviderMap map[string]*pb.PlatformProvider
	currencyMap         map[string]*pb.Currency
}

// DefaultManager .
func DefaultManager() *Manager {
	if defaultManagerInstance == nil {
		defaultManagerInstance = &Manager{
			serverInfoMap:       make(map[pb.ServerType]map[string]*pb.ServerInfo),
			platformProviderMap: make(map[string]*pb.PlatformProvider),
			currencyMap:         make(map[string]*pb.Currency),
		}
	}
	return defaultManagerInstance
}

// LoadFromStore .
func (manager *Manager) LoadFromStore() {
	manager.loadServerInfoFromStore()
	manager.loadPlatformProviderFromStore()
	manager.loadCurrencyFromStore()
}

// GetServerInfoMap .
func (manager *Manager) GetServerInfoMap() map[pb.ServerType]map[string]*pb.ServerInfo {
	return manager.serverInfoMap
}

// GetServerInfoMapByType .
func (manager *Manager) GetServerInfoMapByType(serverType pb.ServerType) map[string]*pb.ServerInfo {
	if serverInfoMap, isExists := manager.serverInfoMap[serverType]; isExists {
		return serverInfoMap
	}

	glog.Warningf("Server info map by type [%s] is nil", serverType)
	return nil
}

// GetServerInfo .
func (manager *Manager) GetServerInfo(serverType pb.ServerType, serverName string) *pb.ServerInfo {
	serverInfoMap := manager.GetServerInfoMapByType(serverType)

	if serverInfoMap != nil {
		if serverInfo, isExists := serverInfoMap[serverName]; isExists {
			return serverInfo
		}

		glog.Warningf("Server info [%s] is nil", serverName)
	}

	return nil
}

// GetPlatformProviderMap .
func (manager *Manager) GetPlatformProviderMap() map[string]*pb.PlatformProvider {
	return manager.platformProviderMap
}

// GetCurrencyMap .
func (manager *Manager) GetCurrencyMap() map[string]*pb.Currency {
	return manager.currencyMap
}

func (manager *Manager) loadServerInfoFromStore() {
	serverInfoStoreList, err := store.DefaultManager().GetAllServerInfo()

	if err != nil {
		panic(err)
	}

	for _, serverInfoStore := range serverInfoStoreList {
		if manager.serverInfoMap[serverInfoStore.Type] == nil {
			manager.serverInfoMap[serverInfoStore.Type] = make(map[string]*pb.ServerInfo)
		}

		manager.serverInfoMap[serverInfoStore.Type][serverInfoStore.Name] = serverInfoStore
	}
}

func (manager *Manager) loadPlatformProviderFromStore() {
	storePlatformProviderList, err := store.DefaultManager().GetAllPlatformProvider()

	if err != nil {
		panic(err)
	}

	for _, storePlatformProvider := range storePlatformProviderList {
		manager.platformProviderMap[storePlatformProvider.Name] = storePlatformProvider
	}
}

func (manager *Manager) loadCurrencyFromStore() {
	storeCurrencyList, err := store.DefaultManager().GetAllCurrency()

	if err != nil {
		panic(err)
	}

	for _, storeCurrency := range storeCurrencyList {
		manager.currencyMap[storeCurrency.Name] = storeCurrency
	}
}
