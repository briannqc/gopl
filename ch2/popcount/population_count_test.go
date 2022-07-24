package popcount_test

import (
	"math/rand"
	"testing"

	"github.com/briannqc/gopl/ch2/popcount"
)

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
