package search

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
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
				x: 4,
			},
			want: 2,
		},
		{
			name: "示例2",
			args: args{
				x: 8,
			},
			want: 2,
		},
		{
			name: "示例3",
			args: args{
				x: 0,
			},
			want: 0,
		},
		{
			name: "示例4",
			args: args{
				x: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
