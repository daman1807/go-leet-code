package main

//https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	result := make([]int, 2)

	for idx, num := range nums {
		if val, ok := m[target-num]; ok {
			result[0] = idx
			result[1] = val
		} else {
			m[num] = idx
		}
	}

	return result
}
