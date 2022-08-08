package recursion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountElements(t *testing.T) {
	doc, err := Fetch("https://golang.org")
	assert.NoError(t, err)

	m := map[string]int{}
	CountElements(m, doc)
	fmt.Println(m)
}
