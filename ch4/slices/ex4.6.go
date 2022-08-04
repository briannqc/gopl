package main

import "unicode"

func SquashAdjacentSpaces(s []byte) []byte {
	inSpaces := false
	n := 0
	for i := 0; i < len(s); i++ {
		if unicode.IsSpace(rune(s[i])) {
			if !inSpaces {
				s[n] = s[i]
				n++
				inSpaces = true
			}
			continue
		}

		if n != i {
			s[n] = s[i]
		}
		n++
		inSpaces = false
	}

	return s[:n]
}
