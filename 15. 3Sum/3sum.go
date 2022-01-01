package main

// https://leetcode.com/problems/3sum

import "sort"

func threeSum(nums []int) [][]int {

	var res [][]int

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		low, high := i+1, len(nums)-1

		for low < high {
			if nums[i]+nums[low]+nums[high] > 0 {
				high--
			} else if nums[i]+nums[low]+nums[high] < 0 {
				low++
			} else {
				res = append(res, []int{nums[i], nums[low], nums[high]})

				for low < high && nums[low] == nums[low+1] {
					low++
				}

				for low < high && nums[high] == nums[high-1] {
					high--
				}

				low++
				high--
			}
		}
	}
	return res
}
