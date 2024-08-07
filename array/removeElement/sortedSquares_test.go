package removeelement

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "示例1",
			args: args{
				nums: []int{-4, -1, 0, 3, 10},
			},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "示例2",
			args: args{
				nums: []int{-7, -3, 2, 3, 11},
			},
			want: []int{4, 9, 9, 49, 121},
		},
		{
			name: "示例3",
			args: args{
				nums: []int{-7, -3, -2, -1, 1, 2, 3, 11},
			},
			want: []int{1, 1, 4, 4, 9, 9, 49, 121},
		},
		{
			name: "示例4",
			args: args{
				nums: []int{-7, -3, -2, -1, 1, 11},
			},
			want: []int{1, 1, 4, 9, 49, 121},
		},
		{
			name: "示例5",
			args: args{
				nums: []int{-7, -3, -2, -1, 1},
			},
			want: []int{1, 1, 4, 9, 49},
		},
		{
			name: "示例7",
			args: args{
				nums: []int{-1},
			},
			want: []int{1},
		},
		{
			name: "示例6",
			args: args{
				nums: []int{1},
			},
			want: []int{1},
		},
		{
			name: "示例8",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
		{
			name: "示例9",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{1, 4, 9},
		},
		{
			name: "示例10",
			args: args{
				nums: []int{-3, -2, -1},
			},
			want: []int{1, 4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
