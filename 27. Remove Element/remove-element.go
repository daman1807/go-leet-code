package main

//https://leetcode.com/problems/remove-element/solution/

func removeElement(nums []int, val int) int {
	var slow, fast int
	for fast=0; fast<len(nums); fast+=1{
		if nums[fast] != val{
			nums[slow] = nums[fast]
			slow+=1
		}
	}
	return slow
}
