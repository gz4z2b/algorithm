package myatoi

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "用例一",
			args: args{s: "42"},
			want: 42,
		},
		{
			name: "用例二",
			args: args{s: "-42"},
			want: -42,
		},
		{
			name: "用例三",
			args: args{s: "1337c0d3"},
			want: 1337,
		},
		{
			name: "用例四",
			args: args{s: "0-1"},
			want: 0,
		},
		{
			name: "用例五",
			args: args{s: "words and 987"},
			want: 0,
		},
		{
			name: "用例六",
			args: args{s: " 987"},
			want: 987,
		},
		{
			name: "用例七",
			args: args{s: "-91283472332"},
			want: -2147483648,
		},
		{
			name: "用例八",
			args: args{s: "91283472332"},
			want: 2147483647,
		},
		{
			name: "用例九",
			args: args{s: "-+12"},
			want: 0,
		},
		{
			name: "用例十",
			args: args{s: "20000000000000000000"},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
