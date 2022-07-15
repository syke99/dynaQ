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

type ResultRows struct {
	CurrentRow int
	Results    []Row
}

func (m *ResultRows) NextRow() (bool, Row) {
	if m.CurrentRow > len(m.Results) {
		var dud Row
		return false, dud
	}

	m.CurrentRow++
	return true, m.Results[m.CurrentRow-1]
}

func (m *ResultRows) Unmarshal(dest *interface{}) error {
	var jsonMap map[string]Row

	for i, row := range m.Results {
		jsonMap[fmt.Sprintf("row-%d", i)] = row
	}

	marshalled, _ := json.Marshal(jsonMap)

	return json.Unmarshal(marshalled, dest)
}
