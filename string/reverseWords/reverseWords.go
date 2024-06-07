package reversewords

import (
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	newWords := make([]string, 0, len(words))
	for i := len(words) - 1; i >= 0; i-- {
		if words[i] == "" || words[i] == " " {
			continue
		}
		newWords = append(newWords, words[i])
	}
	return strings.Join(newWords, " ")
}
