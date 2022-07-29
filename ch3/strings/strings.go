package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(commaSignedFloat(arg))
	}
}

// comma inserts commas in a non-negative decimal integer string.
//
// Exercise 3.10. Write a non-recursive version of comma, using bytes.Buffer
// instead of string concatenation
func comma(s string) string {
	formatted := bytes.NewBuffer(make([]byte, 0, len(s)))

	firstSectionLastIdx := len(s) % 3
	formatted.Write([]byte(s[0:firstSectionLastIdx]))

	for i := firstSectionLastIdx; i < len(s); i += 3 {
		from := i
		to := i + 3
		if from > 0 {
			formatted.WriteByte(',')
		}

		formatted.Write([]byte(s[from:to]))
	}
	return formatted.String()
}

// commaSignedFloat inserts commas in a signed float string.
//
// Exercise 3.11. Enhance comma so that it deals correctly with floating-point
// numbers and an optional sign.
func commaSignedFloat(s string) string {
	formatted := bytes.NewBuffer(make([]byte, 0, len(s)))

	if strings.HasPrefix(s, "-") {
		formatted.WriteByte('-')
		s = s[1:]
	}

	if dotIdx := strings.Index(s, "."); dotIdx >= 0 {
		formatted.WriteString(comma(s[0:dotIdx]))
		s = s[dotIdx:]
		formatted.WriteString(s)
	} else {
		formatted.WriteString(comma(s))
	}

	return formatted.String()
}
