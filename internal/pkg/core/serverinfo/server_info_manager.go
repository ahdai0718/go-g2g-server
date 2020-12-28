package serverinfo

import "ohdada/g2gserver/internal/pkg/pb"

var (
	defaultManager *Manager
)

// DefaultManager .
func DefaultManager() *Manager {
	if defaultManager == nil {
		defaultManager = &Manager{
			localhost: &pb.ServerInfo{},
			currentRunningServerMap: &pb.ServerInfoMapByName{
				Data: make(map[string]*pb.ServerInfo),
			},
		}
	}
	return defaultManager
}

// Manager .
type Manager struct {
	localhost               *pb.ServerInfo
	currentRunningServerMap *pb.ServerInfoMapByName
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
func (manager *Manager) CurrentRunningServerMap() *pb.ServerInfoMapByName {
	return manager.currentRunningServerMap
}

// UpdateRunningServer .
func (manager *Manager) UpdateRunningServer(serverInfo *pb.ServerInfo) {
	manager.currentRunningServerMap.Data[serverInfo.Name] = serverInfo
}

// RemoveRunningServer .
func (manager *Manager) RemoveRunningServer(serverInfo *pb.ServerInfo) {
	delete(manager.currentRunningServerMap.Data, serverInfo.Name)
}

// RemoveRunningServerByName .
func (manager *Manager) RemoveRunningServerByName(serverName string) {
	delete(manager.currentRunningServerMap.Data, serverName)
}
