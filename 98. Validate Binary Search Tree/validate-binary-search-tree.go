package main

import "math"

// https://leetcode.com/problems/validate-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValid(root *TreeNode, min, max float64) bool {
	if root == nil {
		return true
	} else {
		val := float64(root.Val)
		return val > min && val < max && isValid(root.Left, min, val) && isValid(root.Right, val, max)
	}
}
func isValidBST(root *TreeNode) bool {
	return isValid(root, math.Inf(-1), math.Inf(1))
}
