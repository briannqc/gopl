package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/briannqc/gopl/ch9/sec9.7/memo"
)

func main() {
	m := memo.New(httpGetBody)
	var wg sync.WaitGroup
	for _, url := range os.Args[1:] {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			start := time.Now()
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			value, err := m.Get(ctx, url)
			if err != nil {
				fmt.Printf("%s. %s, err: %v\n", url, time.Since(start), err)
			} else {
				fmt.Printf("%s. %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			}
		}(url)
	}

	wg.Wait()
}

func httpGetBody(ctx context.Context, url string) (interface{}, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return ioutil.ReadAll(resp.Body)
}
