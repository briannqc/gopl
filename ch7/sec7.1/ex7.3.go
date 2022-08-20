package sec7_1

import (
	"bytes"
	"fmt"
)

type Tree struct {
	value       int
	left, right *Tree
}

// String reveals the sequence of values in the tree.
func (t *Tree) String() string {
	if t == nil {
		return ""
	}
	buf := &bytes.Buffer{}
	if t.left != nil {
		_, _ = fmt.Fprintf(buf, "%s ", t.left.String())
	}
	_, _ = fmt.Fprintf(buf, "%d", t.value)
	if t.right != nil {
		_, _ = fmt.Fprintf(buf, " %s", t.right.String())
	}
	return buf.String()
}

// Sort sorts values in place.
func Sort(values []int) *Tree {
	var root *Tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	return root
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *Tree, value int) *Tree {
	if t == nil {
		// Equivalent to return &Tree{value: value}.
		t = new(Tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
