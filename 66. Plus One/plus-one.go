package main

//https://leetcode.com/problems/plus-one/

func plusOne(digits []int) []int {
	result := make([]int, len(digits)+1)
	carry := false

	result[0] = 1

	for i := len(digits) - 1; i >= 0; i -= 1 {
		if i == len(digits)-1 {
			digits[i] += 1
		}
		if carry {
			digits[i] += 1
		}

		if digits[i] > 9 {
			carry = true
			digits[i] = 0
		} else {
			carry = false
		}
		result[i+1] = digits[i]
	}

	if carry {
		return result
	}
	return result[1:]
}
