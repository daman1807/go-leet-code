package main

// https://leetcode.com/problems/pascals-triangle-ii/

func getRow(rowIndex int) []int {

	if rowIndex == 0 {
		return []int{1}
	}

	res := make([][]int, rowIndex+1)
	res[0] = []int{1}
	res[1] = []int{1, 1}

	for i := 2; i < rowIndex+1; i++ {
		row := make([]int, rowIndex+1)
		row[0] = 1
		for j := 0; j < i-1; j++ {
			row[j+1] = res[i-1][j] + res[i-1][j+1]
		}
		row[i] = 1
		res[i] = row
	}

	return res[rowIndex]
}
