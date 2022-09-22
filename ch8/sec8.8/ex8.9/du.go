package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	verbose := flag.Bool("v", false, "Show verbose progress messages")

	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(100 * time.Millisecond)
	}
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go func(dir string) {
			defer wg.Done()
			computeAndDisplayDiskUsage(dir, tick)
		}(root)
	}

	wg.Wait()
}

func computeAndDisplayDiskUsage(root string, tick <-chan time.Time) {
	var wg sync.WaitGroup
	fileSizes := make(chan int64)

	wg.Add(1)
	go walkDir(root, &wg, fileSizes)
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(root, nfiles, nbytes)
		}
	}
	printDiskUsage(root, nfiles, nbytes)
}

func printDiskUsage(root string, nfiles int64, nbytes int64) {
	fmt.Printf("%s: %d files %.1f GB\n", root, nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	for _, entry := range dirEntries(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var semaphore = make(chan struct{}, 20)

func dirEntries(dir string) []os.FileInfo {
	semaphore <- struct{}{}
	defer func() {
		<-semaphore
	}()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
