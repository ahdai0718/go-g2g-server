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
	GetAllPlatformProvider() (platformProviderList []*pb.PlatformProvider, err error)
	GetPlatformProviderByName(name string) (platformProvider *pb.PlatformProvider, err error)
}

// Manager .
type Manager struct {
	Storer
	// storer Storer
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
	case pb.StoreDriver_SD_ORM:
		storer = &StorerORM{}
	default:
		storer = &StorerORM{}
	}

	err := storer.init(storeConnection)
	if err != nil {
		glog.Error(err)
	}

	manager.Storer = storer

	return err
}
