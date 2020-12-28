package db

import "ohdada/g2gserver/internal/pkg/store/db/entity"

// ServerInfoGetAll .
func ServerInfoGetAll() ([]entity.ServerInfo, error) {
	return serverInfoModel.GetAll()
}

// ServerInfoGetByName .
func ServerInfoGetByName(name string) (entity.ServerInfo, error) {
	return serverInfoModel.GetByName(name)
}
