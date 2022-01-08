package main

// https://leetcode.com/problems/missing-number/

func missingNumber(nums []int) int {

	res := len(nums) * (len(nums) + 1) / 2
	for i := 0; i < len(nums); i++ {
		res -= nums[i]
	}
	return res
}
