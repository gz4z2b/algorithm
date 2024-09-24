package hash

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{
			name: "用例1",
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			name: "用例2",
			args: args{
				strs: []string{""},
			},
			want: [][]string{
				{""},
			},
		},
		{
			name: "用例3",
			args: args{
				strs: []string{"a"},
			},
			want: [][]string{
				{"a"},
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.args.strs)
			assert.True(t, compare2DArrays(tt.want, got))
		})
	}
}

// 将二维数组展平成一维数组
func flattenArray(arr [][]string) []string {
	var result []string
	for _, row := range arr {
		result = append(result, row...)
	}
	return result
}

// 比较两个二维数组，忽略元素顺序
func compare2DArrays(arr1, arr2 [][]string) bool {
	flatArr1 := flattenArray(arr1)
	flatArr2 := flattenArray(arr2)

	// 如果展平后的数组长度不同，则返回 false
	if len(flatArr1) != len(flatArr2) {
		return false
	}

	// 排序两个一维数组
	sort.Strings(flatArr1)
	sort.Strings(flatArr2)

	// 逐个比较元素
	for i := range flatArr1 {
		if flatArr1[i] != flatArr2[i] {
			return false
		}
	}

	return true
}
