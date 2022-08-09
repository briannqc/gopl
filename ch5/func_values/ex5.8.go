package funcvalues

import (
	"io"

	"golang.org/x/net/html"
)

func FindElementByID(doc *html.Node, id string) *html.Node {
	var found *html.Node
	findInStart := func(w io.Writer, n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					found = n
					return false
				}
			}
		}
		return true
	}
	forEachNode(io.Discard, doc, findInStart, nil)
	return found
}
