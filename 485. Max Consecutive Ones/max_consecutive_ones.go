package main

// https://leetcode.com/problems/max-consecutive-ones/

func findMaxConsecutiveOnes(nums []int) int {
	res := 0

	start, end := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if end-start > res {
				res = end - start
			}
			start, end = i, i
		} else {
			end++
		}
	}

	if end-start > res {
		res = end - start
	}

	return res
}
