package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllLinks(t *testing.T) {
	err := FindLinks("https://golang.org")
	assert.NoError(t, err)
}
