package funcvalues

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

func PrintPrettyHTML(w io.Writer, n *html.Node) {
	forEachNode(w, n, startElement, endElement)
}

func forEachNode(w io.Writer, n *html.Node, pre, post func(io.Writer, *html.Node) bool) {
	if pre != nil {
		if cont := pre(w, n); !cont {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(w, c, pre, post)
	}

	if post != nil {
		if cont := post(w, n); !cont {
			return
		}
	}
}

var depth int

func startElement(w io.Writer, n *html.Node) bool {
	if n.Type == html.ElementNode {
		fmt.Fprintf(w, "%*s<%s ", depth*2, "", n.Data)
		depth++

		var attrs []string
		for _, a := range n.Attr {
			attrs = append(attrs, fmt.Sprintf("%s='%s'", a.Key, a.Val))
		}
		fmt.Fprint(w, strings.Join(attrs, " "))

		if n.FirstChild != nil {
			fmt.Fprint(w, ">\n")
		}
	} else if n.Type == html.TextNode {
		fmt.Fprint(w, n.Data)
	}
	return true
}

func endElement(w io.Writer, n *html.Node) bool {
	if n.Type == html.ElementNode {
		depth--

		if n.FirstChild != nil {
			fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
		} else {
			fmt.Fprint(w, "/>\n")
		}
	}
	return true
}
