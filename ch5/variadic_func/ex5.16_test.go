package variadicfunc

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func JoinStrings(sep string, s ...string) string {
	return strings.Join(s, sep)
}

func TestJoinStrings(t *testing.T) {
	assert.Equal(t, "a b c", JoinStrings(" ", "a", "b", "c"))
	assert.Equal(t, "a", JoinStrings(" ", "a"))
	assert.Equal(t, "", JoinStrings(" ", ""))
}
