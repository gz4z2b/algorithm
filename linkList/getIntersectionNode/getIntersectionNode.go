package getintersectionnode

import linklist "github.com/gz4z2b/algorithm/linkList"

// 面试题 02.07. 链表相交

// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

type ListNode = linklist.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB, curA, curB := 0, 0, headA, headB
	for curA != nil || curB != nil {
		if curA != nil {
			lenA++
			curA = curA.Next
		}
		if curB != nil {
			lenB++
			curB = curB.Next
		}
	}

	curA = headA
	curB = headB
	if lenA > lenB {
		for i := 1; i <= lenA-lenB; i++ {
			curA = curA.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			curB = curB.Next
		}
	}

	for curA != nil && curB != nil {
		if curA == curB {
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}
	return nil
}
