package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	semaphore := NewSemaphore(20)

	worklist := make(chan string, 100*len(os.Args[1:]))
	for _, web := range os.Args[1:] {
		worklist <- web
	}

	seen := make(map[string]bool)
	for web := range worklist {
		if seen[web] {
			continue
		}
		seen[web] = true

		semaphore.Acquire()
		go func(web string) {
			defer semaphore.Release()

			u, err := url.Parse(web)
			if err != nil {
				log.Println("url.Parse failed", err)
				return
			}

			links, err := downLoadAndExtractURLs(u.Host, web)
			if err != nil {
				panic(err)
				return
			}
			for _, l := range links {
				worklist <- l
			}
		}(web)
	}
}

type Semaphore struct {
	tokens chan bool
}

func NewSemaphore(weight int) *Semaphore {
	return &Semaphore{
		tokens: make(chan bool, weight),
	}
}

func (s *Semaphore) Acquire() {
	s.tokens <- true
}

func (s *Semaphore) Release() {
	<-s.tokens
}

func downLoadAndExtractURLs(origin, url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	buf := &bytes.Buffer{}
	tee := io.TeeReader(resp.Body, buf)

	file, err := os.Create(filenameForURL(resp.Request.URL))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()

	_, err = io.Copy(file, tee)
	if err != nil {
		return nil, err
	}

	node, err := html.Parse(buf)
	if err != nil {
		return nil, err
	}

	var links []string
	processNode := func(n *html.Node) {
		ref, ok := getLinkFromNode(n)
		if !ok {
			return
		}

		link, err := resp.Request.URL.Parse(ref)
		if err != nil {
			log.Println("Parse URL failed", err)
			return
		}
		if link.Host == origin {
			links = append(links, link.String())
		}
	}
	forEachNode(node, processNode)
	return links, nil
}

func forEachNode(n *html.Node, f func(node *html.Node)) {
	if f != nil {
		f(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, f)
	}
}

func filenameForURL(u *url.URL) string {
	if u.Path == "/" {
		return "index.html"
	}
	fpath := u.Path
	if strings.HasPrefix(fpath, "/") {
		fpath = fpath[1:]
	}
	if dir := path.Dir(fpath); dir != "." && dir != "/" {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	return fpath
}

func getLinkFromNode(n *html.Node) (string, bool) {
	if n.Type == html.ElementNode {
		return "", false
	}
	if n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				return a.Val, true
			}
		}
	} else if n.Data == "img" {
		for _, a := range n.Attr {
			if a.Key == "src" {
				return a.Val, true
			}
		}
	}
	return "", false
}
