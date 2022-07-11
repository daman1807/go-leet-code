package main

// https://leetcode.com/problems/longest-palindrome/

func longestPalindrome(s string) int {

	var pickedOdd bool
	var ans int

	counter := make(map[rune]int)

	for _, r := range s {
		counter[r]++
	}

	for _, c := range counter {
		if c%2 == 0 {
			ans += c
		} else {
			ans += c - 1
			if !pickedOdd {
				ans += 1
				pickedOdd = true
			}
		}
	}

	return ans
}
