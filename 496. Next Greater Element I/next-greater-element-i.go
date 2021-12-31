package main

// https://leetcode.com/problems/next-greater-element-i/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))

	for i := 0; i < len(nums1); i++ {
		found := false
		nge := -1
		for j := 0; j < len(nums2); j++ {
			if nums2[j] == nums1[i] {
				found = true
			}
			if found {
				if nums2[j] > nums1[i] {
					nge = nums2[j]
					break
				}
			}
		}
		res[i] = nge

	}
	return res
}
