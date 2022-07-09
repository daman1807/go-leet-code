package main

// https://leetcode.com/problems/reverse-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var prev *ListNode
	prev = nil

	for head != nil {
		dummy := &ListNode{Val: head.Val}
		dummy.Next = prev

		prev = dummy
		head = head.Next
	}

	return prev

}
