package main

import "math"

// https://leetcode.com/problems/third-maximum-number/

func thirdMax(nums []int) int {

	first := math.MinInt64
	second := math.MinInt64
	third := math.MinInt64

	for i := 0; i < len(nums); i += 1 {

		if nums[i] > first {
			second, third = first, second
			first = nums[i]
		} else if nums[i] > second && nums[i] < first {
			third = second
			second = nums[i]
		} else if nums[i] > third && nums[i] < second {
			third = nums[i]
		}

	}

	if third == math.MinInt64 {
		return first
	}

	return third
}
