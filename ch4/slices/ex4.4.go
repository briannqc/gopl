package main

func Rotate(s []int, n int) []int {
	n = n % len(s)
	s = append(s, s[:n]...)
	s = s[n:]
	return s
}
