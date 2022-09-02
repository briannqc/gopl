package main

import (
	"io"
	"testing"
)

/*
goos: darwin
goarch: amd64
pkg: github.com/briannqc/gopl/ch8/sec8.4/ex8.5
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_render
Benchmark_render-12               	       5	 213851703 ns/op
*/
func Benchmark_render(b *testing.B) {
	for i := 0; i < b.N; i++ {
		render(io.Discard)
	}
}

/*
Benchmark_renderInParallel1
Benchmark_renderInParallel1-12    	       3	 496690642 ns/op
*/
func Benchmark_renderInParallel1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		renderInParallel1(io.Discard)
	}
}

/*
Benchmark_renderInParallel2
Benchmark_renderInParallel2-12    	       3	 388499770 ns/op
*/
func Benchmark_renderInParallel2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		renderInParallel2(io.Discard)
	}
}
