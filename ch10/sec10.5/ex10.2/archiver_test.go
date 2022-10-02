package archive_test

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	archive "github.com/briannqc/gopl/ch10/sec10.5/ex10.2"
	_ "github.com/briannqc/gopl/ch10/sec10.5/ex10.2/tar"
	_ "github.com/briannqc/gopl/ch10/sec10.5/ex10.2/zip"
)

func TestOpen(t *testing.T) {
	for _, file := range []string{"rah.zip", "rah.tar"} {
		t.Run(file, func(t *testing.T) {
			b := &bytes.Buffer{}
			f, err := os.Open(filepath.Join("testdata", file))
			if err != nil {
				t.Error(file, err)
			}
			r, err := archive.Open(f)
			if err != nil {
				t.Error(file, err)
			}
			_, err = io.Copy(b, r)
			if err != nil {
				t.Error(file, err)
			}
			want := `rah/b:
contentsB
rah/a:
contentsA
`
			got := b.String()
			if got != want {
				t.Errorf("%s: got %q, want %q", file, got, want)
			}
		})
	}
}
