package main

import (
	"crypto/sha256"
	"fmt"
)

func mainEx41() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	diff := uint(0)
	for i := 0; i < len(c1); i++ {
		b1, b2 := c1[i], c2[i]
		diff += uint(pc[b1&b2])
	}

	fmt.Println("Bits diff:", diff)
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
