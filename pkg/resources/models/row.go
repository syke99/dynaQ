package models

type Row struct {
	CurrentColumn int
	Columns       []ColumnValue
}

func (r Row) NextColumn() (bool, ColumnValue) {
	if r.CurrentColumn > len(r.Columns) {
		var dud ColumnValue
		return false, dud
	}

	r.CurrentColumn++
	return true, r.Columns[r.CurrentColumn-1]
}
