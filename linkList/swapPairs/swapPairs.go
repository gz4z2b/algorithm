package swappairs

import linklist "github.com/gz4z2b/algorithm/linkList"

// 24. 两两交换链表中的节点

// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

type ListNode = linklist.ListNode

func swapPairs(head *ListNode) *ListNode {
	cur := head
	haveChangeHead := false
	var pre *ListNode
	for cur != nil && cur.Next != nil {
		slow := cur
		fast := cur.Next
		slow.Next, fast.Next = fast.Next, slow
		if pre != nil {
			pre.Next = fast
		}
		cur, pre = slow.Next, slow
		if !haveChangeHead {
			head = fast
			haveChangeHead = true
		}
	}
	return head
}
