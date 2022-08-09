package variadicfunc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Max[N Number](first N, others ...N) N {
	max := first
	for _, n := range others {
		if n > max {
			max = n
		}
	}
	return max
}

func TestMax(t *testing.T) {
	got := Max(1, 2, 3, 4, 5, 6, 7, 8, 1, 20)
	assert.Equal(t, 20, got)
}

func Min[N Number](first N, others ...N) N {
	min := first
	for _, n := range others {
		if n < min {
			min = n
		}
	}
	return min
}

func TestMin(t *testing.T) {
	got := Min(1, 2, 3, 4, 5, 6, 7, 8, 1, 20)
	assert.Equal(t, 1, got)
}
