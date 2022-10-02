package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"

	archive "github.com/briannqc/gopl/ch10/sec10.5/ex10.2"
)

func init() {
	archive.RegisterFormat("zip", "PK", 0, NewReader)
}

type reader struct {
	zipReader *zip.Reader
	filesLeft []*zip.File
	rc        io.ReadCloser
	toWrite   string
}

func (r *reader) Read(b []byte) (int, error) {
	if r.rc == nil && len(r.filesLeft) == 0 {
		return 0, io.EOF
	}

	if r.rc == nil {
		f := r.filesLeft[0]
		r.filesLeft = r.filesLeft[1:]
		var err error
		r.rc, err = f.Open()
		if err != nil {
			return 0, fmt.Errorf("read zip: %s", err)
		}
		if f.Mode()&os.ModeDir == 0 {
			r.toWrite = f.Name + ":\n"
		}
	}
	written := 0
	if len(r.toWrite) > 0 {
		n := copy(b, r.toWrite)
		b = b[n:]
		r.toWrite = r.toWrite[n:]
		written += n
	}
	n, err := r.rc.Read(b)
	written += n
	if err != nil {
		_ = r.rc.Close()
		r.rc = nil
		if err == io.EOF {
			return written, nil
		}
	}
	return written, nil
}

func NewReader(f *os.File) (io.Reader, error) {
	stat, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("new zip reader: %w", err)
	}
	r, err := zip.NewReader(f, stat.Size())
	if err != nil {
		return nil, fmt.Errorf("new zip reader: %w", err)
	}
	return &reader{
		zipReader: r,
		filesLeft: r.File,
		rc:        nil,
		toWrite:   "",
	}, nil
}
