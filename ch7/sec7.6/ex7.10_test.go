package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	type args struct {
		s sort.Interface
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "GIVEN empty THEN return true",
			args: args{
				s: sort.IntSlice{},
			},
			want: true,
		},
		{
			name: "GIVEN single element THEN return true",
			args: args{
				s: sort.IntSlice{1},
			},
			want: true,
		},
		{
			name: "GIVEN odd palindrome THEN return true",
			args: args{
				s: sort.IntSlice{1, 2, 3, 2, 1},
			},
			want: true,
		},
		{
			name: "GIVEN even palindrome THEN return true",
			args: args{
				s: sort.IntSlice{1, 2, 3, 3, 2, 1},
			},
			want: true,
		},
		{
			name: "GIVEN not palindrome THEN return false",
			args: args{
				s: sort.IntSlice{1, 2, 3, 2, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsPalindrome(tt.args.s), "IsPalindrome(%v)", tt.args.s)
		})
	}
}
