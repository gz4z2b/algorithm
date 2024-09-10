package mylinkedlist

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMyLinkedList(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		// TODO: Add test cases.
		{
			name: "用例1",
			test: func(t *testing.T) {
				list := Constructor()
				t.Log(111)
				list.AddAtHead(1)
				list.AddAtTail(3)
				list.AddAtIndex(1, 2)
				assert.Equal(t, []int{1, 2, 3}, list.ToSlice())
				assert.Equal(t, 2, list.Get(1))
				list.DeleteAtIndex(1)
				assert.Equal(t, []int{1, 3}, list.ToSlice())
				assert.Equal(t, 3, list.Get(1))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.test(t)
		})
	}
}
