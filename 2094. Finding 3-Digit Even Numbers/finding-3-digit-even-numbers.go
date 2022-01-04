package main

import "sort"

// https://leetcode.com/problems/finding-3-digit-even-numbers/

func checkOrigin(num int, counter map[int]int) bool {

	for num > 0 {
		_, ok := counter[num%10]
		if ok && counter[num%10] > 0 {
			counter[num%10]--
		} else {
			return false
		}
		num /= 10
	}
	return true
}

func findEvenNumbers(digits []int) []int {
	var res []int

	counter := make(map[int]int)
	dup := make(map[int]int)

	for i := 0; i < len(digits); i++ {
		_, ok := counter[digits[i]]
		if ok {
			counter[digits[i]]++
			dup[digits[i]]++
		} else {
			counter[digits[i]] = 1
			dup[digits[i]] = 1
		}
	}

	sort.Ints(digits)

	for i := 100; i < 1000; i += 2 {
		counterDup := make(map[int]int)
		for k, v := range counter {
			counterDup[k] = v
		}
		if checkOrigin(i, counterDup) {
			res = append(res, i)
		}
	}

	return res
}
