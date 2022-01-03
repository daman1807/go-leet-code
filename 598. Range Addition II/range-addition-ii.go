package main

// https://leetcode.com/problems/range-addition-ii/

func maxCount(m int, n int, ops [][]int) int {

	x, y := m, n

	for i := 0; i < len(ops); i++ {
		if ops[i][0] < x {
			x = ops[i][0]
		}
		if ops[i][1] < y {
			y = ops[i][1]
		}
	}
	return x * y
}
