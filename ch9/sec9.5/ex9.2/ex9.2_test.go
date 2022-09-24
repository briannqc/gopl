package ex9_2

import "testing"

func TestPopCount(t *testing.T) {
	tests := []struct {
		name string
		x    uint64
		want int
	}{
		{
			name: "PopCount(0) = 0",
			x:    0,
			want: 0,
		},
		{
			name: "PopCount(1) = 1",
			x:    1,
			want: 1,
		},
		{
			name: "PopCount(2) = 1",
			x:    2,
			want: 1,
		},
		{
			name: "PopCount(3) = 2",
			x:    3,
			want: 2,
		},
	}
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopCount(tt.x); got != tt.want {
				t.Errorf("PopCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
