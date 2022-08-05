package maps

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

const (
	Letter = "Letter"
	Digit  = "Digit"
	Space  = "Space"
	Symbol = "Symbol"
	Other  = "Other"
)

func CharCount(reader io.Reader) map[string]int {
	in := bufio.NewReader(reader)
	counts := map[string]int{}
	checkFuncs := []struct {
		kind string
		fn   func(rune) bool
	}{
		{
			kind: Letter,
			fn:   unicode.IsLetter,
		},
		{
			kind: Digit,
			fn:   unicode.IsDigit,
		},
		{
			kind: Space,
			fn:   unicode.IsSpace,
		},
		{
			kind: Symbol,
			fn:   unicode.IsSymbol,
		},
		{
			kind: Other,
			fn: func(rune) bool {
				return true
			},
		},
	}

	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		for _, ck := range checkFuncs {
			if ck.fn(r) {
				counts[ck.kind]++
				break
			}
		}
	}
	return counts
}
