package spiralorder

// LCR 146. 螺旋遍历二维数组
// 给定一个二维数组 array，请返回「螺旋遍历」该数组的结果。

// 螺旋遍历：从左上角开始，按照 向右、向下、向左、向上 的顺序 依次 提取元素，然后再进入内部一层重复相同的步骤，直到提取完所有元素。

func spiralArray(array [][]int) []int {
	if len(array) == 0 {
		return []int{}
	}
	result := make([]int, 0, len(array)*len(array[0]))

	xBegin, xEnd, yBegin, yEnd, x, y := 0, len(array[0])-1, 0, len(array)-1, 0, 0

	for xBegin < xEnd || yBegin < yEnd {
		// 向右
		for x, y = xBegin, yBegin; x < xEnd; x++ {
			result = append(result, array[y][x])
		}
		// 向下
		for ; y < yEnd; y++ {
			result = append(result, array[y][x])
		}
		// 向左
		for ; x > xBegin; x-- {
			result = append(result, array[y][x])
		}
		// 向上
		for ; y > yBegin; y-- {
			result = append(result, array[y][x])
		}

		xBegin++
		yBegin++
		xEnd--
		yEnd--
	}
	if len(result) != len(array)*len(array[0]) {
		result = append(result, array[y][x])
	}
	return result

}
