package main

import "fmt"

func mainEx43() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	reverseArrPtr(&a)
	fmt.Println(a)
}

func reverseArrPtr(a *[10]int) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
}
