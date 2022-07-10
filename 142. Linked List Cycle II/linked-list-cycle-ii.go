package main

// https://leetcode.com/problems/linked-list-cycle-ii/

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	slow, fast := head, head

	for {
		if slow == nil || fast == nil || fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	slow = head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast

}
