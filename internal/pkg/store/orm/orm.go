package orm

import (
	"fmt"
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/pb"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// driverName = DriverNameMySQL
	connection *pb.StoreConnection
	ormDB      *gorm.DB
)

// Init .
func Init(storeConnection *pb.StoreConnection) (err error) {
	glog.Infoln(storeConnection)

	connection = storeConnection

	var dialector gorm.Dialector
	switch dialector {
	default:
		dataSourceName := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8",
			connection.User,
			connection.Password,
			connection.Host,
			connection.Port,
			connection.Schema)

		dialector = mysql.Open(dataSourceName)
	}

	ormDB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ormDB.AutoMigrate(&PlatformProvider{})

	// initModels()

	return nil
}

// SetMaxIdleConns .
func SetMaxIdleConns(n int) {
	if ormDB != nil {
		db, err := ormDB.DB()
		if err != nil {
			db.SetMaxIdleConns(n)
		}
	}
}

// SetMaxOpenConns .
func SetMaxOpenConns(n int) {
	if ormDB != nil {
		db, err := ormDB.DB()
		if err != nil {
			db.SetMaxOpenConns(n)
		}
	}
}
