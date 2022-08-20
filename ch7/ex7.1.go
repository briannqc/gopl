package ch7

import (
	"bufio"
	"bytes"
)

type WordCounter int

// Write counts the words in p
func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
	return 0, scanner.Err()
}

type LineCounter int

// Write counts the lines in p
func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c++
	}
	return 0, scanner.Err()
}
