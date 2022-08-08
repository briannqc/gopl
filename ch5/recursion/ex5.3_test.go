package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintTextNodes(t *testing.T) {
	doc, err := Fetch("https://golang.org")
	assert.NoError(t, err)

	PrintTextNodes(doc)
}
