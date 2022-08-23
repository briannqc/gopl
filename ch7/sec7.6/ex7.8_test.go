package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable_ClickAndSort(t *testing.T) {
	table := NewTable(
		[]string{"Title", "Artist", "Album", "Year", "Length"},
		map[string][]string{
			"Title":  {"Go", "Ready 2 Go", "Go", "Go Ahead"},
			"Artist": {"Moby", "Martin Solveig", "Delilah", "Alicia Keys"},
			"Album":  {"Moby", "Smash", "From the Roots Up", "As I Am"},
			"Year":   {"1992", "2011", "2012", "2007"},
			"Length": {"3m37s", "4m24s", "3m38s", "4m36s"},
		})
	sort.Sort(table)
	wantBeforeClicking :=
		`Title       Artist          Album              Year   Length  
-----       -----           -----              -----  -----   
Go          Moby            Moby               1992   3m37s   
Ready 2 Go  Martin Solveig  Smash              2011   4m24s   
Go          Delilah         From the Roots Up  2012   3m38s   
Go Ahead    Alicia Keys     As I Am            2007   4m36s   
`
	assert.Equal(t, wantBeforeClicking, table.String())

	table.Click("Artist")
	sort.Sort(table)
	wantAfterClickingArtist :=
		`Title       Artist          Album              Year   Length  
-----       -----           -----              -----  -----   
Go Ahead    Alicia Keys     As I Am            2007   4m36s   
Go          Delilah         From the Roots Up  2012   3m38s   
Ready 2 Go  Martin Solveig  Smash              2011   4m24s   
Go          Moby            Moby               1992   3m37s   
`
	assert.Equal(t, wantAfterClickingArtist, table.String())

	table.Click("Title")
	sort.Sort(table)
	wantAfterClickingTitle :=
		`Title       Artist          Album              Year   Length  
-----       -----           -----              -----  -----   
Go          Delilah         From the Roots Up  2012   3m38s   
Go          Moby            Moby               1992   3m37s   
Go Ahead    Alicia Keys     As I Am            2007   4m36s   
Ready 2 Go  Martin Solveig  Smash              2011   4m24s   
`
	assert.Equal(t, wantAfterClickingTitle, table.String())

	table.Click("Year")
	sort.Sort(table)
	wantAfterClickingYear :=
		`Title       Artist          Album              Year   Length  
-----       -----           -----              -----  -----   
Go Ahead    Alicia Keys     As I Am            2007   4m36s   
Ready 2 Go  Martin Solveig  Smash              2011   4m24s   
Go          Moby            Moby               1992   3m37s   
Go          Delilah         From the Roots Up  2012   3m38s   
`
	assert.Equal(t, wantAfterClickingYear, table.String())
}
