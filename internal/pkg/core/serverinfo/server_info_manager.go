package serverinfo

import (
	"ohdada/g2gserver/internal/pkg/pb"
	"sync"
)

var (
	defaultManager *Manager
)

// DefaultManager .
func DefaultManager() *Manager {
	if defaultManager == nil {
		defaultManager = &Manager{
			localhost:               &pb.ServerInfo{},
			currentRunningServerMap: &sync.Map{},
		}
	}
	return defaultManager
}

// Manager .
type Manager struct {
	localhost *pb.ServerInfo
	// currentRunningServerMap *pb.ServerInfoMapByName
	currentRunningServerMap *sync.Map
}

// Init .
func (manager *Manager) Init() (err error) {
	return
}

// SetLocalhost .
func (manager *Manager) SetLocalhost(serverInfo *pb.ServerInfo) {
	manager.localhost = serverInfo
}

// Localhost .
func (manager *Manager) Localhost() *pb.ServerInfo {
	return manager.localhost
}

// CurrentRunningServerMap .
func (manager *Manager) CurrentRunningServerMap() (serverInfoMapByName *pb.ServerInfoMapByName) {
	serverInfoMapByName = &pb.ServerInfoMapByName{
		Data: make(map[string]*pb.ServerInfo),
	}
	manager.currentRunningServerMap.Range(func(key interface{}, value interface{}) bool {
		serverInfoMapByName.Data[key.(string)] = value.(*pb.ServerInfo)
		return true
	})
	return
}

// UpdateRunningServer .
func (manager *Manager) UpdateRunningServer(serverInfo *pb.ServerInfo) {
	manager.currentRunningServerMap.Store(serverInfo.Name, serverInfo)
}

// RemoveRunningServer .
func (manager *Manager) RemoveRunningServer(serverInfo *pb.ServerInfo) {
	manager.currentRunningServerMap.Delete(serverInfo.Name)
}

// RemoveRunningServerByName .
func (manager *Manager) RemoveRunningServerByName(serverName string) {
	manager.currentRunningServerMap.Delete(serverName)
}
