package models

// ColumnValue contains the Type, Name, and Value of each column of each row in the results of a query
type ColumnValue struct {
	Type  string      `json:"type"`
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}
