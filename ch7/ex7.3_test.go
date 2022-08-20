package ch7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tree_String(t *testing.T) {
	tree := Sort([]int{10, 2, 4, 1, 45, 3, 5, 2, 0})
	assert.Equal(t, "0 1 2 2 3 4 5 10 45", tree.String())
}
