package spiralorder

import (
	"reflect"
	"testing"
)

func Test_spiralArray(t *testing.T) {
	type args struct {
		array [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "方形",
			args: args{
				array: [][]int{
					{1, 2, 3},
					{8, 9, 4},
					{7, 6, 5},
				},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "长方形",
			args: args{
				array: [][]int{
					{1, 2, 3, 4},
					{10, 11, 12, 5},
					{9, 8, 7, 6},
				},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			name: "一个",
			args: args{
				array: [][]int{{1}},
			},
			want: []int{1},
		},
		{
			name: "0个",
			args: args{
				array: [][]int{},
			},
			want: []int{},
		},
		{
			name: "长条",
			args: args{
				array: [][]int{{1, 2, 3}},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
