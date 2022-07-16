package internal

import (
	"context"
	"database/sql"
	"github.com/syke99/dynaQ/pkg/resources/models"
)

func DatabaseQuery(dBase *sql.DB, query string, timeFormat string, queryParams QueryArgs) ([]models.Row, error) {
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
	res, err := dBase.Query(query, queryParams.Args...)
	if err != nil {
		var dummyResults []models.Row

		return dummyResults, err
	}

	defer res.Close()

	unmarshalled := UnmarshalRows(&rslt, res, timeFormat)

	return unmarshalled, nil
}

func DatabaseQueryWithContext(dBase *sql.DB, ctx context.Context, query string, timeFormat string, queryParams QueryArgs) ([]models.Row, error) {
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
	res, err := dBase.QueryContext(ctx, query, queryParams.Args...)
	if err != nil {
		var dummyResults []models.Row

		return dummyResults, err
	}

	defer res.Close()

	unmarshalled := UnmarshalRows(&rslt, res, timeFormat)

	return unmarshalled, nil
}

func ConnectionQueryWithContext(conn *sql.Conn, ctx context.Context, query string, timeFormat string, queryParams QueryArgs) ([]models.Row, error) {
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

	unmarshalled := UnmarshalRows(&rslt, res, timeFormat)

	return unmarshalled, nil
}
