package main

// https://leetcode.com/problems/find-pivot-index

func pivotIndex(nums []int) int {
	ans, sum, leftSum := -1, 0, 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if leftSum == sum-leftSum-nums[i] {
			return i
		}
		leftSum += nums[i]
	}

	return ans
}
