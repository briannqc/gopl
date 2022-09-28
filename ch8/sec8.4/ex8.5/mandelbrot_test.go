package main

import (
	"io"
	"runtime"
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

/*
goos: darwin
goarch: amd64
pkg: github.com/briannqc/gopl/ch8/sec8.4/ex8.5
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_renderInParallel2_GOMAXPROCS_1
Benchmark_renderInParallel2_GOMAXPROCS_1-12    	       1	1085448367 ns/op
*/
func Benchmark_renderInParallel2_GOMAXPROCS_1(b *testing.B) {
	runtime.GOMAXPROCS(1)
	for i := 0; i < b.N; i++ {
		renderInParallel2(io.Discard)
	}
}

/*
goos: darwin
goarch: amd64
pkg: github.com/briannqc/gopl/ch8/sec8.4/ex8.5
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_renderInParallel2_GOMAXPROCS_2
Benchmark_renderInParallel2_GOMAXPROCS_2-12    	       2	 524221212 ns/op
*/
func Benchmark_renderInParallel2_GOMAXPROCS_2(b *testing.B) {
	runtime.GOMAXPROCS(2)
	for i := 0; i < b.N; i++ {
		renderInParallel2(io.Discard)
	}
}

/*
goos: darwin
goarch: amd64
pkg: github.com/briannqc/gopl/ch8/sec8.4/ex8.5
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_renderInParallel2_GOMAXPROCS_4
Benchmark_renderInParallel2_GOMAXPROCS_4-12    	       3	 389211096 ns/op
*/
func Benchmark_renderInParallel2_GOMAXPROCS_4(b *testing.B) {
	runtime.GOMAXPROCS(4)
	for i := 0; i < b.N; i++ {
		renderInParallel2(io.Discard)
	}
}

/*
goos: darwin
goarch: amd64
pkg: github.com/briannqc/gopl/ch8/sec8.4/ex8.5
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_renderInParallel2_GOMAXPROCS_8
Benchmark_renderInParallel2_GOMAXPROCS_8-12    	       3	 381482230 ns/op
*/
func Benchmark_renderInParallel2_GOMAXPROCS_8(b *testing.B) {
	runtime.GOMAXPROCS(8)
	for i := 0; i < b.N; i++ {
		renderInParallel2(io.Discard)
	}
}

/*
goos: darwin
goarch: amd64
pkg: github.com/briannqc/gopl/ch8/sec8.4/ex8.5
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_renderInParallel2_GOMAXPROCS_16
Benchmark_renderInParallel2_GOMAXPROCS_16-12    	       3	 394391888 ns/op
*/
func Benchmark_renderInParallel2_GOMAXPROCS_16(b *testing.B) {
	runtime.GOMAXPROCS(16)
	for i := 0; i < b.N; i++ {
		renderInParallel2(io.Discard)
	}
}
