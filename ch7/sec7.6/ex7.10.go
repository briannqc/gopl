package main

import "sort"

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		equal := !s.Less(i, j) && !s.Less(j, i)
		if !equal {
			return false
		}
	}
	return true
}
