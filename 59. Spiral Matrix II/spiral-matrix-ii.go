package main

// https://leetcode.com/problems/spiral-matrix-ii/

func generateMatrix(n int) [][]int {
	var res [][]int
	var i, num, end int
	var left, right, top, bottom int

	for i = 0; i < n; i++ {
		row := make([]int, n)
		res = append(res, row)
	}

	end = n * n
	left, top = 0, 0
	right, bottom = n-1, n-1

	for num < end {
		// right->left
		for i = left; i <= right; i++ {
			num++
			res[top][i] = num
		}
		top++

		// top->bottom
		for i = top; i <= bottom; i++ {
			num++
			res[i][right] = num
		}
		right--

		// left->right
		for i = right; i >= left; i-- {
			num++
			res[bottom][i] = num
		}
		bottom--

		// bottom->top
		for i = bottom; i >= top; i-- {
			num++
			res[i][left] = num
		}
		left++

	}

	return res
}
