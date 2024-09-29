package hash

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "用例1",
			args: args{
				s: "cbaebabacd",
				p: "abc",
			},
			want: []int{0, 6},
		},
		{
			name: "用例2",
			args: args{
				s: "abab",
				p: "ab",
			},
			want: []int{0, 1, 2},
		},
		{
			name: "用例3",
			args: args{
				s: "aaaa",
				p: "a",
			},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "用例4",
			args: args{
				s: "aaaaaaaaaa",
				p: "aaaaaaaaaaaaa",
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceCompaire(t *testing.T) {
	numsA := [3]int{1, 2, 3}
	numsB := [3]int{1, 2, 3}
	if numsA == numsB {
		t.Log("true")
	} else {
		t.Log("false")
	}
}
