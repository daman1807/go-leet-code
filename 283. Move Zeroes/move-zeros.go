package main

// https://leetcode.com/problems/move-zeroes/

func moveZeroes(nums []int) {

	if len(nums) > 1 {
		slow, fast := 0, 0

		for slow < len(nums) && fast < len(nums) {
			if nums[slow] == 0 {
				if nums[fast] != 0 {
					nums[slow], nums[fast] = nums[fast], nums[slow]
					slow++
				}
				fast++
			} else {
				slow++
				fast++
			}
		}
	}
}
