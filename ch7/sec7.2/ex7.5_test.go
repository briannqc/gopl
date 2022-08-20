package sec7_2

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimitReader(t *testing.T) {
	s := `
Exercise 7.5: The LimitReader function in the io package accepts
an io.Reader r and a number of bytes n, and returns another Reader
that reads from r but reports an end-of-file condition after n
bytes. Implement it.`

	r := LimitReader(strings.NewReader(s), 10)
	n, err := r.Read(make([]byte, 100))
	assert.Equal(t, 10, n)
	assert.ErrorIs(t, err, io.EOF)

	r = LimitReader(strings.NewReader(s), int64(len(s)+10))
	n, err = r.Read(make([]byte, len(s)+10))
	assert.Equal(t, len(s), n)
	assert.ErrorIs(t, err, nil)
}
