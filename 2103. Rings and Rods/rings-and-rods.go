package main

import "strconv"

//https://leetcode.com/problems/rings-and-rods/

func countPoints(rings string) int {

	var bars [10][3]int

	result := 0
	index := 0
	color := ""

	for i := 0; i < len(rings); i += 2 {

		color = string(rings[i])
		index, _ = strconv.Atoi(string(rings[i+1]))

		switch color {
		case string('R'):
			bars[index][0] = 1
		case string('B'):
			bars[index][1] = 1
		case string('G'):
			bars[index][2] = 1
		}
	}

	for i := 0; i < 10; i += 1 {
		if bars[i][0]+bars[i][1]+bars[i][2] == 3 {
			result += 1
		}
	}
	return result
}
