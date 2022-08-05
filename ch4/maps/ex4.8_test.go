package maps

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharCount(t *testing.T) {
	input := strings.NewReader("> Test run finished at 04/08/2022, 23:51:42 <")
	want := map[string]int{
		Letter: 17,
		Digit:  14,
		Space:  7,
		Symbol: 2,
		Other:  5,
	}

	got := CharCount(input)
	assert.Equal(t, want, got)
}
