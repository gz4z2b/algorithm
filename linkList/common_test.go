package linklist

import (
	"reflect"
	"testing"
)

func TestToSlice(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "正常",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "空",
			args: args{
				head: nil,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSlice(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToLinkList(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "正常",
			args: args{
				s: []int{1, 2, 3},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
		{
			name: "空",
			args: args{
				s: []int{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLinkList(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToLinkList() = %v, want %v", got, tt.want)
			}
		})
	}
}
