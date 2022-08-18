package ch6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSet_IntersectWith(t *testing.T) {
	var s1, s2 IntSet
	for i := 0; i < 600; i++ {
		s1.Add(i)
	}
	for i := 500; i < 1000; i++ {
		s2.Add(i)
	}

	intersect := s1.IntersectWith(&s2)
	assert.Equal(t, 100, intersect.Len())
	for i := 500; i < 600; i++ {
		assert.True(t, intersect.Has(i))
	}
}

func TestIntSet_DifferentWith(t *testing.T) {
	var s1, s2 IntSet
	for i := 0; i < 600; i++ {
		s1.Add(i)
	}
	for i := 500; i < 1000; i++ {
		s2.Add(i)
	}

	diff := s1.DifferentWith(&s2)
	assert.Equal(t, 500, diff.Len())
	for i := 0; i < 500; i++ {
		assert.True(t, diff.Has(i))
	}
}

func TestIntSet_SymmetricDifference(t *testing.T) {
	var s1, s2 IntSet
	for i := 0; i < 600; i++ {
		s1.Add(i)
	}
	for i := 500; i < 1000; i++ {
		s2.Add(i)
	}

	symDiff := s1.SymmetricDifference(&s2)
	assert.Equal(t, 900, symDiff.Len())
	for i := 0; i < 500; i++ {
		assert.Truef(t, symDiff.Has(i), "Should have: %d but not", i)
	}
	for i := 600; i < 1000; i++ {
		assert.Truef(t, symDiff.Has(i), "Should have: %d but not", i)
	}
}

func TestIntSet_Elems(t *testing.T) {
	tests := []struct {
		set  *IntSet
		want []int
	}{
		{
			set: func() *IntSet {
				var set IntSet
				set.Add(1)
				set.Add(2)
				set.Add(3)
				return &set
			}(),
			want: []int{1, 2, 3},
		},
		{
			set: func() *IntSet {
				var set IntSet
				set.Add(100)
				set.Add(300)
				return &set
			}(),
			want: []int{100, 300},
		},
		{
			set: func() *IntSet {
				var set IntSet
				return &set
			}(),
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.set.String(), func(t *testing.T) {
			got := tt.set.Elems()
			assert.Equal(t, tt.want, got)
		})
	}
}

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

func TestIntSet_AddAll(t *testing.T) {
	var set IntSet
	set.AddAll(1, 2, 3)

	assert.Equal(t, 3, set.Len())
	assert.True(t, set.Has(1))
	assert.True(t, set.Has(2))
	assert.True(t, set.Has(3))
}
