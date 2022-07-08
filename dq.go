package dq

import (
	"database/sql"

	dbase "github.com/syke99/go-dq/db"
	"github.com/syke99/go-dq/stmnt"
)

type Dq struct {
	db           *sql.DB
	dbService    dbase.DataBase
	stmnt        *sql.Stmt
	stmntService stmnt.Statement
	tx           *sql.Tx
}

func NewDq(db *sql.DB) Dq {
	dbService := dbase.NewDbService(db)
	stmntService := stmnt.NewPreparedStatementService()

	return Dq{
		db:           db,
		dbService:    dbService.(dbase.DataBase),
		stmnt:        &sql.Stmt{},
		stmntService: stmntService.(stmnt.Statement),
		tx:           &sql.Tx{},
	}
}

// func (dq Dq) NewPreparedStatement(query string) (Dq, error) {
// 	stmnt, err := dq.db.Prepare(query)
// 	if err != nil {
// 		return dq, err
// 	}

// 	dq.stmnt = stmnt

// 	return dq, nil
// }

// func (dq Dq) NewPreparedStatementWithContext(ctx context.Context, query string) (Dq, error) {
// 	stmnt, err := dq.db.PrepareContext(ctx, query)
// 	if err != nil {
// 		return dq, err
// 	}

// 	dq.stmnt = stmnt

// 	return dq, nil
// }

func (dq Dq) NewTransaction(tx *sql.Tx) Dq {
	dq.tx = tx

	return dq
}
