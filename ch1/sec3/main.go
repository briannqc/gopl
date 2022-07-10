package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	textInFiles := make(map[string][]string)
	fileNames := os.Args[1:]
	if len(fileNames) == 0 {
		fmt.Fprintln(os.Stderr, "At least one file is required")
		return
	}

	for _, fileName := range fileNames {
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Open file: %v failed, err: %v", fileName, err)
			continue
		}
		countLines(f, textInFiles)
		f.Close()
	}

	for line, files := range textInFiles {
		if len(files) > 1 {
			fmt.Printf("Text: %v in %v\n", line, files)
		}
	}
}

func countLines(f *os.File, textInFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		textInFiles[input.Text()] = append(textInFiles[input.Text()], f.Name())
	}
}
