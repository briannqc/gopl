package recursion

import (
	"fmt"

	"golang.org/x/net/html"
)

func FindAllLinks(rootURL string) error {
	doc, err := Fetch(rootURL)
	if err != nil {
		return err
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
	return nil
}

func visitAllLinks(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode {

		if n.Data == "a" || n.Data == "link" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		} else if n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		} else if n.Data == "script" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
	}

	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}
