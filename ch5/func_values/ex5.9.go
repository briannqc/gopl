package funcvalues

import "strings"

func expand(s string, f func(string) string) string {
	replacement := f("foo")
	return strings.ReplaceAll(s, "$foo", replacement)
}
