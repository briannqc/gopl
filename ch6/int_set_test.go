package ch6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSet_Len(t *testing.T) {
	tests := []struct {
		set  *IntSet
		want int
	}{
		{
			set: func() *IntSet {
				var set IntSet
				set.Add(1)
				set.Add(2)
				set.Add(3)
				return &set
			}(),
			want: 3,
		},
		{
			set: func() *IntSet {
				var set IntSet
				set.Add(100)
				set.Add(300)
				return &set
			}(),
			want: 2,
		},
		{
			set: func() *IntSet {
				var set IntSet
				return &set
			}(),
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := tt.set.Len()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIntSet_Remove(t *testing.T) {
	var set IntSet

	for i := 0; i < 1000; i++ {
		set.Add(i)
		assert.True(t, set.Has(i))
		set.Remove(i)
		assert.False(t, set.Has(i))
	}
}

func TestIntSet_Clear(t *testing.T) {
	var set IntSet

	for i := 0; i < 1000; i++ {
		set.Add(i)
		assert.Equal(t, 1, set.Len())
		set.Clear()
		assert.Equal(t, 0, set.Len())
	}
}

func TestIntSet_Copy(t *testing.T) {
	var set IntSet
	for i := 0; i < 1000; i++ {
		set.Add(i)
	}
	copied := set.Copy()
	set.Clear()

	assert.Equal(t, 1000, copied.Len())
	for i := 0; i < 1000; i++ {
		assert.True(t, copied.Has(i))
	}
}
