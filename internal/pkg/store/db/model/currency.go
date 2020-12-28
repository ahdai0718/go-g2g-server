package model

import (
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/pb"
	"ohdada/g2gserver/internal/pkg/store/db/entity"

	"database/sql"
)

const (
	_ = iota
	stmtCurrencyGetAll
	stmtCurrencyGetByName
)

// NewCurrency .
func NewCurrency(db *sql.DB, connection *pb.StoreConnection) *Currency {
	model := &Currency{}
	model.db = db
	model.connection = connection
	model.stmtMap = make(map[int]*sql.Stmt)

	var err error

	if _, err = model.genStmt(stmtCurrencyGetAll, model.genSPStrWithSchema("Currency_SP_GetAll()")); err != nil {
		glog.Error(err)
	}

	if _, err = model.genStmt(stmtCurrencyGetByName, model.genSPStrWithSchema("Currency_SP_GetByName(?)")); err != nil {
		glog.Error(err)
	}

	return model
}

// Currency .
type Currency struct {
	base
}

// GetAll .
func (model *Currency) GetAll() ([]entity.Currency, error) {
	stmt := model.getStmt(stmtCurrencyGetAll)
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

	CurrencyList := make([]entity.Currency, 0)

	for rows.Next() {
		Currency := entity.Currency{}
		err = rows.Scan(
			&Currency.Name,
			&Currency.MultiplyToCredit)

		if err != nil {
			glog.Error(err)
			return nil, err
		}

		CurrencyList = append(CurrencyList, Currency)
	}

	return CurrencyList, nil
}

// GetByName .
func (model *Currency) GetByName(name string) (entity.Currency, error) {
	Currency := entity.Currency{}

	var err error

	stmt := model.getStmt(stmtCurrencyGetByName)

	err = stmt.QueryRow(name).Scan(
		&Currency.Name,
		&Currency.MultiplyToCredit)

	if err != nil {
		glog.Error(err)
		return Currency, err
	}

	return Currency, nil
}
