package model

import (
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/store/db/entity"
	"database/sql"
)

const (
	_ = iota
	stmtServerInfoGetAll
	stmtServerInfoGetByName
)

// NewServerInfo .
func NewServerInfo(db *sql.DB, connection *pb.StoreConnection) *ServerInfo {
	model := &ServerInfo{}
	model.db = db
	model.connection = connection
	model.stmtMap = make(map[int]*sql.Stmt)

	var err error

	if _, err = model.genStmt(stmtServerInfoGetAll, model.genSPStrWithSchema("ServerInfo_SP_GetAll()")); err != nil {
		glog.Error(err)
	}

	if _, err = model.genStmt(stmtServerInfoGetByName, model.genSPStrWithSchema("ServerInfo_SP_GetByName(?)")); err != nil {
		glog.Error(err)
	}

	return model
}

// ServerInfo .
type ServerInfo struct {
	base
}

// GetAll .
func (model *ServerInfo) GetAll() ([]entity.ServerInfo, error) {
	stmt := model.getStmt(stmtServerInfoGetAll)
	rows, err := stmt.Query()

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		glog.Error(err)
		return nil, err
	}

	serverInfoList := make([]entity.ServerInfo, 0)

	for rows.Next() {
		serverInfo := entity.ServerInfo{}
		err = rows.Scan(
			&serverInfo.Name,
			&serverInfo.Group,
			&serverInfo.Type,
			&serverInfo.Host,
			&serverInfo.Port,
			&serverInfo.Protocol,
			&serverInfo.ServerInfoRoutePath,
			&serverInfo.WebsocketProtocol,
			&serverInfo.WebsocketRoutePath,
			&serverInfo.HostForClient,
			&serverInfo.PortForClient,
			&serverInfo.WebsocketProtocolForClient,
			&serverInfo.WebsocketRoutePathForClient,
			&serverInfo.PublicIPAddress,
			&serverInfo.IsOffline,
			&serverInfo.GameType)

		if err != nil {
			glog.Error(err)
			return nil, err
		}

		serverInfoList = append(serverInfoList, serverInfo)
	}

	return serverInfoList, nil
}

// GetByName .
func (model *ServerInfo) GetByName(name string) (entity.ServerInfo, error) {
	serverInfo := entity.ServerInfo{}

	var err error

	stmt := model.getStmt(stmtServerInfoGetByName)

	err = stmt.QueryRow(name).Scan(
		&serverInfo.Name,
		&serverInfo.Group,
		&serverInfo.Type,
		&serverInfo.Host,
		&serverInfo.Port,
		&serverInfo.Protocol,
		&serverInfo.ServerInfoRoutePath,
		&serverInfo.WebsocketProtocol,
		&serverInfo.WebsocketRoutePath,
		&serverInfo.HostForClient,
		&serverInfo.PortForClient,
		&serverInfo.WebsocketProtocolForClient,
		&serverInfo.WebsocketRoutePathForClient,
		&serverInfo.PublicIPAddress,
		&serverInfo.IsOffline,
		&serverInfo.GameType)

	if err != nil {
		glog.Error(err)
		return serverInfo, err
	}

	return serverInfo, nil
}
