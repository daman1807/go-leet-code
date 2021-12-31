package main

// https://leetcode.com/problems/jewels-and-stones/

func numJewelsInStones(jewels string, stones string) int {
	res := 0
	for _, jewel := range jewels {
		for _, stone := range stones {
			if jewel == stone {
				res += 1
			}
		}
	}
	return res
}
