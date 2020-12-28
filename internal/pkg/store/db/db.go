package db

import (
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/store/db/model"
	"database/sql"
	"fmt"

	// Package mysql provides a MySQL driver for Go's database/sql package.
	_ "github.com/go-sql-driver/mysql"
)

// Database driver name
const (
	DriverNameMySQL = "mysql"
)

var (
	driverName = DriverNameMySQL
	connection *pb.StoreConnection
	db         *sql.DB

	serverInfoModel       *model.ServerInfo
	platformProviderModel *model.PlatformProvider
	currencyModel         *model.Currency
	transactionModel      *model.Transaction
)

// Init .
func Init(storeConnection *pb.StoreConnection) (err error) {
	glog.Infoln(storeConnection)

	connection = storeConnection

	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8",
		connection.User,
		connection.Password,
		connection.Host,
		connection.Port,
		connection.Schema)

	db, err = sql.Open(driverName, dataSourceName)

	if err != nil {
		return err
	}

	initModels()

	return nil
}

// SetMaxIdleConns .
func SetMaxIdleConns(n int) {
	if db != nil {
		db.SetMaxIdleConns(n)
	}
}

// SetMaxOpenConns .
func SetMaxOpenConns(n int) {
	if db != nil {
		db.SetMaxOpenConns(n)
	}
}

// CreateTable .
func CreateTable() (err error) {
	_, err = db.Exec(fmt.Sprintf("CALL %s.Table_SP_CreateAllMonth();", connection.Schema))
	if err != nil {
		glog.Error(err)
	}
	return
}

func initModels() {
	serverInfoModel = model.NewServerInfo(db, connection)
	platformProviderModel = model.NewPlatformProvider(db, connection)
	currencyModel = model.NewCurrency(db, connection)
	transactionModel = model.NewTransaction(db, connection)
}
