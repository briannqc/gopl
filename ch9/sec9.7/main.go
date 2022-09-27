package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/briannqc/gopl/ch9/sec9.7/memo"
)

func main() {
	m := memo.NewV2(httpGetBody)
	var wg sync.WaitGroup
	for _, url := range os.Args[1:] {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s. %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
		}(url)
	}

	wg.Wait()
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return ioutil.ReadAll(resp.Body)
}
