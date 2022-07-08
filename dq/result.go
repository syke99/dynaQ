package dq

import (
	"encoding/json"
	"fmt"
	dynamicstruct "github.com/ompluscator/dynamic-struct"
	"reflect"
	"strings"
)

func UnmarshalResult(result map[string]interface{}, dest interface{}) {
	resultStruct := dynamicstruct.ExtendStruct(&dest)
	for column, value := range result {
		valueType := fmt.Sprintf("%v", reflect.TypeOf(value))
		capitalizedColumn := fmt.Sprintf("%s%s", strings.ToUpper(column[:1]), column[1:])
		tag := fmt.Sprintf("json:%s", column)

		switch valueType {
		case "int", "int8", "int16", "int32", "int64":
			resultStruct.
				AddField(capitalizedColumn, 0, tag)
		case "bool":
			resultStruct.
				AddField(capitalizedColumn, false, tag)
		case "uint", "uint8", "uint16", "uint32", "uint64":
			resultStruct.
				AddField(capitalizedColumn, 0, tag)
		case "float32", "float64":
			resultStruct.
				AddField(capitalizedColumn, 0.0, tag)
		case "string":
			resultStruct.
				AddField(capitalizedColumn, "", tag)
		}
	}
}

func UnmarshalResults(results []map[string]interface{}, dest interface{}) {
	// resultsStruct := dynamicstruct.ExtendStruct(&dest)

	type rs struct {
		rslts     []dynamicstruct.Builder
		rsltsData [][]byte
	}

	var rslts []dynamicstruct.Builder
	var rsltsData [][]byte

	res := rs{
		rslts:     rslts,
		rsltsData: rsltsData,
	}
	for _, rowMap := range results {
		resultStruct := dynamicstruct.NewStruct()
		rowJson := "{"
		for column, value := range rowMap {
			valueType := fmt.Sprintf("%v", reflect.TypeOf(value))
			valueString := fmt.Sprintf("%v", reflect.ValueOf(1).Interface())
			capitalizedColumn := fmt.Sprintf("%s%s", strings.ToUpper(column[:1]), column[1:])
			tag := fmt.Sprintf("json:%s", column)

			switch valueType {
			case "uint", "uint8", "uint16", "uint32", "uint64":
			case "int", "int8", "int16", "int32", "int64":
				resultStruct.
					AddField(capitalizedColumn, 0, tag)

				rowJson = fmt.Sprintf("%s\"%s\": %s,", rowJson, column, valueString)
			case "bool":
				resultStruct.
					AddField(capitalizedColumn, false, tag)

				rowJson = fmt.Sprintf("%s\"%s\": %s,", rowJson, column, valueString)
			case "float32", "float64":
				resultStruct.
					AddField(capitalizedColumn, 0.0, tag)

				rowJson = fmt.Sprintf("%s\"%s\": %s,", rowJson, column, valueString)
			case "string":
				resultStruct.
					AddField(capitalizedColumn, "", tag)

				rowJson = fmt.Sprintf("%s\"%s\": \"%s\",", rowJson, column, valueString)
			}
		}

		// after all rows are added, remove the trailing comma
		rowJson = rowJson[:len(rowJson)-1]

		// then add the closing brace
		rowJson = fmt.Sprintf("%s}", rowJson)

		// convert the row's data to bytes
		rowData := []byte(rowJson)

		res.rslts = append(res.rslts, resultStruct)
		res.rsltsData = append(res.rsltsData, rowData)
	}

	var resultsSlice []interface{}
	var resultsSliceData [][]byte

	for i, result := range res.rslts {
		r := result.Build().New()

		_ = json.Unmarshal(res.rsltsData[i], &r)
		data, _ := json.Marshal(r)

		resultsSlice = append(resultsSlice, r)
		resultsSliceData = append(resultsSliceData, data)
	}

	// TODO: loop through resultsSlice and create json to unmarshall until &dest with the corresponding value in resultsSliceData
}
