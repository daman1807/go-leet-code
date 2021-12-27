package main

//https://leetcode.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0

	for i < m && j < n {
		if nums1[i] > nums2[j] {
			nums1[i], nums2[j] = nums2[j], nums1[i]
			for k := j; k < n-1; k += 1 {
				if nums2[k] > nums2[k+1] {
					nums2[k], nums2[k+1] = nums2[k+1], nums2[k]
				}
			}
		}
		i += 1
	}

	for i = m; i < len(nums1); i++ {
		nums1[i] = nums2[i-m]
	}

}
