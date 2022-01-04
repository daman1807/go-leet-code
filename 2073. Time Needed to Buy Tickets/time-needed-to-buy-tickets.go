package main

// https://leetcode.com/problems/time-needed-to-buy-tickets/

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func timeRequiredToBuy(tickets []int, k int) int {
	res := 0
	for i := 0; i <= k; i++ {
		res += Min(tickets[i], tickets[k])
	}
	for i := k + 1; i < len(tickets); i++ {
		res += Min(tickets[i], tickets[k]-1)
	}
	return res
}
