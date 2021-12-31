package main

// https://leetcode.com/problems/contains-duplicate-ii/

func isInBuffer(buffer []int, val int) bool {
	for _, v := range buffer {
		if v == val {
			return true
		}
	}
	return false
}

func nonZero(val int) int {
	if val < 0 {
		return 0
	}
	return val
}

func containsNearbyDuplicate(nums []int, k int) bool {
	for idx, num := range nums {
		buffer := nums[nonZero(idx-k):nonZero(idx)]
		if isInBuffer(buffer, num) {
			return true
		}

	}
	return false
}
