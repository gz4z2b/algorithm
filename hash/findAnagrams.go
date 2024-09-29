package hash

// 438. 找到字符串中所有字母异位词

// 给定两个字符串 s 和 p，找到 s 中所有 p 的
// 异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	result := make([]int, 0, len(s)-len(p))
	for i := 0; i <= len(s)-len(p); i++ {
		if isAnagram2(s[i:i+len(p)], p) {
			result = append(result, i)
		}
	}
	return result
}
