package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

func BuildXKCDIndexFile(filename string, fromID int, toID int) error {
	if fromID > toID {
		return fmt.Errorf("fromID (%v) must be less than or equal to toID (%v)", fromID, toID)
	}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()
	buf := bufio.NewWriter(file)

	ch := make(chan *Comic, 100)
	go func() {
		count := 0
		for comic := range ch {
			b, err := json.Marshal(comic)
			if err != nil {
				log.Printf("Marshalling comic failed, num: %d, err: %v", comic.Num, err)
				continue
			}
			if _, err := buf.Write(b); err != nil {
				log.Printf("Writing comic to file failed, num: %d, err: %v", comic.Num, err)
				continue
			}
			buf.WriteString("\n")

			count++
			if count%50 == 0 {
				buf.Flush()
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Add(toID - fromID + 1)
	for id := fromID; id <= toID; id++ {
		go func(ch chan *Comic, id int) {
			defer wg.Done()

			comic, err := fetchComicByID(id)
			if err != nil {
				return
			}
			ch <- comic
		}(ch, id)
	}

	wg.Wait()
	return nil
}

func fetchComicByID(id int) (*Comic, error) {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Getting comic %d failed, err: %v", id, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got %v code while getting comic: %v", resp.Status, id)
	}

	var comic Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		log.Printf("Decoding comic %d failed, err: %v", id, err)
		return nil, err
	}

	comic.URL = url
	return &comic, nil
}

type Comic struct {
	URL        string
	Num        int
	Title      string
	Transcript string
}
