package search

import (
	"fmt"
)

func BinarySearch(nums []int, find int) int {
	low, high := 0, len(nums)
	for {
		mid := (low + high) / 2
		if nums[mid] < find {
			low = mid
		} else if nums[mid] > find {
			high = mid
		} else {
			return mid
		}
		if high == low || mid == low {
			return -1
		}
	}
}

func SimpleHash(msg string) string {
	hash_num := 0
	for _, c := range msg {
		hash_num += int(c)
	}
	return fmt.Sprintf("%x", hash_num)
}
