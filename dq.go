package dynaQ

import (
	"context"
	"database/sql"
	"github.com/syke99/dynaQ/internal"
	"github.com/syke99/dynaQ/pkg/resources/models"
	"github.com/syke99/dynaQ/pkg/resources/timeFmt"
	conServ "github.com/syke99/dynaQ/pkg/services/conn"
	dbServ "github.com/syke99/dynaQ/pkg/services/db"
)

type DynaQ struct {
	db         *sql.DB
	stmnt      *sql.Stmt
	tx         *sql.Tx
	conn       *sql.Conn
	timeFormat string
	dbService  dbServ.DataBase
	conService conServ.Connection
}

func NewDynaQ(db *sql.DB, timeFormat string) DynaQ {
	dbService := dbServ.NewDbService()
	var tmFmt string

	if timeFormat == "" {
		tmFmt = timeFmt.DefaultTimeFormat
	} else {
		tmFmt = timeFormat
	}

	return DynaQ{
		db:         db,
		timeFormat: tmFmt,
		dbService:  dbService.(dbServ.DataBase),
		conService: conServ.Connection{},
	}
}

func (dq DynaQ) NewDqConn(con *sql.Conn) DynaQ {
	conService := conServ.NewConnectionService()

	dq.conn = con
	dq.conService = conService.(conServ.Connection)

	return dq
}

func (dq DynaQ) DatabaseQuery(query string, args ...interface{}) (models.MultiRowResult, error) {
	var dud []models.Row
	rows := models.MultiRowResult{
		CurrentRow: 1,
		Results:    dud,
	}

	queryArgs := internal.QueryArgs{Args: args}

	r, err := dq.dbService.Query(dq.db, query, dq.timeFormat, queryArgs)
	if err != nil {
		return rows, err
	}

	rows.Results = r

	return rows, nil
}

func (dq DynaQ) DatabaseQueryContext(ctx context.Context, query string, args ...interface{}) (models.MultiRowResult, error) {
	var dud []models.Row
	rows := models.MultiRowResult{
		CurrentRow: 1,
		Results:    dud,
	}

	queryArgs := internal.QueryArgs{Args: args}

	r, err := dq.dbService.QueryWithContext(dq.db, ctx, query, dq.timeFormat, queryArgs)
	if err != nil {
		return rows, err
	}

	rows.Results = r

	return rows, nil
}

func (dq DynaQ) ConnectionQueryContext(ctx context.Context, query string, args ...interface{}) (models.MultiRowResult, error) {
	var dud []models.Row
	rows := models.MultiRowResult{
		CurrentRow: 1,
		Results:    dud,
	}

	queryArgs := internal.QueryArgs{Args: args}

	r, err := dq.conService.QueryWithContext(dq.conn, ctx, query, dq.timeFormat, queryArgs)
	if err != nil {
		return rows, err
	}

	rows.Results = r

	return rows, nil
}
