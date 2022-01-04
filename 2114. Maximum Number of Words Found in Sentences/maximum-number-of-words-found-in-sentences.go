package main

// https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/

func mostWordsFound(sentences []string) int {
	res := 0

	for i := 0; i < len(sentences); i++ {
		words := 1
		for j := 0; j < len(sentences[i]); j++ {
			if string(sentences[i][j]) == " " {
				words++
			}
		}
		if words > res {
			res = words
		}
	}
	return res
}
