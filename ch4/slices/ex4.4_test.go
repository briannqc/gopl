package main

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Ogirinal:%v\n", s)

	n := 17
	s = Rotate(s, n)
	fmt.Printf("After rotate:%v, n=%v\n", s, n)
}
