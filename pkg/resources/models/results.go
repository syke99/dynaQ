package models

import (
	"encoding/json"
	"fmt"
)

// Result is used during unmarshaling
type Result struct {
	Columns      map[string]ColumnValue
	ColumnValues []interface{}
	ColumnNames  []string
	ColumnTypes  []string
}

// ResultRows represents all the rows of the result set returned by a query
type ResultRows struct {
	CurrentRow int
	Results    []Row
}

// NextRow returns the next Row in the result set, preceded by a boolean to denote whether there are any subsequent rows following the one returned by this method
func (m ResultRows) NextRow() (bool, Row) {
	if m.CurrentRow > len(m.Results) {
		var dud Row
		return false, dud
	}

	m.CurrentRow++
	return true, m.Results[m.CurrentRow-1]
}

// Unmarshal takes a pointer to an empty interface as an argument for unmarshaling results into a stringified json object; if dest isn't an empty interface, this method will fail and return an error
func (m ResultRows) Unmarshal(dest *interface{}) error {
	var jsonMap map[string]Row

	for i, row := range m.Results {
		jsonMap[fmt.Sprintf("result-%d", i)] = row
	}

	marshalled, _ := json.Marshal(jsonMap)

	return json.Unmarshal(marshalled, dest)
}
