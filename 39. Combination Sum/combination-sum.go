package main

import "sort"

// https://leetcode.com/problems/combination-sum/

func addToTrail(num int, trail []int) []int {
	newTrail := make([]int, len(trail)+1)
	i := 0
	for i = 0; i < len(trail); i++ {
		newTrail[i] = trail[i]
	}
	newTrail[i] = num
	return newTrail
}

func search(index, curr, target int, trail, candidates []int, res *[][]int) {
	if curr > target {
		return
	} else if curr < target {
		for i := index; i < len(candidates); i++ {
			search(i, curr+candidates[i], target, addToTrail(candidates[i], trail), candidates, res)
		}
	} else {
		*res = append(*res, trail)
	}
}

func combinationSum(candidates []int, target int) [][]int {

	res := make([][]int, 0)
	sort.Ints(candidates)

	for i := 0; i < len(candidates); i++ {
		trail := []int{candidates[i]}
		search(i, candidates[i], target, trail, candidates, &res)
	}

	return res

}
