package sec11_2

import (
	"math/rand"
	"testing"
	"time"
)

func TestIsPalindrome(t *testing.T) {
	type TestCase struct {
		name string
		args string
		want bool
	}

	seed := time.Now().UTC().UnixNano()
	t.Logf("seed: %v", seed)
	rng := rand.New(rand.NewSource(seed))

	var tests []TestCase
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		tests = append(tests, TestCase{name: p, args: p, want: true})
	}
	for i := 0; i < 1000; i++ {
		np := randomNonPalindrome(rng)
		tests = append(tests, TestCase{name: np, args: np, want: false})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func randomNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) + 5
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		c := 'a' + rng.Intn('z'-'a')
		r := rune(c)
		runes[i] = r
		runes[n-1-i] = r + 1
	}
	return string(runes)
}
