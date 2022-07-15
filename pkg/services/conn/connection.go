package conn

import (
	"context"
	"database/sql"
	"github.com/syke99/dynaQ/internal"
	"github.com/syke99/dynaQ/pkg/resources/models"
)

// Connection provides a dynamic querier on a single database connection
type Connection struct{}

type service interface {
	QueryWithContext(conn *sql.Conn, ctx context.Context, query string, timeFormat string, queryParams internal.QueryArgs) ([]models.Row, error)
}

// NewConnectionService creates a new Connection with its associated service method(s)
func NewConnectionService() service {
	return Connection{}
}

// QueryWithContext is a Connection service method to execute a dynamic query, with a context, on a single database connection
func (c Connection) QueryWithContext(conn *sql.Conn, ctx context.Context, query string, timeFormat string, queryParams internal.QueryArgs) ([]models.Row, error) {
	var columnMap map[string]models.ColumnValue
	var columnValuesSlice []interface{}
	var columnNamesSlice []string
	var columnTypesSlice []string

	rslt := models.Result{
		Columns:      columnMap,
		ColumnValues: columnValuesSlice,
		ColumnNames:  columnNamesSlice,
		ColumnTypes:  columnTypesSlice,
	}

	// query the db with the dynamic query and it’s params
	res, err := conn.QueryContext(ctx, query, queryParams.Args...)
	if err != nil {
		var dummyResults []models.Row

		return dummyResults, err
	}

	defer res.Close()

	unmarshalled, err := internal.UnmarshalRows(&rslt, res, timeFormat)
	if err != nil {
		var dummyResults []models.Row

		return dummyResults, err
	}

	return unmarshalled, nil
}
