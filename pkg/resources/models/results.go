package models

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Columns      map[string]QueryValue
	ColumnValues []interface{}
	ColumnNames  []string
	ColumnTypes  []string
}

type SingleRowResult struct {
	Result []QueryValue
}

func (s SingleRowResult) Row() []QueryValue {
	return s.Result
}

func (s SingleRowResult) Unmarshal(dest *interface{}) {
	marshaled, _ := json.Marshal(s.Row())

	json.Unmarshal(marshaled, dest)
}

type MultiRowResult struct {
	CurrentRow int
	Results    [][]QueryValue
}

func (m MultiRowResult) NextRow() (bool, []QueryValue) {
	if m.CurrentRow > len(m.Results) {
		var dud []QueryValue
		return false, dud
	}

	m.CurrentRow++
	return true, m.Results[m.CurrentRow-1]
}

func (m MultiRowResult) Unmarshal(dest *interface{}) {
	var jsonMap map[string][]QueryValue

	for i, row := range m.Results {
		jsonMap[fmt.Sprintf("result-%d", i)] = row
	}

	marshalled, _ := json.Marshal(jsonMap)

	json.Unmarshal(marshalled, dest)
}
