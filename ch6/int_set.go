package ch6

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith creates and returns an IntSet of elements present in both sets.
func (s *IntSet) IntersectWith(t *IntSet) *IntSet {
	lenS := len(s.words)
	lenT := len(t.words)

	lenIntersect := lenS
	if lenIntersect > lenT {
		lenIntersect = lenT
	}

	var intersect IntSet
	intersect.words = make([]uint64, 0, lenIntersect)
	for i := 0; i < lenIntersect; i++ {
		word := s.words[i] & t.words[i]
		intersect.words = append(intersect.words, word)
	}
	return &intersect
}

// DifferentWith creates and returns an IntSet of elements present in s but not t.
func (s *IntSet) DifferentWith(t *IntSet) *IntSet {
	diff := s.Copy()
	lenT := len(t.words)
	lenDiff := len(diff.words)

	minLen := lenDiff
	if minLen > lenT {
		minLen = lenT
	}

	for i := 0; i < minLen; i++ {
		diff.words[i] &= ^t.words[i]
	}
	return diff
}

// SymmetricDifference creates and returns an IntSet of
// elements present in one set or the other but not both.
func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	lenS := len(s.words)
	lenT := len(t.words)

	minLen := lenS
	if minLen > lenT {
		minLen = lenT
	}

	symDiffLen := lenS
	if symDiffLen < lenT {
		symDiffLen = lenT
	}

	var symDiff IntSet
	symDiff.words = make([]uint64, 0, symDiffLen)

	for i := 0; i < minLen; i++ {
		symDiff.words = append(symDiff.words, s.words[i]^t.words[i])
	}
	if lenS < lenT {
		symDiff.words = append(symDiff.words, t.words[lenS:]...)
	} else {
		symDiff.words = append(symDiff.words, s.words[lenT:]...)
	}
	return &symDiff
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the number of elements
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for word > 0 {
			count++
			word = word & (word - 1)
		}
	}
	return count
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		return
	}

	s.words[word] &= ^(1 << bit)
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	var copied IntSet
	copied.words = make([]uint64, len(s.words))
	copy(copied.words, s.words)
	return &copied
}

func (s *IntSet) AddAll(nums ...int) {
	for _, n := range nums {
		s.Add(n)
	}
}
