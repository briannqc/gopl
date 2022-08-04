package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquashAdjacentSpaces(t *testing.T) {
	s1 := []byte("abc bcd    cde def        efg")
	want1 := []byte("abc bcd cde def efg")
	got1 := SquashAdjacentSpaces(s1)
	assert.Equal(t, want1, got1)

	s2 := []byte("    abc bcd    cde def        efg")
	want2 := []byte(" abc bcd cde def efg")
	got2 := SquashAdjacentSpaces(s2)
	assert.Equal(t, want2, got2)

	s3 := []byte("abc bcd    cde def        efg      ")
	want3 := []byte("abc bcd cde def efg ")
	got3 := SquashAdjacentSpaces(s3)
	assert.Equal(t, want3, got3)

	s4 := []byte(" abc bcd    cde def        efg ")
	want4 := []byte(" abc bcd cde def efg ")
	got4 := SquashAdjacentSpaces(s4)
	assert.Equal(t, want4, got4)
}
