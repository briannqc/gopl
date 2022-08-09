package anonymousfunc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func topoSortV3(m map[string][]string) ([]string, error) {
	var order []string
	var err error

	seenAndResolved := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			resolved, seen := seenAndResolved[item]
			if seen && !resolved {
				err = fmt.Errorf("%w cycle on: %s", err, item)
			}
			if !seen {
				seenAndResolved[item] = false
				visitAll(m[item])
				seenAndResolved[item] = true
				order = append(order, item)
			}
		}
	}
	for key := range m {
		visitAll([]string{key})
	}

	return order, err
}

func TestToposortV3(t *testing.T) {
	cyclicPrereqs := map[string][]string{
		"algorithms": {
			"data structures",
		},
		"calculus": {
			"linear algebra",
		},
		"linear algebra": {
			"calculus",
		},

		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},

		"data structures": {
			"discrete math",
		},
		"databases": {
			"data structures",
		},
		"discrete math": {
			"intro to programming",
		},
		"formal languages": {
			"discrete math",
		},
		"networks": {
			"operating systems",
		},
		"operating systems": {
			"data structures",
			"computer organization",
		},
		"programming languages": {
			"data structures",
			"computer organization",
		},
	}
	got, err := topoSortV3(cyclicPrereqs)
	assert.ErrorContains(t, err, "cycle")
	fmt.Println(got)
}
