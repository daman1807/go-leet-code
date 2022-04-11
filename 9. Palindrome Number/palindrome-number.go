package main

// https://leetcode.com/problems/palindrome-number/

import "fmt"

func isPalindrome(x int) bool {
	s := fmt.Sprintf("%d", x)

	start, end := 0, len(s)-1

	for start < end {
		if s[start] != s[end] {
			return false
		}
		start += 1
		end -= 1
	}
	return true
}
