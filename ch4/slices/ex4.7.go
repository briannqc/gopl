package main

import "unicode/utf8"

func reverse(b []byte) []byte {
	output := make([]byte, len(b))

	for n := len(b); n > 0; {
		r, size := utf8.DecodeRune(b)

		n -= size
		b = b[size:]

		runeInBytes := make([]byte, size)
		utf8.EncodeRune(runeInBytes, r)
		copy(output[n:], runeInBytes)
	}
	return output
}
