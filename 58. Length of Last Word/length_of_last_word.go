package main

// https://leetcode.com/problems/length-of-last-word/

func lengthOfLastWord(s string) int {
	end := len(s) - 1
	res := 0
	for end >= 0 {
		if string(s[end]) == " " {

		} else {
			for end >= 0 {
				if string(s[end]) == " " {
					return res
				} else {
					res += 1
				}
				end -= 1
			}
		}
		end -= 1
	}
	return res

}
