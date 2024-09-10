package reverselist

import linklist "github.com/gz4z2b/algorithm/linkList"

// 206. 反转链表

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

func reverseList(head *linklist.ListNode) *linklist.ListNode {
	cur := head
	var pre *linklist.ListNode
	for cur != nil {
		if pre != nil {
			cur, cur.Next, pre = cur.Next, pre, cur
		} else {
			pre = cur
			cur = cur.Next
			pre.Next = nil
		}
	}
	return pre
}
