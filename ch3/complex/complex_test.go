package main

import (
	"io"
	"testing"
)

/**
BenchmarkRenderComplex128
BenchmarkRenderComplex128-12                   5         211868036 ns/op        13049676 B/op 1997031 allocs/op
PASS
ok      github.com/briannqc/gopl/ch3/complex    2.586s
*/
func BenchmarkRenderComplex128(b *testing.B) {
	for n := 0; n < b.N; n++ {
		renderComplex128(io.Discard)
	}
}

/**
BenchmarkRenderComplex64
BenchmarkRenderComplex64-12            4         311535405 ns/op        13049552 B/op    1997013 allocs/op
PASS
ok      github.com/briannqc/gopl/ch3/complex    2.820s
*/
func BenchmarkRenderComplex64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		renderComplex64(io.Discard)
	}
}

/**
BenchmarkRenderBigFloat
BenchmarkRenderBigFloat-12             1        15757807446 ns/op       11459572120 B/op        245704952 allocs/op
PASS
ok      github.com/briannqc/gopl/ch3/complex    16.078s
*/
func BenchmarkRenderBigFloat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		renderBigFloat(io.Discard)
	}
}

/**
*** Test killed with quit: ran too long (11m0s).
exit status 2
FAIL    github.com/briannqc/gopl/ch3/complex    660.014s
*/
func BenchmarkRenderBigRat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		renderBigRat(io.Discard)
	}
}
