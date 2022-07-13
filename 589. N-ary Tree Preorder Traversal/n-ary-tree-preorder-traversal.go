package main

// https://leetcode.com/problems/n-ary-tree-preorder-traversal/

type Node struct {
	Val      int
	Children []*Node
}

func traverse(root *Node, ans *[]int) {
	for _, child := range root.Children {
		*ans = append(*ans, child.Val)
		traverse(child, ans)
	}

}

func preorder(root *Node) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	ans = append(ans, root.Val)
	traverse(root, &ans)
	return ans
}
