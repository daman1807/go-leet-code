package main

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	res := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		if val, ok := m[nums[i]]; ok {
			res[0] = val
			res[1] = i
			break
		} else {
			m[target-nums[i]] = i
		}
	}
	return res
}
