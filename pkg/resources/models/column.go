package models

type ColumnValue struct {
	Type  string      `json:"type"`
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}
