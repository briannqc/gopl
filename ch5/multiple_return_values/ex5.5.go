package multiplereturnvalues

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (nwords, nimages int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	nwords, nimages = countWordsAndImages(doc)
	return nwords, nimages, nil
}

func countWordsAndImages(n *html.Node) (nwords, nimages int) {
	if n == nil {
		return 0, 0
	}
	if n.Type == html.TextNode {
		nwords += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		nimages++
	}

	firstChildWords, firstChildImages := countWordsAndImages(n.FirstChild)
	nextSiblingWords, nextSiblingImages := countWordsAndImages(n.NextSibling)

	nwords = nwords + firstChildWords + nextSiblingWords
	nimages = nimages + firstChildImages + nextSiblingImages
	return nwords, nimages
}
