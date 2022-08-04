package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	b := []byte("Hello, 世界")
	want := []byte("界世 ,olleH")
	got := reverse(b)
	assert.Equal(t, want, got)
}
