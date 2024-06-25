package search

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: []int{3, 4},
		},
		{
			name: "示例2",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: []int{-1, -1},
		},
		{
			name: "示例3",
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: []int{-1, -1},
		},
		{
			name: "示例4",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 5,
			},
			want: []int{0, 0},
		},
		{
			name: "示例5",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 10,
			},
			want: []int{5, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLeftBorder(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "示例1",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 7,
			},
			want: 1,
		},
		{
			name: "示例2",
			args: args{
				nums:   []int{7, 7, 8, 8, 10},
				target: 7,
			},
			want: 0,
		},
		{
			name: "示例3",
			args: args{
				nums:   []int{7, 7, 8, 8, 10},
				target: 6,
			},
			want: -1,
		},
		{
			name: "示例4",
			args: args{
				nums:   []int{7, 7, 8, 8, 10},
				target: 11,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLeftBorder(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findLeftBorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRightBorder(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "示例1",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 7,
			},
			want: 2,
		},
		{
			name: "示例2",
			args: args{
				nums:   []int{7, 7, 8, 8, 10},
				target: 7,
			},
			want: 1,
		},
		{
			name: "示例3",
			args: args{
				nums:   []int{7, 7, 8, 8, 10},
				target: 6,
			},
			want: -1,
		},
		{
			name: "示例4",
			args: args{
				nums:   []int{7, 7, 8, 8, 10},
				target: 11,
			},
			want: -1,
		},
		{
			name: "示例5",
			args: args{
				nums:   []int{7, 7, 8, 8, 10, 10},
				target: 10,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRightBorder(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findRightBorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
