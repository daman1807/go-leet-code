package main

// https://leetcode.com/problems/binary-search/

func search(nums []int, target int) int {
	var left, mid, right int
	left, right = 0, len(nums)-1

	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
