package models

type Result struct {
	Columns      map[string]QueryValue
	ColumnValues []interface{}
	ColumnNames  []string
	ColumnTypes  []string
}

type SingleRowResult struct {
	Result map[string]QueryValue
}

func (s SingleRowResult) Row() map[string]QueryValue {
	return s.Result
}

type MultiRowResult struct {
	CurrentRow int
	Results    []map[string]QueryValue
}

func (m MultiRowResult) NextRow() (bool, map[string]QueryValue) {
	if m.CurrentRow > len(m.Results) {
		dud := make(map[string]QueryValue)
		return false, dud
	}

	m.CurrentRow++
	return true, m.Results[m.CurrentRow-1]
}
