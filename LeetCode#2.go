package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	ret := ListNode{0,nil}
	next1 := l1
	next2 := l2

	cur := &ret
	add := 0

	for next1 != nil || next2 != nil {
		var left int
		var right int
		if next1 != nil {
			left = next1.Val
		} else {
			left = 0
		}
		if next2 != nil {
			right = next2.Val
		} else {
			right = 0
		}
		total := left+right+add
		add = total/10

		cur.Next = &ListNode{total%10, nil}
		cur = cur.Next

		if next1 != nil {
			next1 = next1.Next
		}
		if next2 != nil {
			next2 = next2.Next
		}
	}
	if add != 0 {
		cur.Next = &ListNode{add, nil}
	}

	return ret.Next
}