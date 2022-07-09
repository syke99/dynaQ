package internal

import (
	"database/sql"
	"fmt"
	"github.com/syke99/dynaQ/pkg/models"
	"reflect"
)

func UnmarshalRow(res *models.Result, rows *sql.Rows) (map[string]models.QueryValue, error) {
	// grab the column names from the result to later create an entry for each in result.Rows
	columnNames, _ := rows.Columns()

	// make a dummy interface to scan each column’s value into
	var dummyColumnValue interface{}

	// make a dummy Result to correctly initialize Columns
	dummyRes := models.QueryValue{}

	// this will keep the column names and column values synchronized to make assigning map entry values a breeze
	for i, columnName := range columnNames {
		// append an empty value to the slices so that
		// we can synchronize the column names to each
		// their corresponding column by assigning
		// the column name to the slice the corresponding
		// index
		res.ColumnTypes = append(res.ColumnTypes, "")
		res.ColumnNames = append(res.ColumnNames, "")
		res.ColumnValues = append(res.ColumnValues, dummyColumnValue)

		res.ColumnTypes[i] = ""
		res.ColumnNames[i] = columnName
		res.Columns[columnName] = dummyRes
	}

	if rows.Next() {
		// scans all values into a slice of interfaces of any size
		err := rows.Scan(res.ColumnValues)
		if err != nil {
			return res.Columns, err
		}

		// loop through the columnValues and assign them to the correct map entry in rslt.columns using the index of the value in rslt.columnValues, which was synchronized with rslt.columnNames above
		for i, value := range res.ColumnValues {
			currentColumnName := res.ColumnNames[i]
			qr := models.QueryValue{
				Type:  fmt.Sprintf("%v", reflect.ValueOf(&value).Kind()),
				Value: value,
			}
			res.Columns[currentColumnName] = qr
		}
	}

	return res.Columns, nil
}

func UnmarshalRows(res *models.Result, rows *sql.Rows, columnTypesSlice []string) ([]map[string]models.QueryValue, error) {
	var results []map[string]models.QueryValue

	// grab the column names from the result to later create an entry for each in result.Rows
	columnNames, _ := rows.Columns()

	// make a dummy interface to scan each column’s value into
	var dummyColumnValue interface{}

	// make a dummy Result to correctly initialize Columns
	dummyRes := models.QueryValue{}

	// this will keep the column names and column values synchronized to make assigning map entry values a breeze
	for i, columnName := range columnNames {
		// append an empty value to the slices so that
		// we can synchronize the column names to each
		// their corresponding column by assigning
		// the column name to the slice the corresponding
		// index
		res.ColumnTypes = append(res.ColumnTypes, "")
		res.ColumnNames = append(res.ColumnNames, "")
		res.ColumnValues = append(res.ColumnValues, dummyColumnValue)

		res.ColumnTypes[i] = ""
		res.ColumnNames[i] = columnName
		res.Columns[columnName] = dummyRes
	}

	for rows.Next() {
		// scans all values into a slice of interfaces of any size
		err := rows.Scan(res.ColumnValues...)
		if err != nil {
			return results, err
		}

		// loop through the columnValues and assign them to the correct
		// map entry in rslt.columns using the index of the value in
		// rslt.columnValues, which was synchronized with rslt.columnNames above
		for i, value := range res.ColumnValues {
			if (i + 1) <= len(res.ColumnValues) {
				columnTypesSlice[i] = fmt.Sprintf("%v", reflect.ValueOf(&value).Kind())
			}
			currentColumnName := res.ColumnNames[i]
			currentColumnType := columnTypesSlice[i]
			qr := models.QueryValue{
				Type:  currentColumnType,
				Value: value,
			}
			res.Columns[currentColumnName] = qr
		}

		results = append(results, res.Columns)
	}

	return results, nil
}
