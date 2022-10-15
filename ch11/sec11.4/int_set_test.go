package sec11_4

import (
	"testing"

	"github.com/briannqc/gopl/ch6"
)

/*
goos: darwin
goarch: amd64
pkg: github.com/briannqc/gopl/ch11/sec11.4
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkAdd
BenchmarkAdd-12    	113990029	        10.42 ns/op
*/
func BenchmarkAdd(b *testing.B) {
	var s1 ch6.IntSet
	for i := 0; i < b.N; i++ {
		s1.Add(i)
	}
}

/*
goos: darwin
goarch: amd64
pkg: github.com/briannqc/gopl/ch11/sec11.4
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkHas
BenchmarkHas-12    	147576277	         7.975 ns/op
PASS
*/
func BenchmarkHas(b *testing.B) {
	var s ch6.IntSet
	for i := 0; i < 30000; i++ {
		s.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Has(i)
	}
}

/*
goos: darwin
goarch: amd64
pkg: github.com/briannqc/gopl/ch11/sec11.4
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkUnionWith
BenchmarkUnionWith-12    	 1516719	       805.2 ns/op
PASS
*/
func BenchmarkUnionWith(b *testing.B) {
	var s1, s2 ch6.IntSet
	for i := 0; i < 30000; i++ {
		s1.Add(i)
		s2.Add(i * 2)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s1.UnionWith(&s2)
	}
}
