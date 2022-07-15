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

// DynaQ is the base dynamic querier
type DynaQ struct {
	db         *sql.DB
	conn       *sql.Conn
	timeFormat string
	dbService  dbServ.DataBase
	conService conServ.Connection
}

// NewDynaQ returns a new dynamic querier with the provided configurations. The user must pass in a created database connection, but can optionally configure the time format for dynaQ to use whenver
// analyzing column values to determine strings vs times by passing in the desired time format or by either passing in an empty string or importing dynaQ/pkg/resources/timeFmt and passing in
// timeFmt.DefaultTimeFormat
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

// NewDqConn allows your dynamic querier to make dynamic queries on a specific database connection
func (dq DynaQ) NewDqConn(con *sql.Conn) DynaQ {
	conService := conServ.NewConnectionService()

	dq.conn = con
	dq.conService = conService.(conServ.Connection)

	return dq
}

// DatabaseQuery takes the query the user wishes to execute, along with any arguments required as arguments and executes the query against the database instance. It then returns a
// models.ResultRows holding all the rows of the result set returned by the query executed
func (dq DynaQ) DatabaseQuery(query string, args ...interface{}) (models.ResultRows, error) {
	var dud []models.Row
	rows := models.ResultRows{
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

// DatabaseQueryContext takes the specific context, query the user wishes to execute, and any arguments required as arguments and executes the query against a database instance. It then returns a
// models.ResultRows holding all the rows of the result set returned by the query executed
func (dq DynaQ) DatabaseQueryContext(ctx context.Context, query string, args ...interface{}) (models.ResultRows, error) {
	var dud []models.Row
	rows := models.ResultRows{
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

// ConnectionQueryContext takes the specific context, query the user wishes to execute, and any arguments required as arguments and executes the query against the specific database connection.
// It then returns a models.ResultRows holding all the rows of the result set returned by the query executed
func (dq DynaQ) ConnectionQueryContext(ctx context.Context, query string, args ...interface{}) (models.ResultRows, error) {
	var dud []models.Row
	rows := models.ResultRows{
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
