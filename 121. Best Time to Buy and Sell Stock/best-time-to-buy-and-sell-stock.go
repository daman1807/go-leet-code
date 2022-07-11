package main

import "math"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := math.MaxInt

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}
