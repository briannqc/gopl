package main

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)
	return local, n, err
}

func TestFetch(t *testing.T) {
	url := "https://go.dev/"
	filename, n, err := fetch(url)
	assert.NoError(t, err)
	assert.Greater(t, n, int64(0))

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		t.Fail()
	}
}
