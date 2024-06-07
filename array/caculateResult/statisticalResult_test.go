package caculateResult

import (
	"reflect"
	"testing"
)

func Test_statisticalResult(t *testing.T) {
	type args struct {
		arrayA []int
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
				arrayA: []int{2, 4, 6, 8, 10},
			},
			want: []int{1920, 960, 640, 480, 384},
		},
		{
			name: "含0的",
			args: args{
				arrayA: []int{2, 0, 6, 8, 10},
			},
			want: []int{0, 960, 0, 0, 0},
		},
		{
			name: "含多个0的",
			args: args{
				arrayA: []int{2, 0, 6, 0, 10},
			},
			want: []int{0, 0, 0, 0, 0},
		},
		{
			name: "空数组",
			args: args{
				arrayA: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := statisticalResult(tt.args.arrayA); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("statisticalResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
