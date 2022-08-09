package funcvalues

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestPrintPrettyHTML(t *testing.T) {
	n, err := fetch("https://google.com")
	if !assert.NoError(t, err) {
		return
	}

	PrintPrettyHTML(os.Stdout, n)
}

func fetch(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return html.Parse(resp.Body)
}
