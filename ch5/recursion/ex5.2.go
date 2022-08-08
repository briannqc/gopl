package recursion

import "golang.org/x/net/html"

func CountElements(m map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	CountElements(m, n.FirstChild)
	CountElements(m, n.NextSibling)
}
