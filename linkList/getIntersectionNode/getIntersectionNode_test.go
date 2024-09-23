package getintersectionnode

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "示例1",
			args: args{
				headA: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 8,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
				headB: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 8,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 5,
									},
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
