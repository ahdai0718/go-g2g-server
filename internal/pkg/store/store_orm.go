package store

import (
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/store/orm"
)

// StorerORM .
type StorerORM struct {
	lastCheckCreateTable int64
}

func (storer *StorerORM) init(storeConnection *pb.StoreConnection) (err error) {
	err = orm.Init(storeConnection)
	orm.SetMaxIdleConns(int(storeConnection.MaxConnection))
	orm.SetMaxOpenConns(int(storeConnection.MaxConnection))
	return
}

func (storer *StorerORM) GetPlatformProviderByName(name string) (platformProvider *pb.PlatformProvider, err error) {
	dbPlatformProvider, err := orm.GetByName(name)

	platformProvider = &pb.PlatformProvider{
		FactoryName: dbPlatformProvider.FactoryName,
		Name:        dbPlatformProvider.Name,
		AesKey:      dbPlatformProvider.AESKey,
		AesIv:       dbPlatformProvider.AESIV,
		ApiUrlBase:  dbPlatformProvider.APIURLBase,
		// PublicIpAddress: dbPlatformProvider,
		// RunMode:         dbPlatformProvider,
		Auth: &pb.Auth{
			// 			state         protoimpl.MessageState
			// sizeCache     protoimpl.SizeCache
			// unknownFields protoimpl.UnknownFields
			Type:      pb.AuthType(dbPlatformProvider.AuthType),
			Id:        dbPlatformProvider.AuthID,
			Secret:    dbPlatformProvider.AuthSecret,
			GrantType: dbPlatformProvider.AuthGrantType,
			Scope:     dbPlatformProvider.AuthScope,
		},
	}
	return
}

func (storer *StorerORM) GetAllPlatformProvider() (platformProviderList []*pb.PlatformProvider, err error) {
	dbPlatformProviderList, err := orm.GetAll()
	for _, pp := range dbPlatformProviderList {
		platformProviderList = append(platformProviderList, &pb.PlatformProvider{
			FactoryName: pp.FactoryName,
			Name:        pp.Name,
			AesKey:      pp.AESKey,
			AesIv:       pp.AESIV,
			ApiUrlBase:  pp.APIURLBase,
			// PublicIpAddress: pp,
			// RunMode:         pp,
			Auth: &pb.Auth{
				// 			state         protoimpl.MessageState
				// sizeCache     protoimpl.SizeCache
				// unknownFields protoimpl.UnknownFields
				// Type:      pb.AuthType(pp.AuthType),
				Id:        pp.AuthID,
				Secret:    pp.AuthSecret,
				GrantType: pp.AuthGrantType,
				Scope:     pp.AuthScope,
			},
		})
	}
	return
}
