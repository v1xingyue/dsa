package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 4, 4, 5, 8, 9, 10, 22}
	pos := BinarySearch(nums, 23)
	t.Log(pos)
	v := SimpleHash("abcdefgmooAofjjhgsadkfas")
	t.Log("hash abcdefg : ", v)
}
