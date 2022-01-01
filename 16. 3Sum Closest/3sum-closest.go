package main

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/3sum-closest/

func Abs(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}

func threeSumClosest(nums []int, target int) int {
	res, sum, minDelta := 0, 0, math.MaxInt64
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums)-1 {
			low, high := i+1, len(nums)-1
			for low < high {
				sum = nums[i] + nums[low] + nums[high]
				if sum > target {
					if Abs(target-sum) < minDelta {
						minDelta = Abs(target - sum)
						res = sum
					}
					high--
				} else if sum < target {
					if Abs(target-sum) < minDelta {
						minDelta = Abs(target - sum)
						res = sum
					}
					low++
				} else {
					return target
				}
			}
		}
	}
	return res
}
