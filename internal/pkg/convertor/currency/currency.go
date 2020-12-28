package currency

import (
	"ohdada/g2gserver/internal/pkg/pb"
	"fmt"
)

var (
	defaultConvertor *Convertor
)

// DefaultConvertor .
func DefaultConvertor() *Convertor {
	if defaultConvertor == nil {
		defaultConvertor = &Convertor{}
	}
	return defaultConvertor
}

// Convertor .
type Convertor struct {
	currencyMap map[string]*pb.Currency
}

// SetCurrencyMap .
func (convertor *Convertor) SetCurrencyMap(currencyMap map[string]*pb.Currency) {
	convertor.currencyMap = currencyMap
}

// GetCurrency .
func (convertor *Convertor) GetCurrency(currencyName string) *pb.Currency {
	return convertor.currencyMap[currencyName]
}

// CurrencyToCredit .
func (convertor *Convertor) CurrencyToCredit(amount float64, currencyName string) (int64, error) {

	creidt := int64(0)

	currencyConfig, isExists := convertor.currencyMap[currencyName]

	if !isExists {
		return creidt, fmt.Errorf("currency [%s] not exists", currencyName)
	}

	creidt = int64(amount * float64(currencyConfig.MultiplyToCredit))

	return creidt, nil
}

// CreditToCurrency .
func (convertor *Convertor) CreditToCurrency(amount int64, currencyName string) (float64, error) {
	currency := float64(0)

	currencyConfig, isExists := convertor.currencyMap[currencyName]

	if !isExists {
		return currency, fmt.Errorf("currency [%s] not exists", currencyName)
	}

	currency = float64(amount) / float64(currencyConfig.MultiplyToCredit)

	return currency, nil
}
