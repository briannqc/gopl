package sec7_5

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaveat(t *testing.T) {
	var buf *bytes.Buffer
	assert.False(t, isNil(buf), "buf is (type=*bytes.Buffer, value=nil), hence not nil")

	var file *os.File
	assert.False(t, isNil(file), "buf is (type=*os.File, value=nil), hence not nil")

	var w io.Writer
	assert.True(t, isNil(w), "w is (type=nil, value=nil), hence nil")
}

// isNil return true if w is nil (type=nil, value=nil)
func isNil(w io.Writer) bool {
	return w == nil
}
