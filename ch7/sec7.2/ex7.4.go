package sec7_2

import (
	"bytes"
	"io"
)

type stringReader struct {
	data    string
	pointer int
}

func (s *stringReader) Read(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	n, err := buf.WriteString(s.data)
	s.pointer += n
	if err != nil {
		return n, err
	}
	if s.pointer >= len(s.data) {
		return n, io.EOF
	}
	return n, nil
}

func NewStringReader(s string) io.Reader {
	return &stringReader{data: s, pointer: 0}
}
