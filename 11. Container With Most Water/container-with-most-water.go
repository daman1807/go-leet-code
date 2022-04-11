package main

//https://leetcode.com/problems/container-with-most-water

func Min(a, b int) int {
	if b > a {
		return a
	}
	return b
}

func maxArea(height []int) int {
	low, high := 0, len(height)-1
	area := 0
	max := 0

	for low < high {
		area = Min(height[low], height[high]) * (high - low)
		if area > max {
			max = area
		}

		if height[low] > height[high] {
			high--
		} else {
			low++
		}

	}
	return max
}
