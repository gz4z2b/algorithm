package removeelement

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "示例1",
			args: args{
				s: "ab#c",
				t: "ad#c",
			},
			want: true,
		},
		{
			name: "示例2",
			args: args{
				s: "ab##",
				t: "c#d#",
			},
			want: true,
		},
		{
			name: "示例3",
			args: args{
				s: "a#c",
				t: "b",
			},
			want: false,
		},
		{
			name: "示例4",
			args: args{
				s: "y#fo##f",
				t: "y#f#o##f",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := backspaceCompare(tt.args.s, tt.args.t)
			if got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
