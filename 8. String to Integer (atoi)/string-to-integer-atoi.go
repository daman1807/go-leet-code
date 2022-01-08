package main

// https://leetcode.com/problems/string-to-integer-atoi/

import "math"

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	sign := 1
	res := 0
	low, high := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			low++
		} else {
			break
		}
	}

	if low >= len(s) {
		return res
	}

	if s[low] == '-' {
		sign = -1
		low++
	} else if s[low] == '+' {
		sign = 1
		low++
	}

	for i := low; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			high = i + 1
		} else {
			break
		}
	}

	if low > high {
		return res
	}

	for {
		if low < len(s)-1 && s[low] == '0' {
			low++
		} else {
			break
		}
	}

	pow := 1

	for low < high {
		res += int(s[high-1]-'0') * pow
		if res > math.MaxInt32 || pow > math.MaxInt32 {
			if sign > 0 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		high--
		pow *= 10
	}
	return sign * res
}
