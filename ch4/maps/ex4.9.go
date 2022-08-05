package maps

import (
	"bufio"
	"io"
	"strings"
)

func WordFred(rd io.Reader) map[string]int {
	in := bufio.NewScanner(rd)
	in.Split(bufio.ScanWords)

	freq := map[string]int{}
	for in.Scan() {
		word := in.Text()
		freq[strings.ToLower(word)]++
	}

	return freq
}
