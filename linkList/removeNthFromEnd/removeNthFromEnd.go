package removenthfromend

import linklist "github.com/gz4z2b/algorithm/linkList"

// 19. 删除链表的倒数第 N 个结点

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

type ListNode = linklist.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}
	slow := &ListNode{
		Next: head,
	}
	fakeHead := slow
	fast := slow
	gap := 0
	for fast.Next != nil {
		if gap == n {
			slow = slow.Next
		} else {
			gap++
		}
		fast = fast.Next
	}
	if gap >= n && slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return fakeHead.Next
}
