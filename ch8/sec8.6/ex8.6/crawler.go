package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	log.Println("Discovering", url)
	tokens <- struct{}{} // acquire a token
	list, err := extract(url)
	if err != nil {
		log.Println("Extract failed, err:", err)
	}
	<-tokens // release the token

	return list
}

func extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	node, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var links []string
	extractLinksFromNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}

				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				if link.Scheme == "http" || link.Scheme == "https" {
					links = append(links, link.String())
				}
			}
		}
	}
	forEachNode(node, extractLinksFromNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

type DepthLink struct {
	urls  []string
	depth int
}

func main() {
	depthLimit := flag.Int("depth", 5, "depth limiting")
	flag.Parse()

	worklist := make(chan DepthLink)
	var noPending int

	noPending++
	go func() {
		worklist <- DepthLink{
			urls:  flag.Args(),
			depth: 0,
		}
	}()

	var count int
	seen := make(map[string]bool)
	for ; noPending > 0; noPending-- {
		depthLinks := <-worklist
		for _, link := range depthLinks.urls {
			if !seen[link] {
				seen[link] = true
				noPending++
				go func(link string, depth int) {
					links := crawl(link)

					msg := &strings.Builder{}
					_, _ = fmt.Fprintf(msg, "Links discovered from: %s\n", link)
					for _, l := range links {
						_, _ = fmt.Fprintf(msg, "\t%s\n", l)
					}
					log.Print(msg.String())

					count += len(links)
					if depth+1 <= *depthLimit {
						worklist <- DepthLink{
							urls:  links,
							depth: depth + 1,
						}
					}
				}(link, depthLinks.depth)
			}
		}
	}

	log.Printf("Discovered: %d links with depth: %d", count, *depthLimit)
}
