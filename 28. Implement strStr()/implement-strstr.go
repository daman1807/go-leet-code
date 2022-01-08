package main

// https://leetcode.com/problems/implement-strstr/

func strStr(haystack string, needle string) int {

	if len(haystack) == 0 && len(needle) == 0{
		return 0
	}

	if len(haystack) < len(needle){
		return -1
	}

	res := -1
	for i:=0; i<=len(haystack) - len(needle); i++{
		res = i
		for j:=0; i+j < len(haystack) && j<len(needle); j++{
			if haystack[i+j] != needle[j]{
				res = -1
				break
			}
		}
		if res != -1 {
			break
		}
	}
	return res
}