package models

// Row represents an entire row of the result set of a query
type Row struct {
	CurrentColumn int
	Columns       []ColumnValue
}

// NextColumn returns the next ColumnValue in the row, preceded by a boolean to denote whether there are any subsequent columns following the one returned by this method
func (r Row) NextColumn() (bool, ColumnValue) {
	if r.CurrentColumn > len(r.Columns) {
		var dud ColumnValue
		return false, dud
	}

	r.CurrentColumn++
	return true, r.Columns[r.CurrentColumn-1]
}
