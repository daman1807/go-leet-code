package main

// https://leetcode.com/problems/permutations/

func addToTrail(num int, trail []int) []int {
	newTrail := make([]int, len(trail)+1)
	for i := 0; i < len(trail); i++ {
		newTrail[i] = trail[i]
	}
	newTrail[len(trail)] = num
	return newTrail
}

func notInTrail(num int, trail []int) bool {
	for i := 0; i < len(trail); i++ {
		if num == trail[i] {
			return false
		}
	}
	return true
}

func getPermutations(nums, trail []int, res *[][]int) {
	if len(trail) == len(nums) {
		*res = append(*res, trail)
	} else {
		for i := 0; i < len(nums); i++ {
			if notInTrail(nums[i], trail) {
				getPermutations(nums, addToTrail(nums[i], trail), res)
			}
		}
	}
}

func permute(nums []int) [][]int {
	var res [][]int
	getPermutations(nums, []int{}, &res)
	return res
}
