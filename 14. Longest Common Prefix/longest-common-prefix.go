package main

// https://leetcode.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	var res string
	j := 0

	if len(strs) == 1 {
		return strs[0]
	}

	for {
		for i := 1; i < len(strs); i++ {
			if j > len(strs[i])-1 || j > len(strs[0])-1 {
				return res
			}
			if strs[i][j] != strs[0][j] {
				return res
			}
		}
		res += string(strs[0][j])
		j += 1
	}
	return res
}
