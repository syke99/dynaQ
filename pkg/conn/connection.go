package conn

import (
	"context"
	"database/sql"
	"github.com/syke99/dynaQ/internal"
	models2 "github.com/syke99/dynaQ/pkg/models"
)

type Connection struct{}

type service interface {
	QueryWithContext(conn *sql.Conn, ctx context.Context, query string, queryParams ...interface{}) ([]map[string]models2.QueryValue, error)
	QueryRowWithContext(conn *sql.Conn, ctx context.Context, query string, queryParams ...interface{}) (map[string]models2.QueryValue, error)
}

func NewConnectionService(conn *sql.Conn) service {
	return Connection{}
}

func (db Connection) QueryWithContext(conn *sql.Conn, ctx context.Context, query string, queryParams ...interface{}) ([]map[string]models2.QueryValue, error) {

	var results []map[string]models2.QueryValue

	var columnMap map[string]models2.QueryValue
	var columnValuesSlice []interface{}
	var columnNamesSlice []string
	var columnTypesSlice []string

	rslt := models2.Result{
		Columns:      columnMap,
		ColumnValues: columnValuesSlice,
		ColumnNames:  columnNamesSlice,
		ColumnTypes:  columnTypesSlice,
	}

	// query the db with the dynamic query and it’s params
	res, err := conn.QueryContext(ctx, query, queryParams)
	if err != nil {
		return results, err
	}

	defer res.Close()

	if err != nil {
		var dummyResults []map[string]models2.QueryValue

		return dummyResults, err
	}

	defer res.Close()

	unmarshalled, err := internal.UnmarshalRows(&rslt, res, columnTypesSlice)
	if err != nil {
		var dummyResults []map[string]models2.QueryValue

		return dummyResults, err
	}

	return unmarshalled, nil
}

func (db Connection) QueryRowWithContext(conn *sql.Conn, ctx context.Context, query string, queryParams ...interface{}) (map[string]models2.QueryValue, error) {
	var columnMap map[string]models2.QueryValue
	var columnValuesSlice []interface{}
	var columnNamesSlice []string
	var columnTypesSlice []string

	rslt := models2.Result{
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

	unmarshalled, err := internal.UnmarshalRow(&rslt, res)
	if err != nil {
		return rslt.Columns, err
	}

	return unmarshalled, nil
}
