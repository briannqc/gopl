package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	firstResp := fetchFirst(os.Args[1:])
	fmt.Println(firstResp)
}

func fetchFirst(urls []string) string {
	responses := make(chan string, len(urls))

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(ctx context.Context, url string) {
			defer wg.Done()
			resp, err := fetch(ctx, url)
			if err != nil {
				log.Println("Fetch failed, url:", url, "err:", err)
				return
			}

			responses <- resp
		}(ctx, url)
	}

	go func() {
		wg.Wait()
		close(responses)
	}()

	select {
	case resp := <-responses:
		cancel()
		return resp
	}
}

func fetch(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	buf := &bytes.Buffer{}
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return "", err
	}
	return buf.String(), nil
}
