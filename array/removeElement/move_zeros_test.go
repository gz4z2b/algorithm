package removeelement

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "示例1",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "示例2",
			args: args{
				nums: []int{0},
			},
			want: []int{0},
		},
		{
			name: "示例3",
			args: args{
				nums: []int{1, 2, 0, 3, 0},
			},
			want: []int{1, 2, 3, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			assert.Equal(t, tt.want, tt.args.nums)
		})
	}
}
