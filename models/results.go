package models

type Result struct {
	Columns      map[string]QueryValue
	ColumnValues []interface{}
	ColumnNames  []string
	ColumnTypes  []string
}
