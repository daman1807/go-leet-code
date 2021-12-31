package main

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

func findDisappearedNumbers(nums []int) []int {

	exists := make([]bool, len(nums))
	res := make([]int, 0)

	for i := 0; i < len(nums); i += 1 {
		exists[nums[i]-1] = true
	}

	for i := 0; i < len(exists); i++ {
		if !exists[i] {
			res = append(res, i+1)
		}
	}

	return res

}
