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

func (dq Dq) DatabaseQuery(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return dq.dbService.Query(query, args)
}

func (dq Dq) DatabaseQueryRow(query string, args ...interface{}) (map[string]interface{}, error) {
	return dq.dbService.QueryRow(query, args)
}

func (dq Dq) DatabaseQueryContext(ctx context.Context, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return dq.dbService.QueryWithContext(ctx, query, args)
}

func (dq Dq) DatabaseQueryRowContext(ctx context.Context, query string, args ...interface{}) (map[string]interface{}, error) {
	return dq.dbService.QueryRowWithContext(ctx, query, args)
}

func (dq Dq) PreparedStatementQuery(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return dq.stmntService.Query(query, args)
}

func (dq Dq) PreparedStatementQueryRow(query string, args ...interface{}) (map[string]interface{}, error) {
	return dq.stmntService.QueryRow(query, args)
}

func (dq Dq) PreparedStatementQueryContext(ctx context.Context, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return dq.stmntService.QueryWithContext(ctx, query, args)
}

func (dq Dq) PreparedStatementQueryRowContext(ctx context.Context, query string) (map[string]interface{}, error) {
	return dq.stmntService.QueryRowWithContext(ctx, query)
}

func (dq Dq) TransactionQuery(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return dq.txService.Query(query, args)
}

func (dq Dq) TransactionQueryRow(query string, args ...interface{}) (map[string]interface{}, error) {
	return dq.txService.QueryRow(query, args)
}

func (dq Dq) TransactionQueryContext(ctx context.Context, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return dq.txService.QueryWithContext(ctx, query, args)
}

func (dq Dq) TransactionQueryRowContext(ctx context.Context, query string, args ...interface{}) (map[string]interface{}, error) {
	return dq.txService.QueryRowWithContext(ctx, query, args)
}

func (dq Dq) ConnectionQueryContext(ctx context.Context, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return dq.conService.QueryWithContext(ctx, query, args)
}

func (dq Dq) ConnectionQueryRowContext(ctx context.Context, query string, args ...interface{}) (map[string]interface{}, error) {
	return dq.conService.QueryRowWithContext(ctx, query, args)
}
