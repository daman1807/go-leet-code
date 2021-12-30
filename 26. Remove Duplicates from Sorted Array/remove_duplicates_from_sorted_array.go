package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-array

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	var slow, fast int

	for fast = 1; fast < len(nums); fast += 1 {
		if nums[slow] != nums[fast] {
			slow += 1
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}
