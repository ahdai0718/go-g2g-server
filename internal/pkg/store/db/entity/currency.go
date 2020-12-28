package entity

import "database/sql"

// Currency .
type Currency struct {
	Name             sql.NullString
	MultiplyToCredit sql.NullFloat64
}
