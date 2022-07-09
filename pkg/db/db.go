package db

import (
	"context"
	"database/sql"
	"github.com/syke99/dynaQ/internal"
	models2 "github.com/syke99/dynaQ/pkg/models"
)

type DataBase struct{}

type service interface {
	Query(db *sql.DB, query string, queryParams ...interface{}) ([]map[string]models2.QueryValue, error)
	QueryWithContext(db *sql.DB, ctx context.Context, query string, queryParams ...interface{}) ([]map[string]models2.QueryValue, error)
	QueryRow(db *sql.DB, query string, queryParams ...interface{}) (map[string]models2.QueryValue, error)
	QueryRowWithContext(db *sql.DB, ctx context.Context, query string, queryParams ...interface{}) (map[string]models2.QueryValue, error)
}

func NewDbService() service {
	return DataBase{}
}

func (db DataBase) Query(dBase *sql.DB, query string, queryParams ...interface{}) ([]map[string]models2.QueryValue, error) {
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
	res, err := dBase.Query(query, queryParams)
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

func (db DataBase) QueryWithContext(dBase *sql.DB, ctx context.Context, query string, queryParams ...interface{}) ([]map[string]models2.QueryValue, error) {
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
	res, err := dBase.QueryContext(ctx, query, queryParams)
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

func (db DataBase) QueryRow(dBase *sql.DB, query string, queryParams ...interface{}) (map[string]models2.QueryValue, error) {
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
	res, err := dBase.Query(query, queryParams)
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

func (db DataBase) QueryRowWithContext(dBase *sql.DB, ctx context.Context, query string, queryParams ...interface{}) (map[string]models2.QueryValue, error) {
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
	res, err := dBase.QueryContext(ctx, query, queryParams...)
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
