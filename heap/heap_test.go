package heap

import (
	"testing"
)

func TestHeapify(t *testing.T) {
	tree := []int{9, 10, 8, 7, 6, 5, 4}
	heapify(tree, 0)
	t.Log(tree)
}

func TestBuildHeap(t *testing.T) {
	tree := []int{1, 2, 3, 5, 8, 10, 20, 6, 9}
	build_heap(tree)
	t.Log(tree)
}
