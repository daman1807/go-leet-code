package main

// https://leetcode.com/problems/spiral-matrix/

func spiralOrder(matrix [][]int) []int {
	var i, k int
	var left, right, top, bottom int

	n, m := len(matrix), len(matrix[0])

	res := make([]int, n*n)

	top, bottom = 0, n-1
	left, right = 0, m-1

	for k < n*m {
		// left->right
		for i = left; i <= right; i++ {
			res[k] = matrix[top][i]
			k++
		}
		top++

		// top->bottom
		for i = top; i <= bottom; i++ {
			res[k] = matrix[i][right]
			k++
		}
		right--

		// right->left
		for i = right; i >= left; i-- {
			res[k] = matrix[bottom][i]
			k++
		}
		bottom--

		// bottom->top
		for i = bottom; i >= top; i-- {
			res[k] = matrix[i][left]
			k++
		}
		left++
	}

	return res
}
