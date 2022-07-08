package dq

import (
	"context"
	"database/sql"
	"github.com/syke99/go-dq/pkg/conn"
	dbase "github.com/syke99/go-dq/pkg/db"
	"github.com/syke99/go-dq/pkg/stmnt"
	"github.com/syke99/go-dq/pkg/tx"
)

type Dq struct {
	db           *sql.DB
	stmnt        *sql.Stmt
	tx           *sql.Tx
	con          *sql.Conn
	dbService    dbase.DataBase
	stmntService stmnt.Statement
	txService    tx.Transaction
	conService   conn.Connection
}

func NewDq(db *sql.DB) Dq {
	dbService := dbase.NewDbService(db)
	stmntService := stmnt.NewPreparedStatementService()
	txService := tx.NewTransactionService()
	conService := conn.NewConnectionService()

	return Dq{
		db:           db,
		stmnt:        &sql.Stmt{},
		tx:           &sql.Tx{},
		con:          &sql.Conn{},
		dbService:    dbService.(dbase.DataBase),
		stmntService: stmntService.(stmnt.Statement),
		txService:    txService.(tx.Transaction),
		conService:   conService.(conn.Connection),
	}
}

func (dq Dq) NewPreparedStatement(query string) (Dq, error) {
	stm, err := dq.db.Prepare(query)
	if err != nil {
		return dq, err
	}

	dq.stmnt = stm

	return dq, nil
}

func (dq Dq) NewPreparedStatementWithContext(ctx context.Context, query string) (Dq, error) {
	stm, err := dq.db.PrepareContext(ctx, query)
	if err != nil {
		return dq, err
	}

	dq.stmnt = stm

	return dq, nil
}

func (dq Dq) NewTransaction(tx *sql.Tx) Dq {
	dq.tx = tx

	return dq
}
