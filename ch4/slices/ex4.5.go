package main

func EleminateAdjacentDupplicates(s []string) []string {
	i, j := 0, 0
	for ; j < len(s); j++ {
		if s[i] != s[j] {
			i++
			s[i] = s[j]
		}
	}
	return s[:i+1]
}
