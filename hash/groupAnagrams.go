package hash

// 49. 字母异位词分组

// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

func groupAnagrams(strs []string) [][]string {
	charMaps := make(map[[26]int][]string)
	for _, str := range strs {
		charMap := makeCharMap(str)
		if _, exist := charMaps[charMap]; exist {
			charMaps[charMap] = append(charMaps[charMap], str)
		} else {
			charMaps[charMap] = []string{str}
		}
	}
	result := make([][]string, 0)
	for _, charSlice := range charMaps {
		result = append(result, charSlice)
	}
	return result
}

func makeCharMap(str string) [26]int {
	result := [26]int{}
	for _, char := range str {
		result[char-'a']++
	}
	return result
}
