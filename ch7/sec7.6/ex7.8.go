package sec7_6

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
)

// Table is sortable upto 3 recently clicked columns.
//
// Exercise 7.8: Many GUIs provide a table widget with
// a stateful multi-tier sort: the primary sort key is
// the most recently clicked column head, the secondary
// sort key is the second-most recently clicked column
// head, and so on. Define an implementation of
// sort.Interface for use by such a table. Compare that
// approach with repeated sorting using sort.Stable.
type Table struct {
	recentlyClickedColumns []string
	columns                []string
	data                   map[string][]string
	length                 int
}

func NewTable(columns []string, data map[string][]string) *Table {
	copiedColumns := make([]string, len(columns))
	copy(copiedColumns, columns)

	length := 0
	copiedData := make(map[string][]string)
	for col, values := range data {
		if length == 0 {
			length = len(values)
		} else if length != len(values) {
			panic(fmt.Errorf("all columns in the table must have the same length, found: %d and %d", length, len(values)))
		}
		copiedValues := make([]string, len(values))
		copy(copiedValues, values)
		copiedData[col] = copiedValues
	}

	return &Table{
		recentlyClickedColumns: nil,
		columns:                copiedColumns,
		data:                   copiedData,
		length:                 length,
	}
}

func (t *Table) String() string {
	buf := &bytes.Buffer{}

	format := strings.Repeat("%v\t", len(t.columns))
	format = format + "\n"

	tw := new(tabwriter.Writer).Init(buf, 0, 8, 2, ' ', 0)

	var columns []interface{}
	for _, c := range t.columns {
		columns = append(columns, c)
	}
	_, _ = fmt.Fprintf(tw, format, columns...)

	var dashes []interface{}
	for range t.columns {
		dashes = append(dashes, "-----")
	}
	_, _ = fmt.Fprintf(tw, format, dashes...)

	for i := 0; i < t.length; i++ {
		var lineValues []interface{}
		for _, col := range t.columns {
			lineValues = append(lineValues, t.data[col][i])
		}
		_, _ = fmt.Fprintf(tw, format, lineValues...)
	}

	_ = tw.Flush()
	return buf.String()
}

func (t *Table) Click(column string) {
	t.recentlyClickedColumns = append([]string{column}, t.recentlyClickedColumns...)
	if len(t.recentlyClickedColumns) > 3 {
		t.recentlyClickedColumns = t.recentlyClickedColumns[:3]
	}
}

func (t *Table) Len() int {
	return t.length
}

func (t *Table) Less(i, j int) bool {
	for _, col := range t.recentlyClickedColumns {
		if t.data[col][i] < t.data[col][j] {
			return true
		}
	}
	return false
}

func (t *Table) Swap(i, j int) {
	for col := range t.data {
		t.data[col][i], t.data[col][j] = t.data[col][j], t.data[col][i]
	}
}
