package entity

import "database/sql"

// PlatformProvider .
type PlatformProvider struct {
	Name          sql.NullString
	FactoryName   sql.NullString
	AESKey        sql.NullString
	AESIV         sql.NullString
	AuthType      sql.NullInt32
	AuthID        sql.NullString
	AuthSecret    sql.NullString
	AuthGrantType sql.NullString
	AuthScope     sql.NullString
	APIURLBase    sql.NullString
}
