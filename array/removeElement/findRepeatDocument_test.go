package removeelement

import (
	"testing"
)

func Test_findRepeatDocument(t *testing.T) {
	type args struct {
		documents []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "错误1",
			args: args{
				documents: []int{3, 4, 2, 1, 1, 0},
			},
			want: []int{1},
		},
		{
			name: "示例1",
			args: args{
				documents: []int{2, 5, 3, 0, 5, 0},
			},
			want: []int{5, 0},
		},
		{
			name: "没找到",
			args: args{
				documents: []int{2, 4, 3, 1, 5, 0},
			},
			want: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := findRepeatDocument(tt.args.documents); inSlice(tt.want, got) {
			// 	t.Errorf("findRepeatDocument() = %v, want %v", got, tt.want)
			// }
			if got := findRepeatDocumentV1(tt.args.documents); !inSlice(tt.want, got) {
				t.Errorf("findRepeatDocumentV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func inSlice(list []int, a int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
