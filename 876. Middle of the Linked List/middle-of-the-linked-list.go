package main

// https://leetcode.com/problems/middle-of-the-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	curr := head
	count := 0

	for curr != nil {
		curr = curr.Next
		count++
	}

	for i := 0; i < count/2; i++ {
		head = head.Next
	}
	return head

}
