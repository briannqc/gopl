package main

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"
)

var (
	//go:embed table.html
	templateTable []byte
)

func (t *Table) Columns() []string {
	columns := make([]string, len(t.columns))
	copy(columns, t.columns)
	return columns
}

func (t *Table) Rows() [][]string {
	rows := make([][]string, 0, t.length)
	for i := 0; i < t.length; i++ {
		row := make([]string, 0, len(t.columns))
		for _, col := range t.columns {
			row = append(row, t.data[col][i])
		}
		rows = append(rows, row)
	}
	return rows
}

func main() {
	table := NewTable(
		[]string{"Title", "Artist", "Album", "Year", "Length"},
		map[string][]string{
			"Title":  {"Go", "Ready 2 Go", "Go", "Go Ahead"},
			"Artist": {"Moby", "Martin Solveig", "Delilah", "Alicia Keys"},
			"Album":  {"Moby", "Smash", "From the Roots Up", "As I Am"},
			"Year":   {"1992", "2011", "2012", "2007"},
			"Length": {"3m37s", "4m24s", "3m38s", "4m36s"},
		})
	tablePage, err := template.New("milestoneList").Parse(string(templateTable))
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/table", func(w http.ResponseWriter, req *http.Request) {
		_ = tablePage.Execute(w, table)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
