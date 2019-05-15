package heap

import (
	"log"
)

//from i down must be heapify
func heapify(tree []int, i int) {
	max := i
	l := len(tree)
	if i >= l {
		return
	}
	left := i*2 + 1
	right := i*2 + 2
	if left < l && tree[i] < tree[left] {
		max = left
	}
	if right < l && tree[i] < tree[right] {
		max = right
	}
	if max != i {
		tree[i], tree[max] = tree[max], tree[i]
		heapify(tree, max)
	} else {
		return
	}
}

func build_heap(tree []int) {
	l := len(tree)
	n := (l - 1) / 2
	for ; n >= 0; n-- {
		log.Println(n)
		heapify(tree, n)
		log.Println(tree)
	}
}
