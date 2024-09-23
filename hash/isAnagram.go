package hash

// 242. 有效的字母异位词

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的
// 字母异位词

func isAnagram(s string, t string) bool {
	chatMap := make(map[rune]int)
	if len(s) != len(t) {
		return false
	}
	for k, v := range s {
		if _, exist := chatMap[v]; exist {
			chatMap[v]++
		} else {
			chatMap[v] = 1
		}
		tv := rune(t[k])
		if _, exist := chatMap[tv]; exist {
			chatMap[tv]--
		} else {
			chatMap[tv] = -1
		}
	}

	for _, v := range chatMap {
		if v != 0 {
			return false
		}
	}
	return true
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	chatMap := make([]int, 26)
	for k, v := range s {
		chatMap[v-'a']++
		chatMap[t[k]-'a']--
	}

	for _, v := range chatMap {
		if v != 0 {
			return false
		}
	}
	return true
}
