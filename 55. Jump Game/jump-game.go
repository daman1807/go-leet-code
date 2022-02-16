package main

// https://leetcode.com/problems/jump-game/

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func canJump(nums []int) bool {
	n := len(nums)

	if n == 0 {
		return false
	} else if n == 1 {
		return true
	}

	maxIndex := 0
	for i, v := range nums {
		maxIndex = MaxInt(maxIndex, i+v)
		if maxIndex >= n-1 {
			return true
		} else if maxIndex == i {
			return false
		}
	}
	return true
}
