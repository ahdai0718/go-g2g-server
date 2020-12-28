package model

import (
	"ohdada/g2gserver/internal/pkg/pb"
	"database/sql"
	"fmt"
)

type base struct {
	db         *sql.DB
	connection *pb.StoreConnection
	stmtMap    map[int]*sql.Stmt
}

func (model *base) getStmt(id int) *sql.Stmt {
	return model.stmtMap[id]
}

func (model *base) genStmt(id int, sql string) (*sql.Stmt, error) {
	stmt, err := model.db.Prepare(sql)
	if stmt != nil {
		model.stmtMap[id] = stmt
	}
	return stmt, err
}

func (model *base) genSPStrWithSchema(sp string) string {
	return fmt.Sprintf("CALL %s.%s;", model.connection.Schema, sp)
}
