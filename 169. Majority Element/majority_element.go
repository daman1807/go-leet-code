package main

import "sort"

//https://leetcode.com/problems/majority-element/

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[int(len(nums)/2)]
}
