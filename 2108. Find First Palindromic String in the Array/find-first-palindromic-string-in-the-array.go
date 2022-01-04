package main

//  https://leetcode.com/problems/find-first-palindromic-string-in-the-array/

func isPalindrome(word string) bool {
	low, high := 0, len(word)-1
	for low < high {
		if word[low] != word[high] {
			return false
		}
		low++
		high--
	}
	return true
}

func firstPalindrome(words []string) string {
	res := ""
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return res
}
