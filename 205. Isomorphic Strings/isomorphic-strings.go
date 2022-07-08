package main

// https://leetcode.com/problems/isomorphic-strings/

func isIsomorphic(s string, t string) bool {
	chrMap := make(map[byte]byte)
	revMap := make(map[byte]byte)

	var chr, rev byte
	var ok bool

	for i := 0; i < len(s); i++ {
		chr, ok = chrMap[s[i]]
		if !ok {
			rev, ok = revMap[t[i]]
			if ok {
				if s[i] != rev {
					return false
				}
			} else {
				revMap[t[i]] = s[i]
			}
			chrMap[s[i]] = t[i]
		} else {
			if chr != t[i] {
				return false
			}
		}
	}
	return true
}
