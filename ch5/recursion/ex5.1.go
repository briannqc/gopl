package recursion

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func Fetch(rootURL string) (*html.Node, error) {
	resp, err := http.Get(rootURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got: %v for url: %v", resp.Status, rootURL)
	}

	return html.Parse(resp.Body)
}

func FindLinks(rootURL string) error {
	doc, err := Fetch(rootURL)
	if err != nil {
		return err
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
	return nil
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}