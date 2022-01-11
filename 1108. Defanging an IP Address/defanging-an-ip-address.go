package main

// https://leetcode.com/problems/defanging-an-ip-address/

func defangIPaddr(address string) string {
	var res string

	for i := 0; i < len(address); i++ {
		if string(address[i]) == "." {
			res += "[.]"
		} else {
			res += string(address[i])
		}
	}
	return res
}
