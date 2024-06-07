package caculateResult

// 为了深入了解这些生物群体的生态特征，你们进行了大量的实地观察和数据采集。数组 arrayA 记录了各个生物群体数量数据，其中 arrayA[i] 表示第 i 个生物群体的数量。请返回一个数组 arrayB，该数组为基于数组 arrayA 中的数据计算得出的结果，其中 arrayB[i] 表示将第 i 个生物群体的数量从总体中排除后的其他数量的乘积。

func statisticalResult(arrayA []int) []int {
	result := make([]int, len(arrayA))
	if len(arrayA) == 0 {
		return result
	}
	result[0] = 1
	zeroCount := 0
	uniqVal := arrayA[0]
	for i := 1; i <= len(arrayA)-1; i++ {
		if arrayA[i] == 0 {
			zeroCount++
		} else {
			uniqVal *= arrayA[i]
		}
		result[0] *= arrayA[i]
	}
	if zeroCount > 1 {
		return result
	}
	for i := 1; i <= len(arrayA)-1; i++ {
		if arrayA[i] != 0 {
			result[i] = result[i-1] / arrayA[i] * arrayA[i-1]
		} else {
			result[i] = uniqVal
		}
	}
	return result
}
