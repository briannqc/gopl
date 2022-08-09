package anonymousfunc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Graph map[int][]int

var graph = Graph{
	1:  {2, 10},
	2:  {1, 3, 9},
	3:  {2, 4, 7, 8},
	4:  {3, 5, 6},
	5:  {4},
	6:  {4},
	7:  {3},
	8:  {3, 9, 13},
	9:  {2, 8, 11},
	10: {1, 12},
	11: {9, 13},
	12: {10, 13, 14},
	13: {8, 11, 12},
	14: {12},
}

func breadthFirst(g Graph, root int) []int {
	seen := map[int]bool{}
	queue := new(Queue)
	order := make([]int, 0, len(g))
	queue.Push(root)
	seen[root] = true
	for !queue.IsEmpty() {
		node := queue.Pop()
		order = append(order, node)

		for _, n := range g[node] {
			if !seen[n] {
				seen[n] = true
				queue.Push(n)
			}
		}
	}

	return order
}

type Queue struct {
	data []int
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Push(v int) {
	q.data = append(q.data, v)
}

func (q *Queue) Pop() int {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func TestBreadthFirst(t *testing.T) {
	got := breadthFirst(graph, 1)
	want := []int{1, 2, 10, 3, 9, 12, 4, 7, 8, 11, 13, 14, 5, 6}
	assert.Equal(t, want, got)
}
