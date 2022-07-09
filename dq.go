package dynaQ

import (
	"context"
	"database/sql"
	conServ "github.com/syke99/dynaQ/pkg/conn"
	dbServ "github.com/syke99/dynaQ/pkg/db"
	"github.com/syke99/dynaQ/pkg/models"
	stmntServ "github.com/syke99/dynaQ/pkg/stmnt"
	txserv "github.com/syke99/dynaQ/pkg/tx"
)

type DynaQ struct {
	db           *sql.DB
	stmnt        *sql.Stmt
	tx           *sql.Tx
	conn         *sql.Conn
	timeFormat   string
	dbService    dbServ.DataBase
	stmntService stmntServ.Statement
	txService    txserv.Transaction
	conService   conServ.Connection
}

func NewDynaQ(db *sql.DB, timeFormat string) DynaQ {
	dbService := dbServ.NewDbService()
	var tmFmt string

	if timeFormat == "" {
		tmFmt = DefaultTimeFormat
	} else {
		tmFmt = timeFormat
	}

	return DynaQ{
		db:           db,
		timeFormat:   tmFmt,
		dbService:    dbService.(dbServ.DataBase),
		stmntService: stmntServ.Statement{},
		txService:    txserv.Transaction{},
		conService:   conServ.Connection{},
	}
}

func (dq DynaQ) NewDqPreparedStatement(query string) (DynaQ, error) {
	stm, err := dq.db.Prepare(query)
	if err != nil {
		return dq, err
	}

	stmntService := stmntServ.NewPreparedStatementService()

	dq.stmnt = stm
	dq.stmntService = stmntService.(stmntServ.Statement)

	return dq, nil
}

func (dq DynaQ) NewDqPreparedStatementWithContext(ctx context.Context, query string) (DynaQ, error) {
	stm, err := dq.db.PrepareContext(ctx, query)
	if err != nil {
		return dq, err
	}

	stmntService := stmntServ.NewPreparedStatementService()

	dq.stmnt = stm
	dq.stmntService = stmntService.(stmntServ.Statement)

	return dq, nil
}

func (dq DynaQ) NewDqTransaction(tx *sql.Tx) DynaQ {
	txService := txserv.NewTransactionService()

	dq.tx = tx
	dq.txService = txService.(txserv.Transaction)

	return dq
}

func (dq DynaQ) NewDqConn(con *sql.Conn) DynaQ {
	conService := conServ.NewConnectionService(con)

	dq.conn = con
	dq.conService = conService.(conServ.Connection)

	return dq
}

func (dq DynaQ) DatabaseQuery(query string, args ...interface{}) ([]map[string]models.QueryValue, error) {
	return dq.dbService.Query(dq.db, query, dq.timeFormat, args)
}

func (dq DynaQ) DatabaseQueryRow(query string, args ...interface{}) (map[string]models.QueryValue, error) {
	return dq.dbService.QueryRow(dq.db, query, dq.timeFormat, args)
}

func (dq DynaQ) DatabaseQueryContext(ctx context.Context, query string, args ...interface{}) ([]map[string]models.QueryValue, error) {
	return dq.dbService.QueryWithContext(dq.db, ctx, query, dq.timeFormat, args)
}

func (dq DynaQ) DatabaseQueryRowContext(ctx context.Context, query string, args ...interface{}) (map[string]models.QueryValue, error) {
	return dq.dbService.QueryRowWithContext(dq.db, ctx, query, dq.timeFormat, args)
}

func (dq DynaQ) PreparedStatementQuery(query string, args ...interface{}) ([]map[string]models.QueryValue, error) {
	return dq.stmntService.Query(dq.stmnt, query, dq.timeFormat, args)
}

func (dq DynaQ) PreparedStatementQueryRow(query string, args ...interface{}) (map[string]models.QueryValue, error) {
	return dq.stmntService.QueryRow(dq.stmnt, query, dq.timeFormat, args)
}

func (dq DynaQ) PreparedStatementQueryContext(ctx context.Context, query string, args ...interface{}) ([]map[string]models.QueryValue, error) {
	return dq.stmntService.QueryWithContext(dq.stmnt, ctx, query, dq.timeFormat, args)
}

func (dq DynaQ) PreparedStatementQueryRowContext(ctx context.Context, query string) (map[string]models.QueryValue, error) {
	return dq.stmntService.QueryRowWithContext(dq.stmnt, ctx, query, dq.timeFormat)
}

func (dq DynaQ) TransactionQuery(query string, args ...interface{}) ([]map[string]models.QueryValue, error) {
	return dq.txService.Query(dq.tx, query, dq.timeFormat, args)
}

func (dq DynaQ) TransactionQueryRow(query string, args ...interface{}) (map[string]models.QueryValue, error) {
	return dq.txService.QueryRow(dq.tx, query, dq.timeFormat, args)
}

func (dq DynaQ) TransactionQueryContext(ctx context.Context, query string, args ...interface{}) ([]map[string]models.QueryValue, error) {
	return dq.txService.QueryWithContext(dq.tx, ctx, query, dq.timeFormat, args)
}

func (dq DynaQ) TransactionQueryRowContext(ctx context.Context, query string, args ...interface{}) (map[string]models.QueryValue, error) {
	return dq.txService.QueryRowWithContext(dq.tx, ctx, query, dq.timeFormat, args)
}

func (dq DynaQ) ConnectionQueryContext(ctx context.Context, query string, args ...interface{}) ([]map[string]models.QueryValue, error) {
	return dq.conService.QueryWithContext(dq.conn, ctx, query, dq.timeFormat, args)
}

func (dq DynaQ) ConnectionQueryRowContext(ctx context.Context, query string, args ...interface{}) (map[string]models.QueryValue, error) {
	return dq.conService.QueryRowWithContext(dq.conn, ctx, query, dq.timeFormat, args)
}
