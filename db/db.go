package db

import (
	"context"
	"database/sql"

	"github.com/syke99/go-dq/internal"
)

type DataBase struct {
	db *sql.DB
}

type service interface {
	Query(query string, queryParams ...interface{}) ([]map[string]interface{}, error)
	QueryWithContext(ctx context.Context, query string, queryParams ...interface{}) ([]map[string]interface{}, error)
	QueryRow(query string, queryParams ...interface{}) (map[string]interface{}, error)
	QueryRowWithContext(ctx context.Context, query string, queryParams ...interface{}) (map[string]interface{}, error)
}

func NewDbService(db *sql.DB) service {
	return DataBase{
		db: db,
	}
}

func (db DataBase) Query(query string, queryParams ...interface{}) ([]map[string]interface{}, error) {

	var results []map[string]interface{}

	var columnMap map[string]interface{}
	var columnValuesSlice []interface{}
	var columnNamesSlice []string

	rslt := internal.Result{
		Columns:      columnMap,
		ColumnValues: columnValuesSlice,
		ColumnNames:  columnNamesSlice,
	}

	// query the db with the dynamic query and it’s params
	res, err := db.db.Query(query, queryParams)
	if err != nil {
		return results, err
	}

	defer res.Close()

	// grab the column names from the result to later create an entry for each in result.Rows
	columnNames, _ := res.Columns()

	// make a dummy interface to scan each column’s value into
	var dummyColumnValue interface{}

	// this will keep the column names and column values synchronized to make assigning map entry values a breeze
	for i, columnName := range columnNames {
		rslt.ColumnNames[i] = columnName
		rslt.ColumnValues[i] = dummyColumnValue
		rslt.Columns[columnName] = dummyColumnValue
	}

	for res.Next() {
		// scans all values into a slice of interfaces of any size
		res.Scan(rslt.ColumnValues)

		// loop through the columnValues and assign them to the correct map entry in rslt.columns using the index of the value in rslt.columnValues, which was synchronized with rslt.columnNames above
		for i, value := range rslt.ColumnValues {
			currentColumnName := rslt.ColumnNames[i]
			rslt.Columns[currentColumnName] = value
		}

		results = append(results, rslt.Columns)
	}

	return results, nil
}

func (db DataBase) QueryWithContext(ctx context.Context, query string, queryParams ...interface{}) ([]map[string]interface{}, error) {

	var results []map[string]interface{}

	var columnMap map[string]interface{}
	var columnValuesSlice []interface{}
	var columnNamesSlice []string

	rslt := internal.Result{
		Columns:      columnMap,
		ColumnValues: columnValuesSlice,
		ColumnNames:  columnNamesSlice,
	}

	// query the db with the dynamic query and it’s params
	res, err := db.db.QueryContext(ctx, query, queryParams)
	if err != nil {
		return results, err
	}

	defer res.Close()

	// grab the column names from the result to later create an entry for each in result.Rows
	columnNames, _ := res.Columns()

	// make a dummy interface to scan each column’s value into
	var dummyColumnValue interface{}

	// this will keep the column names and column values synchronized to make assigning map entry values a breeze
	for i, columnName := range columnNames {
		rslt.ColumnNames[i] = columnName
		rslt.ColumnValues[i] = dummyColumnValue
		rslt.Columns[columnName] = dummyColumnValue
	}

	for res.Next() {
		// scans all values into a slice of interfaces of any size
		res.Scan(rslt.ColumnValues)

		// loop through the columnValues and assign them to the correct map entry in rslt.columns using the index of the value in rslt.columnValues, which was synchronized with rslt.columnNames above
		for i, value := range rslt.ColumnValues {
			currentColumnName := rslt.ColumnNames[i]
			rslt.Columns[currentColumnName] = value
		}

		results = append(results, rslt.Columns)
	}

	return results, nil
}

func (db DataBase) QueryRow(query string, queryParams ...interface{}) (map[string]interface{}, error) {
	var columnMap map[string]interface{}
	var columnValuesSlice []interface{}
	var columnNamesSlice []string

	rslt := internal.Result{
		Columns:      columnMap,
		ColumnValues: columnValuesSlice,
		ColumnNames:  columnNamesSlice,
	}

	// query the db with the dynamic query and it’s params
	res, err := db.db.Query(query, queryParams)
	if err != nil {
		return rslt.Columns, err
	}

	defer res.Close()

	// grab the column names from the result to later create an entry for each in result.Rows
	columnNames, _ := res.Columns()

	// make a dummy interface to scan each column’s value into
	var dummyColumnValue interface{}

	// this will keep the column names and column values synchronized to make assigning map entry values a breeze
	for i, columnName := range columnNames {
		rslt.ColumnNames[i] = columnName
		rslt.ColumnValues[i] = dummyColumnValue
		rslt.Columns[columnName] = dummyColumnValue
	}

	if res.Next() {
		// scans all values into a slice of interfaces of any size
		res.Scan(rslt.ColumnValues)

		// loop through the columnValues and assign them to the correct map entry in rslt.columns using the index of the value in rslt.columnValues, which was synchronized with rslt.columnNames above
		for i, value := range rslt.ColumnValues {
			currentColumnName := rslt.ColumnNames[i]
			rslt.Columns[currentColumnName] = value
		}
	}

	return rslt.Columns, nil
}

func (db DataBase) QueryRowWithContext(ctx context.Context, query string, queryParams ...interface{}) (map[string]interface{}, error) {
	var columnMap map[string]interface{}
	var columnValuesSlice []interface{}
	var columnNamesSlice []string

	rslt := internal.Result{
		Columns:      columnMap,
		ColumnValues: columnValuesSlice,
		ColumnNames:  columnNamesSlice,
	}

	// query the db with the dynamic query and it’s params
	res, err := db.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return rslt.Columns, err
	}

	// grab the column names from the result to later create an entry for each in result.Rows
	columnNames, _ := res.Columns()

	// make a dummy interface to scan each column’s value into
	var dummyColumnValue interface{}

	// this will keep the column names and column values synchronized to make assigning map entry values a breeze
	for i, columnName := range columnNames {
		rslt.ColumnNames[i] = columnName
		rslt.ColumnValues[i] = dummyColumnValue
		rslt.Columns[columnName] = dummyColumnValue
	}

	if res.Next() {
		// scans all values into a slice of interfaces of any size
		res.Scan(rslt.ColumnValues)

		// loop through the columnValues and assign them to the correct map entry in rslt.columns using the index of the value in rslt.columnValues, which was synchronized with rslt.columnNames above
		for i, value := range rslt.ColumnValues {
			currentColumnName := rslt.ColumnNames[i]
			rslt.Columns[currentColumnName] = value
		}
	}
	return rslt.Columns, nil
}
