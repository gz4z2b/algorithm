package removeelement

// 844. 比较含退格的字符串
// 给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。

// 注意：如果对空文本输入退格字符，文本继续为空。

func backspaceCompare(s string, t string) bool {
	sResult := getBackspaceString(s)
	tResult := getBackspaceString(t)
	return sResult == tResult
}

func getBackspaceString(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if len(result) > 0 {
				result = result[0 : len(result)-1]
			}
		} else {
			result += string(s[i])
		}
	}
	return result
}
