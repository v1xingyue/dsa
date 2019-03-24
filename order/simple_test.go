package order

import (
	"fmt"
	"testing"
)

func BubbleSort(nums []int) []int {
	l := len(nums)
	t := 0
	if l == 0 {
		return nums
	}
	for i := 0; i < l; i++ {
		for j := 0; j < (l - i - 1); j++ {
			if nums[j] > nums[j+1] {
				t = nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = t
			}
		}
	}
	return nums
}

func InsertSort(nums []int) []int {
	l := len(nums)
	t := 0
	if 0 == l {
		return nums
	}
	for i := 1; i < l; i++ {
		t = nums[i]
		j := i - 1
		for ; j >= 0 && nums[j] > t; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = t
	}
	return nums
}

func TestBubble(t *testing.T) {
	test_nums := []int{10, 2, 9, 4, 5, 8}
	sort_nums := BubbleSort(test_nums)
	t.Log("BubbleSort :", sort_nums)
	test_nums2 := []int{10, 2, 9, 4, 5, 8}
	sort_nums2 := InsertSort(test_nums2)
	t.Log("Insert Sort :", sort_nums2)
}
