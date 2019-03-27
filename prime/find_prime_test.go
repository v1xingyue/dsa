package prime

import (
	"testing"
)

func IsPrime(num int) bool {
	switch num {
	case 0, 1:
		return false
	case 2:
		return true
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func TestIsPrime(t *testing.T) {
	var nums []int
	for i := 0; i < 1000000; i++ {
		if IsPrime(i) {
			nums = append(nums, i)
		}
	}
	t.Log(nums)
}
