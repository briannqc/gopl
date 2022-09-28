package ex9_4

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRunGoroutinePipeline(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "size is 1",
			size: 1,
		},
		{
			name: "size is 10",
			size: 10,
		},
		{
			name: "size is 100",
			size: 100,
		},
		{
			name: "size is 1000",
			size: 1000,
		},
		{
			name: "size is 10_000",
			size: 10_000,
		},
		{
			name: "size is 100_000",
			size: 100_000,
		},
		{
			name: "size is 1_000_000",
			size: 1_000_000,
		},
		{
			name: "size is 10_000_000",
			size: 10_000_000,
		},
		{
			name: "size is 20_000_000",
			size: 20_000_000,
		},
		{
			name: "size is 30_000_000",
			size: 30_000_000,
		},
		{
			name: "size is 40_000_000",
			size: 40_000_000,
		},
		{
			name: "size is 50_000_000",
			size: 50_000_000,
		},
		{
			name: "size is 60_000_000",
			size: 60_000_000,
		},
		{
			name: "size is 70_000_000",
			size: 70_000_000,
		},
		{
			name: "size is 80_000_000",
			size: 80_000_000,
		},
		{
			name: "size is 90_000_000",
			size: 90_000_000,
		},
		{
			name: "size is 100_000_000",
			size: 100_000_000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in, out := make(chan int), make(chan int)
			go RunGoroutinePipeline(in, out, tt.size)

			start := time.Now()
			in <- 100
			got := <-out
			taken := time.Since(start)

			assert.Equal(t, 100, got)
			fmt.Printf("Time taken for size: %d is %v\n", tt.size, taken)
		})
	}
}
