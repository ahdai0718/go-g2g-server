package model

import (
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/store/db/entity"

	"database/sql"
)

const (
	_ = iota
	stmtPlatformProviderGetAll
	stmtPlatformProviderGetByName
)

// NewPlatformProvider .
func NewPlatformProvider(db *sql.DB, connection *pb.StoreConnection) *PlatformProvider {
	model := &PlatformProvider{}
	model.db = db
	model.connection = connection
	model.stmtMap = make(map[int]*sql.Stmt)

	var err error

	if _, err = model.genStmt(stmtPlatformProviderGetAll, model.genSPStrWithSchema("PlatformProvider_SP_GetAll()")); err != nil {
		glog.Error(err)
	}

	if _, err = model.genStmt(stmtPlatformProviderGetByName, model.genSPStrWithSchema("PlatformProvider_SP_GetByName(?)")); err != nil {
		glog.Error(err)
	}

	return model
}

// PlatformProvider .
type PlatformProvider struct {
	base
}

// GetAll .
func (model *PlatformProvider) GetAll() ([]entity.PlatformProvider, error) {
	stmt := model.getStmt(stmtPlatformProviderGetAll)
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

	PlatformProviderList := make([]entity.PlatformProvider, 0)

	for rows.Next() {
		PlatformProvider := entity.PlatformProvider{}
		err = rows.Scan(
			&PlatformProvider.Name,
			&PlatformProvider.FactoryName,
			&PlatformProvider.AESKey,
			&PlatformProvider.AESIV,
			&PlatformProvider.AuthType,
			&PlatformProvider.AuthID,
			&PlatformProvider.AuthSecret,
			&PlatformProvider.AuthGrantType,
			&PlatformProvider.AuthScope,
			&PlatformProvider.APIURLBase)

		if err != nil {
			glog.Error(err)
			return nil, err
		}

		PlatformProviderList = append(PlatformProviderList, PlatformProvider)
	}

	return PlatformProviderList, nil
}

// GetByName .
func (model *PlatformProvider) GetByName(name string) (entity.PlatformProvider, error) {
	PlatformProvider := entity.PlatformProvider{}

	var err error

	stmt := model.getStmt(stmtPlatformProviderGetByName)

	err = stmt.QueryRow(name).Scan(
		&PlatformProvider.Name,
		&PlatformProvider.FactoryName,
		&PlatformProvider.AESKey,
		&PlatformProvider.AESIV,
		&PlatformProvider.AuthType,
		&PlatformProvider.AuthID,
		&PlatformProvider.AuthSecret,
		&PlatformProvider.AuthGrantType,
		&PlatformProvider.AuthScope,
		&PlatformProvider.APIURLBase)

	if err != nil {
		glog.Error(err)
		return PlatformProvider, err
	}

	return PlatformProvider, nil
}
