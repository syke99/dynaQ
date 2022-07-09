package stmnt

import (
	"context"
	"database/sql"
	"github.com/syke99/dynaQ/internal"
	"github.com/syke99/dynaQ/models"
)

type Statement struct{}

type service interface {
	Query(stm *sql.Stmt, query string, queryParams ...interface{}) ([]map[string]models.QueryValue, error)
	QueryWithContext(stm *sql.Stmt, ctx context.Context, query string, queryParams ...interface{}) ([]map[string]models.QueryValue, error)
	QueryRow(stm *sql.Stmt, query string, queryParams ...interface{}) (map[string]models.QueryValue, error)
	QueryRowWithContext(stm *sql.Stmt, ctx context.Context, query string) (map[string]models.QueryValue, error)
}

func NewPreparedStatementService() service {
	return Statement{}
}

func (db Statement) Query(stm *sql.Stmt, query string, queryParams ...interface{}) ([]map[string]models.QueryValue, error) {
	var results []map[string]models.QueryValue

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
	res, err := stm.Query(query, queryParams)
	if err != nil {
		return results, err
	}

	defer res.Close()

	if err != nil {
		var dummyResults []map[string]models.QueryValue

		return dummyResults, err
	}

	defer res.Close()

	unmarshalled, err := internal.UnmarshalRows(&rslt, res, columnTypesSlice)
	if err != nil {
		var dummyResults []map[string]models.QueryValue

		return dummyResults, err
	}

	return unmarshalled, nil
}

func (db Statement) QueryWithContext(stm *sql.Stmt, ctx context.Context, query string, queryParams ...interface{}) ([]map[string]models.QueryValue, error) {
	var results []map[string]models.QueryValue

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
	res, err := stm.QueryContext(ctx, query, queryParams)
	if err != nil {
		return results, err
	}

	defer res.Close()

	if err != nil {
		var dummyResults []map[string]models.QueryValue

		return dummyResults, err
	}

	defer res.Close()

	unmarshalled, err := internal.UnmarshalRows(&rslt, res, columnTypesSlice)
	if err != nil {
		var dummyResults []map[string]models.QueryValue

		return dummyResults, err
	}

	return unmarshalled, nil
}

func (db Statement) QueryRow(stm *sql.Stmt, query string, queryParams ...interface{}) (map[string]models.QueryValue, error) {
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
	res, err := stm.Query(query, queryParams)
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

func (db Statement) QueryRowWithContext(stm *sql.Stmt, ctx context.Context, query string) (map[string]models.QueryValue, error) {
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
	res, err := stm.QueryContext(ctx, query)
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
