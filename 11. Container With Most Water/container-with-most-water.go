package main

//https://leetcode.com/problems/container-with-most-water/

func Min(x int, y int) int {
	if x >= y {
		return y
	}
	return x

}

func maxArea(height []int) int {
	var area, maximumArea int

	x, y := 0, len(height)-1
	maximumArea = Min(height[x], height[y]) * (y - x)

	for _ = range height {
		if x == y {
			break
		}

		if height[x] < height[y] {
			x += 1
		} else {
			y -= 1
		}

		area = Min(height[x], height[y]) * (y - x)
		if area > maximumArea {
			maximumArea = area
		}

	}
	return maximumArea
}
