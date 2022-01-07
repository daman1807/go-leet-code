package main

// https://leetcode.com/problems/roman-to-integer/

func romanToInt(s string) int {
	var res int
	var symbolVal = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else if symbolVal[string(s[i])] < symbolVal[string(s[i-1])] {
			res += symbolVal[string(s[i-1])] * count
			count = 1
		} else {
			res -= symbolVal[string(s[i-1])] * count
		}
	}

	res += count * symbolVal[string(s[len(s)-1])]
	return res
}
