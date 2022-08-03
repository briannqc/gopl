package main

import (
	"fmt"
	"testing"
)

func TestEleminateAdjacentDupplicates(t *testing.T) {
	s := []string{"a", "a", "b", "c", "c", "d", "d", "e"}
	fmt.Printf("Original: %v\n", s)

	s = EleminateAdjacentDupplicates(s)
	fmt.Printf("After eliminating adjacent dupplicates: %v\n", s)
}
