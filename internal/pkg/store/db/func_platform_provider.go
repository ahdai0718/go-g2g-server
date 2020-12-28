package db

import "ohdada/g2gserver/internal/pkg/store/db/entity"

// PlatformProviderGetAll .
func PlatformProviderGetAll() ([]entity.PlatformProvider, error) {
	return platformProviderModel.GetAll()
}

// PlatformProviderGetByName .
func PlatformProviderGetByName(name string) (entity.PlatformProvider, error) {
	return platformProviderModel.GetByName(name)
}
