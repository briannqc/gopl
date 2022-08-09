package anonymousfunc

import (
	"fmt"
	"testing"
)

func topoSortV2(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	for key := range m {
		visitAll([]string{key})
	}

	return order
}

func TestToposortV2(t *testing.T) {
	got := topoSortV2(prereqs)
	fmt.Println(got)
}
