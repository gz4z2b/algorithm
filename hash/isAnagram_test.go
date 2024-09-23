package hash

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "用例1",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "用例2",
			args: args{
				s: "rat",
				t: "cat",
			},
			want: false,
		},
		{
			name: "用例3",
			args: args{
				s: "abc",
				t: "ab",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
