package main

// https://leetcode.com/problems/distribute-candies

func distributeCandies(candyType []int) int {

	res := 0
	counts := make(map[int]bool, len(candyType))

	for i := 0; i < len(candyType); i++ {
		if res == len(candyType)/2 {
			break
		}

		_, ok := counts[candyType[i]]
		if !ok {
			res += 1
			counts[candyType[i]] = true
		}
	}
	return res
}
