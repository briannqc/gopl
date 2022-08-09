package variadicfunc

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	found := make([]*html.Node, 0)

	var forEachNode func(*html.Node)
	forEachNode = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if contains(names, n.Data) {
				found = append(found, n)
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			forEachNode(c)
		}
	}

	forEachNode(doc)
	return found
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TestElementsByTagName(t *testing.T) {
	resp, err := http.Get("https://google.com")
	if !assert.NoError(t, err) {
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if !assert.NoError(t, err) {
		return
	}

	got := ElementsByTagName(doc, "div")
	assert.NotEmpty(t, got)
}
