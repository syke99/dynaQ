package db

import (
	"context"
	"database/sql"
	"github.com/syke99/dynaQ/internal"
	"github.com/syke99/dynaQ/pkg/resources/models"
)

type DataBase struct{}

type service interface {
	Query(db *sql.DB, query string, timeFormat string, queryParams ...interface{}) ([]models.Row, error)
	QueryWithContext(db *sql.DB, ctx context.Context, query string, timeFormat string, queryParams ...interface{}) ([]models.Row, error)
	QueryRow(db *sql.DB, query string, timeFormat string, queryParams ...interface{}) (models.Row, error)
	QueryRowWithContext(db *sql.DB, ctx context.Context, query string, timeFormat string, queryParams ...interface{}) (models.Row, error)
}

func NewDbService() service {
	return DataBase{}
}

func (db DataBase) Query(dBase *sql.DB, query string, timeFormat string, queryParams ...interface{}) ([]models.Row, error) {
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
	res, err := dBase.Query(query, queryParams...)
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

func (db DataBase) QueryWithContext(dBase *sql.DB, ctx context.Context, query string, timeFormat string, queryParams ...interface{}) ([]models.Row, error) {
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
	res, err := dBase.QueryContext(ctx, query, queryParams...)
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

func (db DataBase) QueryRow(dBase *sql.DB, query string, timeFormat string, queryParams ...interface{}) (models.Row, error) {
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
	res, err := dBase.Query(query, queryParams...)
	if err != nil {
		var dummyResults models.Row

		return dummyResults, err
	}

	defer res.Close()

	unmarshalled, err := internal.UnmarshalRow(&rslt, res, timeFormat)
	if err != nil {
		var dummyResults models.Row

		return dummyResults, err
	}

	return unmarshalled, nil
}

func (db DataBase) QueryRowWithContext(dBase *sql.DB, ctx context.Context, query string, timeFormat string, queryParams ...interface{}) (models.Row, error) {
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
	res, err := dBase.QueryContext(ctx, query, queryParams...)
	if err != nil {
		var dummyResults models.Row

		return dummyResults, err
	}

	defer res.Close()

	unmarshalled, err := internal.UnmarshalRow(&rslt, res, timeFormat)
	if err != nil {
		var dummyResults models.Row

		return dummyResults, err
	}

	return unmarshalled, nil
}
