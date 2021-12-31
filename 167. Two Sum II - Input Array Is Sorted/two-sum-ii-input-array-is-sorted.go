package main

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

func twoSum(numbers []int, target int) []int {
	res := make([]int, 2)
	low, high := 0, len(numbers)-1

	for high > low {
		if numbers[high]+numbers[low] > target {
			high -= 1
		} else if numbers[high]+numbers[low] < target {
			low += 1
		} else {
			res[0] = low + 1
			res[1] = high + 1
			break
		}
	}
	return res

}
