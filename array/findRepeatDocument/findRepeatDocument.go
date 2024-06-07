package findrepeatdocument

// 设备中存有 n 个文件，文件 id 记于数组 documents。若文件 id 相同，则定义为该文件存在副本。请返回任一存在副本的文件 id。
// 0 ≤ documents[i] ≤ n-1
// 2 <= n <= 100000
func findRepeatDocument(documents []int) int {
	docMap := make(map[int]bool, len(documents))
	for _, v := range documents {
		if _, exist := docMap[v]; exist {
			return v
		} else {
			docMap[v] = true
		}
	}
	return -1
}

func findRepeatDocumentV1(documents []int) int {
	for idx := 0; idx < len(documents); {
		if idx != documents[documents[idx]] {
			if documents[documents[idx]] == documents[idx] {
				return documents[idx]
			}
			documents[idx], documents[documents[idx]] = documents[documents[idx]], documents[idx]
		} else {
			idx++
		}
	}
	return -1
}
