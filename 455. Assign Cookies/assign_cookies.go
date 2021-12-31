package main

import "sort"

// https://leetcode.com/problems/assign-cookies/

func findContentChildren(g []int, s []int) int {

	res := 0

	i, j := 0, 0

	sort.Ints(g)
	sort.Ints(s)

	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			res += 1
			i++
			j++
		} else {
			j++
		}

	}
	return res
}
