package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ToSlice(head *ListNode) []int {
	cur := head
	result := make([]int, 0)
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}
	return result
}

func ToLinkList(s []int) *ListNode {
	var (
		result *ListNode
		cur    *ListNode
	)
	for _, v := range s {
		node := &ListNode{
			Val: v,
		}
		if cur != nil {
			cur.Next = node
			cur = node
		} else {
			result, cur = node, node
		}
	}
	return result
}
