package main

// https://leetcode.com/problems/binary-tree-level-order-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func add(val int, ans *[][]int, depth int) {
	if len(*ans) <= depth {
		*ans = append(*ans, make([]int, 0))
	}
	(*ans)[depth] = append((*ans)[depth], val)

}
func traverse(root *TreeNode, ans *[][]int, depth int) {
	if root == nil {
		return
	} else {
		add(root.Val, ans, depth)
		defer traverse(root.Right, ans, depth+1)
		defer traverse(root.Left, ans, depth+1)
	}
}

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)

	if root == nil {
		return ans
	}

	traverse(root, &ans, 0)

	return ans

}
