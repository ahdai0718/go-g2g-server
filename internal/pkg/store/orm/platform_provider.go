package orm

import (
	"gorm.io/gorm"
)

// PlatformProvider .
type PlatformProvider struct {
	gorm.Model
	Name          string
	FactoryName   string
	AESKey        string
	AESIV         string
	AuthID        string
	AuthSecret    string
	AuthGrantType string
	AuthScope     string
	APIURLBase    string
	AuthType      uint
}

// GetAll .
func GetAll() (platformProviderList []PlatformProvider, err error) {
	result := ormDB.Find(&platformProviderList)
	err = result.Error
	if err != nil {
		return
	}
	return
}

// GetByName .
func GetByName(name string) (platformProvider PlatformProvider, err error) {
	result := ormDB.Where("name = ?", name).First(&platformProvider)
	err = result.Error
	if err != nil {
		return
	}
	return
}
