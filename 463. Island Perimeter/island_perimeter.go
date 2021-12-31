package main

// https://leetcode.com/problems/island-perimeter/

func cellPerimeter(grid [][]int, i, j int) int {
	peri := 0

	if i == 0 || grid[i-1][j] == 0 {
		peri++
	}

	if j == 0 || grid[i][j-1] == 0 {
		peri++
	}

	if i == len(grid)-1 || grid[i+1][j] == 0 {
		peri++
	}

	if j == len(grid[0])-1 || grid[i][j+1] == 0 {
		peri++
	}

	return peri

}

func islandPerimeter(grid [][]int) int {

	m, n := len(grid), len(grid[0])
	peri := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				peri += cellPerimeter(grid, i, j)
			}
		}
	}

	return peri
}
