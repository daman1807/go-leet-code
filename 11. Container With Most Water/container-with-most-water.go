package main

//https://leetcode.com/problems/container-with-most-water/

func Min(x int, y int) int {
	if x >= y {
		return y
	}
	return x

}

func maxArea(height []int) int {
	var res, area int

	low, high := 0, len(height)-1

	for low < high {

		area = Min(height[low], height[high]) * (high - low)
		if area > res {
			res = area
		}

		if height[low] > height[high] {
			high--
		} else {
			low++
		}

	}
	return res
}
