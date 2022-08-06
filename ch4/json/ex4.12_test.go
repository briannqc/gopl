package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildXKCDIndexFile(t *testing.T) {
	filename := "ex4.12_xkcd.json"
	err := BuildXKCDIndexFile(filename, 0, 10_000)
	assert.NoError(t, err)
}
