package internal

type Result struct {
	Columns      map[string]interface{}
	ColumnValues []interface{}
	ColumnNames  []string
}
