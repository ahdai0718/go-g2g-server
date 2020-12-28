package db

import "ohdada/g2gserver/internal/pkg/store/db/entity"

// CurrencyGetAll .
func CurrencyGetAll() ([]entity.Currency, error) {
	return currencyModel.GetAll()
}

// CurrencyGetByName .
func CurrencyGetByName(name string) (entity.Currency, error) {
	return currencyModel.GetByName(name)
}
