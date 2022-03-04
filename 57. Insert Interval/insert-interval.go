package main

// https://leetcode.com/problems/insert-interval/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if b > a {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int

	if len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}

	curIndex := 0
	for curIndex < len(intervals) && intervals[curIndex][1] < newInterval[0] {
		res = append(res, intervals[curIndex])
		curIndex++
	}

	for curIndex < len(intervals) && intervals[curIndex][0] <= newInterval[1] {
		newInterval = []int{min(newInterval[0], intervals[curIndex][0]), max(newInterval[1], intervals[curIndex][1])}
		curIndex++
	}
	res = append(res, newInterval)

	for curIndex < len(intervals) {
		res = append(res, intervals[curIndex])
		curIndex++
	}

	return res

}
