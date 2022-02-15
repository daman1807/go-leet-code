package main

// https://leetcode.com/problems/sqrtx/

func mySqrt(x int) int {

	if x == 1 {
		return 1
	}

	mid, sqr := 0, 0
	low, high := 0, x

	for low <= high {

		mid = (low + high) / 2
		sqr = mid * mid

		if sqr == x {
			return mid
		} else if sqr > x {
			high = mid - 1
		} else {
			if (mid+1)*(mid+1) <= x {
				low = mid + 1
			} else {
				return mid
			}
		}
	}
	return mid
}
