package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	exec1()
	exec1Inefficient()
	exec2()
}

// exec1 modifies the echo program to also print os.Args[0],
// the name of the cmd that invoked it.
func exec1() {
	start := time.Now()

	fmt.Println("=== Exercise 1 ===")
	fmt.Println(strings.Join(os.Args, " "))

	fmt.Println("exec1() took", time.Now().Sub(start))
}

func exec1Inefficient() {
	start := time.Now()

	fmt.Println("=== Exercise 1 (Inefficient) ===")
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Println("exec1Inefficient() took", time.Now().Sub(start))
}

// exec2 modifies the echo program to print the index and value
// of each of its arguments, one per line.
func exec2() {
	start := time.Now()

	fmt.Println("=== Exercise 2 ===")
	for index, arg := range os.Args {
		fmt.Println(index, arg)
	}

	fmt.Println("exec2() took", time.Now().Sub(start))
}
