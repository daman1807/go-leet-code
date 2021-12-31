package main

// https://leetcode.com/problems/pascals-triangle/

func generate(numRows int) [][]int {

	if numRows == 1 {
		return [][]int{{1}}
	}

	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	res := [][]int{{1}, {1, 1}}

	for i := 2; i < numRows; i += 1 {
		row := make([]int, i+1)

		row[0] = 1
		for j := 0; j < i-1; j += 1 {
			row[j+1] = res[i-1][j] + res[i-1][j+1]
		}
		row[i] = 1
		res = append(res, row)
	}

	return res
}
