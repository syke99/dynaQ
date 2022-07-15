package tx

import (
	"context"
	"database/sql"
	"github.com/syke99/dynaQ/internal"
	"github.com/syke99/dynaQ/pkg/resources/models"
)

type Transaction struct{}

type service interface {
	Query(tx *sql.Tx, query string, timeFormat string, queryParams internal.QueryArgs) ([]models.Row, error)
	QueryWithContext(tx *sql.Tx, ctx context.Context, query string, timeFormat string, queryParams internal.QueryArgs) ([]models.Row, error)
	QueryRow(tx *sql.Tx, query string, timeFormat string, queryParams internal.QueryArgs) ([]models.Row, error)
	QueryRowWithContext(tx *sql.Tx, ctx context.Context, query string, timeFormat string, queryParams internal.QueryArgs) ([]models.Row, error)
}

func NewTransactionService() service {
	return Transaction{}
}

func (t Transaction) Query(tx *sql.Tx, query string, timeFormat string, queryParams internal.QueryArgs) ([]models.Row, error) {
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
	res, err := tx.Query(query, queryParams.Args...)
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

func (t Transaction) QueryWithContext(tx *sql.Tx, ctx context.Context, timeFormat string, query string, queryParams internal.QueryArgs) ([]models.Row, error) {
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
	res, err := tx.QueryContext(ctx, query, queryParams.Args...)
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

func (t Transaction) QueryRow(tx *sql.Tx, query string, timeFormat string, queryParams internal.QueryArgs) ([]models.Row, error) {
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
	res, err := tx.Query(query, queryParams.Args...)
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

func (t Transaction) QueryRowWithContext(tx *sql.Tx, ctx context.Context, query string, timeFormat string, queryParams internal.QueryArgs) ([]models.Row, error) {
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
	res, err := tx.QueryContext(ctx, query, queryParams.Args...)
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
