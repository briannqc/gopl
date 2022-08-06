package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildXKCDIndexFile(t *testing.T) {
	filename := "ex4.12_xkcd.json"
	err := BuildXKCDIndexFile(filename, 0, 10_000)
	assert.NoError(t, err)
}

func TestSearchComics(t *testing.T) {
	filename := "ex4.12_xkcd.json"
	comics, err := SearchComics(filename, "sexual medicine suggests that")
	assert.NoError(t, err)

	for _, c := range comics {
		fmt.Printf("URL: %v\n", c.URL)
		fmt.Println(c.Transcript)
		fmt.Println()
	}
}
