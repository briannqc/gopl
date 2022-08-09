package anonymousfunc

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"golang.org/x/net/html"
)

func PrintPrettyHTML(url string) {
	var dept int
	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			dept++
			fmt.Printf("%*s<%s", dept*2, "", n.Data)
			if n.FirstChild != nil && n.FirstChild.Type != html.TextNode {
				fmt.Println(">")
			}
		}
	}
	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.FirstChild != nil && n.FirstChild.Type != html.TextNode {
				fmt.Printf("%*s</%s>\n", dept*2, "", n.Data)
			} else {
				fmt.Println("/>")
			}
			dept--
		}
	}

	doc, err := fetch(url)
	if err != nil {
		log.Fatalln(err)
	}
	forEachNode(doc, startElement, endElement)
}

func fetch(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return html.Parse(resp.Body)
}

func TestPrintPrettyHTML(t *testing.T) {
	PrintPrettyHTML("https://google.com")
}
