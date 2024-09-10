package reverselist

import (
	"reflect"
	"testing"

	linklist "github.com/gz4z2b/algorithm/linkList"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *linklist.ListNode
	}
	tests := []struct {
		name string
		args args
		want *linklist.ListNode
	}{
		// TODO: Add test cases.
		{
			name: "正常",
			args: args{
				head: &linklist.ListNode{
					Val: 1,
					Next: &linklist.ListNode{
						Val: 2,
						Next: &linklist.ListNode{
							Val: 3,
						},
					},
				},
			},
			want: &linklist.ListNode{
				Val: 3,
				Next: &linklist.ListNode{
					Val: 2,
					Next: &linklist.ListNode{
						Val: 1,
					},
				},
			},
		},
		{
			name: "空",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
