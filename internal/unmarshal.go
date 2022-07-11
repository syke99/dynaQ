package internal

import (
	"database/sql"
	"fmt"
	"github.com/syke99/dynaQ/pkg/resources/models"
	"reflect"
	"time"
)

func UnmarshalRow(res *models.Result, rows *sql.Rows, timeFormat string) (models.Row, error) {
	// grab the column names from the result to later create an entry for each in result.Rows
	columnNames, _ := rows.Columns()

	// make a dummy interface to scan each column’s value into
	var dummyColumnValue interface{}

	// make a dummy Result to correctly initialize Columns
	dummyRes := models.ColumnValue{}

	count := len(columnNames)
	colNames := make(map[string]models.ColumnValue, count)
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

	columns := make([]models.ColumnValue, len(values))

	rowResults := models.Row{CurrentColumn: 0, Columns: columns}

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
				colType := evalType(timeFormat, colVal, values[i])

				columnTypesSlice[i] = colType
			}
			currentColumnName := res.ColumnNames[i]
			currentColumnType := columnTypesSlice[i]

			currentColumnType = columnTypesSlice[i]

			qr := models.ColumnValue{
				Type:  currentColumnType,
				Name:  currentColumnName,
				Value: colVal,
			}

			rowResults.Columns[i] = qr
		}
	}

	return rowResults, nil
}

func UnmarshalRows(res *models.Result, rows *sql.Rows, timeFormat string) ([]models.Row, map[string][]models.ColumnValue, error) {
	var results []models.Row
	var colMap map[string][]models.ColumnValue

	// grab the column names from the result to later create an entry for each in result.Rows
	columnNames, _ := rows.Columns()

	// make a dummy interface to scan each column’s value into
	var dummyColumnValue interface{}

	// make a dummy Result to correctly initialize Columns
	dummyRes := models.ColumnValue{}

	count := len(columnNames)
	colNames := make(map[string]models.ColumnValue, count)
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
			return results, colMap, err
		}

		columns := make([]models.ColumnValue, len(values))

		rowResults := models.Row{CurrentColumn: 0, Columns: columns}

		// loop through the columnValues and assign them to the correct
		// map entry in rslt.columns using the index of the value in
		// rslt.columnValues, which was synchronized with rslt.columnNames above
		for i, _ := range valuePointers {
			colVal := evalVal(values[i])
			if i < len(valuePointers) {
				colType := evalType(timeFormat, colVal, values[i])

				columnTypesSlice[i] = colType
			}
			currentColumnName := res.ColumnNames[i]
			currentColumnType := columnTypesSlice[i]

			currentColumnType = columnTypesSlice[i]

			qr := models.ColumnValue{
				Type:  currentColumnType,
				Name:  currentColumnName,
				Value: colVal,
			}

			rowResults.Columns[i] = qr
		}

		results = append(results, rowResults)
	}

	for i, row := range results {
		colBool, col := row.NextColumn()
		for colBool {
			colSlice := make([]models.ColumnValue, len(results))
			if i == 0 {
				colMap[col.Name] = colSlice
				colMap[col.Name] = append(colMap[col.Name], col)
			}
			colMap[col.Name] = append(colMap[col.Name], col)
		}
	}

	return results, colMap, nil
}

func evalType(timeFormat string, colVal string, value interface{}) string {
	pointsTo := reflect.Indirect(reflect.ValueOf(value))

	cType := fmt.Sprintf("%v", pointsTo.Type())

	if cType == "[]uint8" {
		return "image ([]byte)"
	}

	if cType == "string" {
		_, err := time.Parse(timeFormat, colVal)
		if err != nil {
			cType = "string"
		} else {
			cType = "time.Time"
		}
	}

	return cType
}

func evalVal(value interface{}) string {
	pointsTo := reflect.Indirect(reflect.ValueOf(value))

	return fmt.Sprintf("%v", pointsTo)
}
