package main

import "fmt"

const k = 1024

const (
	KB = k
	MB = k * KB
	GB = k * MB
	TB = k * GB
	PB = k * TB
	EB = k * PB
	ZB = k * EB
	YB = k * ZB
)

func main() {
	fmt.Printf("%T %[1]v\n", KB)
	fmt.Printf("%T %[1]v\n", MB)
	fmt.Printf("%T %[1]v\n", GB)
	fmt.Printf("%T %[1]v\n", TB)
	fmt.Printf("%T %[1]v\n", PB)
	fmt.Printf("%T %[1]v\n", EB)
	// Overflow
	//fmt.Printf("%T %[1]v\n", ZB)

	// But this still works, impressive
	fmt.Println(YB / ZB)
}
