package main

// https://leetcode.com/problems/first-bad-version/
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	if n == 0 {
		return 0
	}

	low, mid, high := 1, n/2, n

	for low+1 < high {
		mid = low + (high-low)/2
		if isBadVersion(mid) {
			high = mid
		} else {
			low = mid
		}
	}

	if isBadVersion(low) {
		return low
	} else {
		return high
	}
}
