package main

//https://leetcode.com/problems/number-of-segments-in-a-string

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}

	res, index := 0, 0
	start, end := 0, len(s)-1

	for {
		if string(s[start]) != " " || start >= len(s)-1 {
			break
		}
		start += 1
	}

	for {
		if string(s[end]) != " " || end <= 0 {
			break
		}
		end -= 1
	}

	if start == end {
		if string(s[start]) == " " {
			return 0
		} else {
			return 1
		}
	}

	index = start
	for index <= end {
		if string(s[index]) == " " && string(s[index-1]) != " " {
			res += 1
		}
		index += 1
	}

	if start < end {
		res += 1
	}

	return res
}
