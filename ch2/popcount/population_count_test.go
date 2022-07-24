package popcount_test

import (
	"math/rand"
	"testing"

	"github.com/briannqc/gopl/ch2/popcount"
)

func TestPopCount(t *testing.T) {
	for n := uint64(0); n < 1000; n++ {
		v1 := popcount.PopCount(n)
		v2 := popcount.PopCountV2(n)
		v3 := popcount.PopCountV3(n)
		if v1 != v2 || v2 != v3 {
			t.Fatalf("Got different result in different solutions, input: %d, results: %d %d %d", n, v1, v2, v3)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	x := rand.Uint64()
	for n := 0; n < b.N; n++ {
		popcount.PopCount(x)
	}
}

func BenchmarkPopCountV2(b *testing.B) {
	x := rand.Uint64()
	for n := 0; n < b.N; n++ {
		popcount.PopCountV2(x)
	}
}

func BenchmarkPopCountV3(b *testing.B) {
	x := rand.Uint64()
	for n := 0; n < b.N; n++ {
		popcount.PopCountV3(x)
	}
}
