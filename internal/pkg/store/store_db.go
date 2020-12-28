package store

import (
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/store/db"
	"time"
)

// StorerDB .
type StorerDB struct {
	lastCheckCreateTable int64
}

func (storer *StorerDB) init(storeConnection *pb.StoreConnection) (err error) {
	err = db.Init(storeConnection)
	db.SetMaxIdleConns(int(storeConnection.MaxConnection))
	db.SetMaxOpenConns(int(storeConnection.MaxConnection))
	return
}

func (storer *StorerDB) tick() (err error) {
	if time.Since(time.Unix(storer.lastCheckCreateTable, 0)) >= time.Hour {
		storer.lastCheckCreateTable = time.Now().Unix()
		return db.CreateTable()
	}
	return
}

func (storer *StorerDB) getAllServerInfo() ([]*pb.ServerInfo, error) {
	serverInfoList := make([]*pb.ServerInfo, 0)

	storeServerInfoList, err := db.ServerInfoGetAll()

	if err != nil {
		return serverInfoList, err
	}

	for _, storeServerInfo := range storeServerInfoList {
		serverInfo := &pb.ServerInfo{}
		db.Copy(serverInfo, storeServerInfo)
		serverInfoList = append(serverInfoList, serverInfo)
	}

	return serverInfoList, nil
}

func (storer *StorerDB) getAllPlatformProvider() ([]*pb.PlatformProvider, error) {
	platformProviderList := make([]*pb.PlatformProvider, 0)

	storePlatformProviderList, err := db.PlatformProviderGetAll()

	if err != nil {
		return platformProviderList, err
	}

	for _, storePlatformProvider := range storePlatformProviderList {
		platformProvider := &pb.PlatformProvider{}
		db.Copy(platformProvider, storePlatformProvider)

		platformProvider.AesKey = storePlatformProvider.AESKey.String
		platformProvider.AesIv = storePlatformProvider.AESIV.String
		platformProvider.ApiUrlBase = storePlatformProvider.APIURLBase.String

		platformProvider.Auth = &pb.Auth{
			Type:      pb.AuthType(storePlatformProvider.AuthType.Int32),
			Id:        storePlatformProvider.AuthID.String,
			Secret:    storePlatformProvider.AuthSecret.String,
			GrantType: storePlatformProvider.AuthGrantType.String,
			Scope:     storePlatformProvider.AuthScope.String,
		}

		platformProviderList = append(platformProviderList, platformProvider)
	}

	return platformProviderList, nil
}

func (storer *StorerDB) getAllCurrency() ([]*pb.Currency, error) {
	currencyList := make([]*pb.Currency, 0)

	storeCurrencyList, err := db.CurrencyGetAll()

	if err != nil {
		return currencyList, err
	}

	for _, storeCurrency := range storeCurrencyList {
		platformProvider := &pb.Currency{}
		db.Copy(platformProvider, storeCurrency)
		currencyList = append(currencyList, platformProvider)
	}

	return currencyList, nil
}

func (storer *StorerDB) addTransactionLog(transaction *pb.Transaction) error {
	return db.TransactionAddLog(transaction)
}
