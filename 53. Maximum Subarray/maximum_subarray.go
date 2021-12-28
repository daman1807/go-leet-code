package main

// https://leetcode.com/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	res := -99999
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > res {
			res = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	return res
}
