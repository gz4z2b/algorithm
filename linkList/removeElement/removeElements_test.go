package removeelement

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func nodeListToArr(head *ListNode) []int {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
func arrToNodeList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	result := &ListNode{Val: arr[0]}
	idx := result
	for i := 1; i <= len(arr)-1; i++ {
		idx.Next = &ListNode{
			Val: arr[i],
		}
		idx = idx.Next
	}
	return result
}

func Test_removeElements(t *testing.T) {
	type args struct {
		head []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例 1",
			args: args{
				head: []int{1, 2, 6, 3, 4, 5, 6},
				val:  6,
			},
			want: []int{
				1, 2, 3, 4, 5,
			},
		},
		{
			name: "示例 2",
			args: args{
				head: []int{},
				val:  1,
			},
			want: []int{},
		},
		{
			name: "示例 3",
			args: args{
				head: []int{7, 7, 7, 7},
				val:  7,
			},
			want: []int{},
		},
		{
			name: "示例 4",
			args: args{
				head: []int{1, 2, 2, 1},
				val:  2,
			},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := arrToNodeList(tt.args.head)
			got := removeElements(head, tt.args.val)
			assert.Equal(t, nodeListToArr(got), tt.want)
		})
	}
}
