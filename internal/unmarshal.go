package internal

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"

	"github.com/syke99/dynaQ/pkg/models"
)

func UnmarshalRow(res *models.Result, rows *sql.Rows, timeFormat string) ([]models.QueryValue, error) {
	// grab the column names from the result to later create an entry for each in result.Rows
	columnNames, _ := rows.Columns()

	// make a dummy interface to scan each column’s value into
	var dummyColumnValue interface{}

	// make a dummy Result to correctly initialize Columns
	dummyRes := models.QueryValue{}

	count := len(columnNames)
	colNames := make(map[string]models.QueryValue, count)
	values := make([]interface{}, count)
	valuePointers := make([]interface{}, count)

	columnValuesSlice := make([]interface{}, count)
	columnNamesSlice := make([]string, count)
	columnTypesSlice := make([]string, count)

	res.Columns = colNames
	res.ColumnValues = columnValuesSlice
	res.ColumnNames = columnNamesSlice
	res.ColumnTypes = columnTypesSlice

	// this will keep the column names and column values synchronized to make assigning map entry values a breeze
	for i, columnName := range columnNames {
		// append an empty value to the slices so that
		// we can synchronize the column names to each
		// their corresponding column by assigning
		// the column name to the slice the corresponding
		// index
		res.ColumnTypes = append(res.ColumnTypes, "")
		res.ColumnNames = append(res.ColumnNames, "")
		res.ColumnValues = append(res.ColumnValues, &dummyColumnValue)

		res.ColumnTypes[i] = ""
		res.ColumnNames[i] = columnName
		res.Columns[columnName] = dummyRes
	}

	rowResults := make([]models.QueryValue, count)

	if rows.Next() {
		for i := range columnNames {
			res.ColumnNames[i] = columnNames[i]
			valuePointers[i] = &values[i]
		}
		// scans all values into a slice of interfaces of any size
		err := rows.Scan(valuePointers...)
		if err != nil {
			return rowResults, err
		}

		// loop through the columnValues and assign them to the correct map entry in rslt.columns using the index of the value in rslt.columnValues, which was synchronized with rslt.columnNames above
		for i, _ := range valuePointers {
			colVal := evalVal(values[i])
			if i < len(valuePointers) {
				colType := evalType(values[i])

				columnTypesSlice[i] = colType
			}
			currentColumnName := res.ColumnNames[i]
			currentColumnType := columnTypesSlice[i]

			if currentColumnType == "string" {
				_, err := time.Parse(timeFormat, colVal)
				if err != nil {
					println(err.Error())
					columnTypesSlice[i] = "string"
				} else {
					columnTypesSlice[i] = "time.Time"
				}
			}

			currentColumnType = columnTypesSlice[i]

			qr := models.QueryValue{
				Type:  currentColumnType,
				Name:  currentColumnName,
				Value: colVal,
			}

			rowResults[i] = qr
		}
	}

	return rowResults, nil
}

func UnmarshalRows(res *models.Result, rows *sql.Rows, timeFormat string) ([][]models.QueryValue, error) {
	var results [][]models.QueryValue

	// grab the column names from the result to later create an entry for each in result.Rows
	columnNames, _ := rows.Columns()

	// make a dummy interface to scan each column’s value into
	var dummyColumnValue interface{}

	// make a dummy Result to correctly initialize Columns
	dummyRes := models.QueryValue{}

	count := len(columnNames)
	colNames := make(map[string]models.QueryValue, count)
	values := make([]interface{}, count)
	valuePointers := make([]interface{}, count)

	columnValuesSlice := make([]interface{}, count)
	columnNamesSlice := make([]string, count)
	columnTypesSlice := make([]string, count)

	res.Columns = colNames
	res.ColumnValues = columnValuesSlice
	res.ColumnNames = columnNamesSlice
	res.ColumnTypes = columnTypesSlice

	// this will keep the column names and column values synchronized to make assigning map entry values a breeze
	for i, columnName := range columnNames {
		// append an empty value to the slices so that
		// we can synchronize the column names to each
		// their corresponding column by assigning
		// the column name to the slice the corresponding
		// index
		res.ColumnTypes = append(res.ColumnTypes, "")
		res.ColumnNames = append(res.ColumnNames, "")
		res.ColumnValues = append(res.ColumnValues, &dummyColumnValue)

		res.ColumnTypes[i] = ""
		res.ColumnNames[i] = columnName
		res.Columns[columnName] = dummyRes
	}

	for rows.Next() {
		for i := range columnNames {
			res.ColumnNames[i] = columnNames[i]
			valuePointers[i] = &values[i]
		}
		// scans all values into a slice of interfaces of any size
		err := rows.Scan(valuePointers...)
		if err != nil {
			return results, err
		}

		rowResults := make([]models.QueryValue, count)

		// loop through the columnValues and assign them to the correct
		// map entry in rslt.columns using the index of the value in
		// rslt.columnValues, which was synchronized with rslt.columnNames above
		for i, _ := range valuePointers {
			colVal := evalVal(values[i])
			if i < len(valuePointers) {
				colType := evalType(values[i])

				columnTypesSlice[i] = colType
			}
			currentColumnName := res.ColumnNames[i]
			currentColumnType := columnTypesSlice[i]

			if currentColumnType == "string" {
				_, err := time.Parse(timeFormat, colVal)
				if err != nil {
					println(err.Error())
					columnTypesSlice[i] = "string"
				} else {
					columnTypesSlice[i] = "time.Time"
				}
			}

			currentColumnType = columnTypesSlice[i]

			qr := models.QueryValue{
				Type:  currentColumnType,
				Name:  currentColumnName,
				Value: colVal,
			}

			rowResults[i] = qr
		}

		results = append(results, rowResults)
	}

	return results, nil
}

func evalType(value interface{}) string {
	test := reflect.ValueOf(value)

	pointsTo := reflect.Indirect(test)

	cType := fmt.Sprintf("%v", pointsTo.Type())

	if cType == "[]uint8" {
		return "[]byte"
	}

	return cType
}

func evalVal(value interface{}) string {
	test := reflect.ValueOf(value)

	pointsTo := reflect.Indirect(test)

	return fmt.Sprintf("%v", pointsTo)
}
