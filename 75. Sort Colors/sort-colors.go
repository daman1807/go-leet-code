package main

// https://leetcode.com/problems/sort-colors/

func sortColors(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i; j > 0 && nums[j-1] > nums[j]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}
