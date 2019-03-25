package order

import (
	"testing"
)

func BubbleSort(nums []int) []int {
	l := len(nums)
	t := 0
	moved := false
	if l == 0 {
		return nums
	}

	//先做第一层循环，用来标记已经排好序的位置,
	//最后边的i个元素是依次从大到小排列的
	for i := 0; i < l; i++ {
		moved = false
		//然后，从0到排好序的位置，两两对比，如果后边的元素小于前边的元素，那么就交换两个数字的位置
		for j := 0; j < (l - i - 1); j++ {
			if nums[j] > nums[j+1] {
				t = nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = t
				moved = true
			}
		}
		// 如果循环内，没有交换任何的两个元素，表示序列已经有序，可以跳出循环
		if moved == false {
			break
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
	test_nums := []int{10, 2, 9, 9, 9, 8, 12, 0, 90, 23}
	t.Log("Input list : ", test_nums)
	BubbleSort(test_nums)
	t.Log("BubbleSort :", test_nums)
	test_nums2 := []int{10, 2, 9, 4, 5, 8}
	t.Log("Input list : ", test_nums2)
	sort_nums2 := InsertSort(test_nums2)
	t.Log("Insert Sort :", sort_nums2)
}
