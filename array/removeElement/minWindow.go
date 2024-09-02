package removeelement

// 76. 最小覆盖子串
// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

func minWindow(s string, t string) string {
	var result string
	needUseDict := make(map[rune]int, len(t))
	needNum := 0
	for _, v := range t {
		needUseDict[v]++
		needNum++
	}
	begin := 0
	haveDict := make(map[rune]int, len(needUseDict))
	for end := 0; end <= len(s)-1; end++ {
		if _, exist := needUseDict[rune(s[end])]; exist {
			haveDict[rune(s[end])]++
			needNum--
		}
		if checkCover(haveDict, needUseDict) {
			for ; true; begin++ {
				tmpStr := s[begin : end+1]
				if len(result) == 0 || len(tmpStr) < len(result) {
					result = tmpStr
				}
				if _, exist := needUseDict[rune(s[begin])]; exist {
					haveDict[rune(s[begin])]--
					if haveDict[rune(s[begin])] < needUseDict[rune(s[begin])] {
						begin++
						break
					}
				}
			}
		}

	}
	return result
}

func checkCover(my, target map[rune]int) bool {
	for k, v := range target {
		if count, exist := my[k]; exist {
			if count < v {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
