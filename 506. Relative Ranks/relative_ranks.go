package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/relative-ranks/

func findRelativeRanks(score []int) []string {
	medals := make([]string, len(score))
	rankMap := make(map[int]int)
	dummy := make([]int, len(score))

	for i := 0; i < len(score); i += 1 {
		dummy[i] = score[i]
	}

	sort.Ints(dummy)

	for i := len(dummy) - 1; i >= 0; i-- {
		rankMap[dummy[i]] = len(dummy) - i
	}

	for idx, val := range score {
		switch rankMap[val] {
		case 1:
			medals[idx] = "Gold Medal"
		case 2:
			medals[idx] = "Silver Medal"
		case 3:
			medals[idx] = "Bronze Medal"
		default:
			medals[idx] = fmt.Sprintf("%d", rankMap[val])
		}
	}

	return medals
}
