package anonymousfunc

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func crawlAndSave(rootURL string) error {
	parsedURL, err := url.Parse(rootURL)
	if err != nil {
		return err
	}

	doc, err := fetch(rootURL)
	if err != nil {
		return err
	}

	host := parsedURL.Host
	dir := strings.ReplaceAll(host, ":", "_")
	if err := os.MkdirAll(dir, 0777); err != nil {
		return err
	}

	saveNodeOfSameDomain := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				absURL, err := parsedURL.Parse(a.Val)
				if err != nil {
					continue
				}
				if absURL.Host != host {
					continue
				}
				if err := savePageToFile(dir, absURL.String()); err != nil {
					log.Printf("Saving %s to file failed, err: %v", absURL, err)
				}
			}
		}
	}
	forEachNode(doc, saveNodeOfSameDomain, nil)
	return nil
}

func savePageToFile(dir string, pageURL string) error {
	resp, err := http.Get(pageURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fileName := strings.NewReplacer(":", "_", "/", "_").Replace(pageURL) + ".html"
	f, err := os.Create(path.Join(dir, fileName))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}

func TestCrawlAndSave(t *testing.T) {
	err := crawlAndSave("https://golang.org")
	assert.NoError(t, err)
}
