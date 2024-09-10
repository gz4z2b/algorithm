package mylinkedlist

// 707. 设计链表

// 你可以选择使用单链表或者双链表，设计并实现自己的链表。

// 单链表中的节点应该具备两个属性：val 和 next 。val 是当前节点的值，next 是指向下一个节点的指针/引用。

// 如果是双向链表，则还需要属性 prev 以指示链表中的上一个节点。假设链表中的所有节点下标从 0 开始。

// 实现 MyLinkedList 类：

// MyLinkedList() 初始化 MyLinkedList 对象。
// int get(int index) 获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1 。
// void addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
// void addAtTail(int val) 将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
// void addAtIndex(int index, int val) 将一个值为 val 的节点插入到链表中下标为 index 的节点之前。如果 index 等于链表的长度，那么该节点会被追加到链表的末尾。如果 index 比长度更大，该节点将 不会插入 到链表中。
// void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为 index 的节点。

type MyLinkedList struct {
	Head   *MyLinkedListNode
	Tail   *MyLinkedListNode
	Length int
}

type MyLinkedListNode struct {
	Val  int
	Next *MyLinkedListNode
	Pre  *MyLinkedListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) getNodeByIndex(index int) *MyLinkedListNode {
	if index > this.Length-1 {
		return nil
	}
	cur := this.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.getNodeByIndex(index)
	if cur == nil {
		return -1
	}

	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := MyLinkedListNode{
		Val:  val,
		Next: this.Head,
	}
	if this.Head != nil {
		this.Head.Pre = &node
	}
	this.Head = &node
	if this.Tail == nil {
		this.Tail = &node
	}
	this.Length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := MyLinkedListNode{
		Val: val,
		Pre: this.Tail,
	}
	if this.Tail != nil {
		this.Tail.Next = &node
	}
	this.Tail = &node
	if this.Head == nil {
		this.Head = &node
	}
	this.Length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Length {
		return
	}
	if index == this.Length {
		this.AddAtTail(val)
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	cur := this.getNodeByIndex(index)

	node := &MyLinkedListNode{
		Val:  val,
		Next: cur,
		Pre:  cur.Pre,
	}
	cur.Pre.Next = node
	cur.Pre = node
	this.Length++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this.getNodeByIndex(index)
	if cur == nil {
		return
	}
	if cur.Pre != nil {
		cur.Pre.Next = cur.Next
	} else {
		this.Head = cur.Next
	}
	if cur.Next != nil {
		cur.Next.Pre = cur.Pre
	} else {
		this.Tail = cur.Pre
	}
	this.Length--
}

func (this *MyLinkedList) ToSlice() []int {
	cur := this.Head
	result := make([]int, 0, this.Length)
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}
	return result
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
