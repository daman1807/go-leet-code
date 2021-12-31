package main

import "fmt"

// https://leetcode.com/problems/summary-ranges/

func summaryRanges(nums []int) []string {

	var res []string

	if len(nums) == 0 {
		return res
	}

	start, end := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			end = nums[i]
		} else {
			if start == end {
				res = append(res, fmt.Sprintf("%d", start))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", start, end))
			}
			start, end = nums[i], nums[i]
		}
	}

	if start == end {
		res = append(res, fmt.Sprintf("%d", start))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", start, end))
	}

	return res
}
