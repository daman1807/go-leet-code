package main

// https://leetcode.com/problems/binary-search/

func search(nums []int, target int) int {
	res := -1

	low, mid, high := 0, 0, len(nums)-1

	for low <= high {
		mid = (high + low) / 2

		if nums[mid] == target {
			res = mid
			break
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return res
}
