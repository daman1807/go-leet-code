package main

// https://leetcode.com/problems/minimum-index-sum-of-two-lists/

func findRestaurant(list1 []string, list2 []string) []string {

	cost := make(map[string]int)

	for i := 0; i < len(list1); i++ {
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				if _, ok := cost[list1[i]]; !ok {
					cost[list1[i]] = i + j
				}
			}
		}
	}

	min := int(^uint(0) >> 1)
	res := make([]string, 0)

	for _, v := range cost {
		if v < min {
			min = v
		}
	}

	for k, v := range cost {
		if v == min {
			res = append(res, k)
		}
	}

	return res
}
