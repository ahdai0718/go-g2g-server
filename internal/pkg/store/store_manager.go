package store

import (
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/pb"
)

var (
	defaultManagerInstance *Manager
)

// Storer .
type Storer interface {
	init(storeConnection *pb.StoreConnection) (err error)
	tick() (err error)
	getAllServerInfo() ([]*pb.ServerInfo, error)
	getAllPlatformProvider() ([]*pb.PlatformProvider, error)
	getAllCurrency() ([]*pb.Currency, error)
	addTransactionLog(transaction *pb.Transaction) error
}

// Manager .
type Manager struct {
	storer Storer
}

// DefaultManager .
func DefaultManager() *Manager {
	if defaultManagerInstance == nil {
		defaultManagerInstance = &Manager{}
	}
	return defaultManagerInstance
}

// Init .
func (manager *Manager) Init(storeConnection *pb.StoreConnection) error {

	var storer Storer

	switch storeConnection.Driver {
	default:
		storer = &StorerDB{}
	}

	err := storer.init(storeConnection)
	if err != nil {
		glog.Error(err)
	}

	manager.storer = storer

	return err
}

// Tick .
func (manager *Manager) Tick() (err error) {
	return manager.storer.tick()
}

// GetAllServerInfo .
func (manager *Manager) GetAllServerInfo() ([]*pb.ServerInfo, error) {
	return manager.storer.getAllServerInfo()
}

// GetAllPlatformProvider .
func (manager *Manager) GetAllPlatformProvider() ([]*pb.PlatformProvider, error) {
	return manager.storer.getAllPlatformProvider()
}

// GetAllCurrency .
func (manager *Manager) GetAllCurrency() ([]*pb.Currency, error) {
	return manager.storer.getAllCurrency()
}

// AddTransactionLog .
func (manager *Manager) AddTransactionLog(transaction *pb.Transaction) error {
	return manager.storer.addTransactionLog(transaction)
}
