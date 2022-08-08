package recursion

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func PrintTextNodes(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.TextNode {
		fmt.Println(strings.TrimSpace(n.Data))
	}
	PrintTextNodes(n.FirstChild)
	PrintTextNodes(n.NextSibling)
}
