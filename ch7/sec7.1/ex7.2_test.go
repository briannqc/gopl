package sec7_1

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingWriter(t *testing.T) {
	s := "Exercise 7.2: Write a function CountingWriter with the signature below that, given an io.Writer, returns a new Writer that wraps the original, and a pointer to an int64 variable that at any moment contains the number of bytes written to the new Writer."
	buf := &bytes.Buffer{}
	writer, written := CountingWriter(buf)
	_, err := writer.Write([]byte(s))

	assert.NoError(t, err)
	assert.Equal(t, s, buf.String())
	assert.Equal(t, int64(len(s)), *written)
}
