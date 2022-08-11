package panicrecover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func add(first int, others ...int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			result = r.(int)
		}
	}()

	sum := first
	for _, n := range others {
		sum += n
	}
	panic(sum)
}

func TestAdd(t *testing.T) {
	result := add(1, 2, 3, 4, 5, 6)
	assert.Equal(t, 21, result)
}
