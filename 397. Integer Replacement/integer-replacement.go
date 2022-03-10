package main

// https://leetcode.com/problems/integer-replacement/

func Min(a, b int) int {
	if b > a {
		return a
	}
	return b
}

func MinSteps(n int, cache map[int]int) int {
	if steps, ok := cache[n]; ok {
		return steps
	}

	if n%2 == 0 {
		cache[n] = MinSteps(n/2, cache) + 1
	} else {
		cache[n] = Min(MinSteps(n-1, cache), MinSteps(n+1, cache)) + 1
	}
	return cache[n]
}

func integerReplacement(n int) int {
	cache := make(map[int]int)

	cache[1] = 0
	cache[2] = 1

	return MinSteps(n, cache)
}
