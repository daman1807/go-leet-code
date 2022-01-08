package main

/// https://leetcode.com/problems/length-of-last-word/

func lengthOfLastWord(s string) int {
	high := len(s) - 1
	res := 0

	for high >= 0 {
		if s[high] == ' ' {
			high--
		} else {
			break
		}
	}

	for high >= 0 {
		if s[high] == ' ' {
			break
		} else {
			res++
		}
		high--
	}
	return res
}
