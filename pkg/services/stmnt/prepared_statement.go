package stmnt

import (
	"context"
	"database/sql"
	"github.com/syke99/dynaQ/internal"
	"github.com/syke99/dynaQ/pkg/resources/models"
)

type Statement struct{}

type service interface {
	Query(stm *sql.Stmt, timeFormat string, queryParams ...interface{}) ([]models.Row, error)
	QueryWithContext(stm *sql.Stmt, ctx context.Context, timeFormat string, queryParams ...interface{}) ([]models.Row, error)
	QueryRow(stm *sql.Stmt, timeFormat string, queryParams ...interface{}) (models.Row, error)
	QueryRowWithContext(stm *sql.Stmt, ctx context.Context, timeFormat string, queryParams ...interface{}) (models.Row, error)
}

func NewPreparedStatementService() service {
	return Statement{}
}

func (s Statement) Query(stm *sql.Stmt, timeFormat string, queryParams ...interface{}) ([]models.Row, error) {
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
	res, err := stm.Query(queryParams...)
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

func (s Statement) QueryWithContext(stm *sql.Stmt, ctx context.Context, timeFormat string, queryParams ...interface{}) ([]models.Row, error) {
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
	res, err := stm.QueryContext(ctx, queryParams...)
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

func (s Statement) QueryRow(stm *sql.Stmt, timeFormat string, queryParams ...interface{}) (models.Row, error) {
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
	res, err := stm.Query(queryParams...)
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

func (s Statement) QueryRowWithContext(stm *sql.Stmt, ctx context.Context, timeFormat string, queryParams ...interface{}) (models.Row, error) {
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
	res, err := stm.QueryContext(ctx, queryParams...)
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
