package main

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			if res[0] == -1 {
				res[0] = i
				res[1] = i
			} else {
				res[1] = i
			}
		}
	}
	return res
}
