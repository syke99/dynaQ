package models

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
