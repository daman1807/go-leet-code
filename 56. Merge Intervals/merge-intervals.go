package main

import "sort"

// https://leetcode.com/problems/merge-intervals/

func min(a, b int) int {
	if b > a {
		return a
	}
	return b
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func merge(intervals [][]int) [][]int {
	var res [][]int

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start, end := intervals[0][0], intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			res = append(res, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		} else {
			start = min(start, intervals[i][0])
			end = max(end, intervals[i][1])
		}
	}

	res = append(res, []int{start, end})
	return res
}
