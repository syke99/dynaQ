package tx

import (
	"context"
	"database/sql"
	"github.com/syke99/dynaQ/internal"
	"github.com/syke99/dynaQ/pkg/models"
)

type Transaction struct{}

type service interface {
	Query(tx *sql.Tx, query string, queryParams ...interface{}) ([]map[string]models.QueryValue, error)
	QueryWithContext(tx *sql.Tx, ctx context.Context, query string, queryParams ...interface{}) ([]map[string]models.QueryValue, error)
	QueryRow(tx *sql.Tx, query string, queryParams ...interface{}) (map[string]models.QueryValue, error)
	QueryRowWithContext(tx *sql.Tx, ctx context.Context, query string, queryParams ...interface{}) (map[string]models.QueryValue, error)
}

func NewTransactionService() service {
	return Transaction{}
}

func (t Transaction) Query(tx *sql.Tx, query string, queryParams ...interface{}) ([]map[string]models.QueryValue, error) {
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
	res, err := tx.Query(query, queryParams)
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

func (t Transaction) QueryWithContext(tx *sql.Tx, ctx context.Context, query string, queryParams ...interface{}) ([]map[string]models.QueryValue, error) {
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
	res, err := tx.QueryContext(ctx, query, queryParams)
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

func (t Transaction) QueryRow(tx *sql.Tx, query string, queryParams ...interface{}) (map[string]models.QueryValue, error) {
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
	res, err := tx.Query(query, queryParams)
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

func (t Transaction) QueryRowWithContext(tx *sql.Tx, ctx context.Context, query string, queryParams ...interface{}) (map[string]models.QueryValue, error) {
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
	res, err := tx.QueryContext(ctx, query, queryParams...)
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
