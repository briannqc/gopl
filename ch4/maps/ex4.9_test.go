package maps

import (
	"fmt"
	"io"
	"os"
	"sort"
	"testing"
)

func TestWordFred(t *testing.T) {
	f, err := os.Open("ex4.9_sample.txt")
	if err != nil {
		t.Fatalf("Failed to open sample text file")
	}
	defer func(c io.Closer) {
		_ = c.Close()
	}(f)

	freq := WordFred(f)

	pl := make(PairList, len(freq))
	i := 0
	for k, v := range freq {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))

	fmt.Println("Top 20 commonly used words in the file")
	for _, p := range pl[:20] {
		fmt.Printf("%v\t%v\n", p.Key, p.Value)
	}
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
