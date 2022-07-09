package conn

import (
	"context"
	"database/sql"
	"github.com/syke99/dynaQ/internal"
	"github.com/syke99/dynaQ/pkg/models"
)

type Connection struct{}

type service interface {
	QueryWithContext(conn *sql.Conn, ctx context.Context, query string, timeFormat string, queryParams ...interface{}) ([]map[string]models.QueryValue, error)
	QueryRowWithContext(conn *sql.Conn, ctx context.Context, query string, timeFormat string, queryParams ...interface{}) (map[string]models.QueryValue, error)
}

func NewConnectionService(conn *sql.Conn) service {
	return Connection{}
}

func (c Connection) QueryWithContext(conn *sql.Conn, ctx context.Context, query string, timeFormat string, queryParams ...interface{}) ([]map[string]models.QueryValue, error) {
	var columnMap map[string]models.QueryValue
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
	res, err := conn.QueryContext(ctx, query, queryParams)
	if err != nil {
		var dummyResults []map[string]models.QueryValue

		return dummyResults, err
	}

	defer res.Close()

	unmarshalled, err := internal.UnmarshalRows(&rslt, res, columnTypesSlice, timeFormat)
	if err != nil {
		var dummyResults []map[string]models.QueryValue

		return dummyResults, err
	}

	return unmarshalled, nil
}

func (c Connection) QueryRowWithContext(conn *sql.Conn, ctx context.Context, query string, timeFormat string, queryParams ...interface{}) (map[string]models.QueryValue, error) {
	var columnMap map[string]models.QueryValue
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
	res, err := conn.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return rslt.Columns, err
	}

	defer res.Close()

	unmarshalled, err := internal.UnmarshalRow(&rslt, res, columnTypesSlice, timeFormat)
	if err != nil {
		return rslt.Columns, err
	}

	return unmarshalled, nil
}
