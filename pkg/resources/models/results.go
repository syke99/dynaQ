package models

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Columns      map[string]ColumnValue
	ColumnValues []interface{}
	ColumnNames  []string
	ColumnTypes  []string
}

type SingleRowResult struct {
	Result Row
}

func (s SingleRowResult) Row() Row {
	return s.Result
}

func (s SingleRowResult) Unmarshal(dest *interface{}) {
	marshaled, _ := json.Marshal(s.Row())

	json.Unmarshal(marshaled, dest)
}

type MultiRowResult struct {
	CurrentRow int
	Rows       []Row
	Columns    map[string][]ColumnValue
}

func (m MultiRowResult) NextRow() (bool, Row) {
	if m.CurrentRow > len(m.Rows) {
		var dud Row
		return false, dud
	}

	m.CurrentRow++
	return true, m.Rows[m.CurrentRow-1]
}

func (m MultiRowResult) Unmarshal(dest *interface{}) {
	var jsonMap map[string]Row

	for i, row := range m.Rows {
		jsonMap[fmt.Sprintf("result-%d", i)] = row
	}

	marshalled, _ := json.Marshal(jsonMap)

	json.Unmarshal(marshalled, dest)
}
