package main

// https://leetcode.com/problems/add-binary

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1

	var res string
	carry := 0
	val := 0
	for {
		if i < 0 && j < 0 {
			break
		}
		val = carry
		if i >= 0 {
			val += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			val += int(b[j] - '0')
			j--
		}

		if val == 3 {
			carry = 1
			res += "1"
		} else if val == 2 {
			carry = 1
			res += "0"
		} else {
			carry = 0
			if val == 0 {
				res += "0"
			} else {
				res += "1"
			}
		}
	}

	if carry == 1 {
		res += "1"
	}

	return Reverse(res)
}
