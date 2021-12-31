package main

//https://leetcode.com/problems/search-insert-position

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums) - 1

	if target > nums[high]{
		return high + 1
	}
	if target < nums[low]{
		return low
	}


	for low < high {
		mid := low + (high - low) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if nums[low] < target {
		return low + 1
	}
	return low
}