package main

// https://leetcode.com/problems/integer-to-roman/

func intToRoman(num int) (s string) {
	literals := [...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	index := 0
	for num != 0 {
		if num-values[index] >= 0 {
			num -= values[index]
			s += literals[index]
		} else {
			index++
		}
	}
	return
}
