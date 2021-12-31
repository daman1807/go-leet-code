package main

import "strings"

// https://leetcode.com/problems/keyboard-row/

func contains(words []string, char string) bool {
	for _, word := range words {
		if char == word {
			return true
		}
	}
	return false

}

func valid(word string, row []string) bool {
	for _, char := range word {
		if !contains(row, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}

func findWords(words []string) []string {

	firstRow := []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}
	secondRow := []string{"a", "s", "d", "f", "g", "h", "j", "k", "l"}
	thirdRow := []string{"z", "x", "c", "v", "b", "n", "m"}

	var result []string

	for _, word := range words {
		if valid(word, firstRow) {
			result = append(result, word)
		} else if valid(word, secondRow) {
			result = append(result, word)
		} else if valid(word, thirdRow) {
			result = append(result, word)
		}
	}
	return result
}
