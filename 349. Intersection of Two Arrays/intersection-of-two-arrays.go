package main

import "sort"

// https://leetcode.com/problems/intersection-of-two-arrays/

func intersection(nums1 []int, nums2 []int) []int {

	var res []int

	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			// removing duplicates here
			if i > 0 && j > 0 {
				if nums1[i] == nums1[i-1] {
					i += 1
					j += 1
					continue
				}
			}
			res = append(res, nums1[i])
			i += 1
			j += 1
		} else if nums1[i] > nums2[j] {
			j += 1
		} else {
			i += 1
		}
	}
	return res
}
