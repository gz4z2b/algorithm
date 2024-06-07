package spiralorder

// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

func spiralOrder(matrix [][]int) []int {
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	x, y := left, top
	result := make([]int, 0)
	for top <= bottom && left <= right {
		for y, x = top, left; x <= right; x++ {
			result = append(result, matrix[y][x])
		}
		top++
		if top > bottom {
			break
		}

		for y, x = top, right; y <= bottom; y++ {
			result = append(result, matrix[y][x])
		}
		right--
		if right < left {
			break
		}

		for y, x = bottom, right; x >= left; x-- {
			result = append(result, matrix[y][x])
		}
		bottom--
		if bottom < top {
			break
		}

		for y, x = bottom, left; y >= top; y-- {
			result = append(result, matrix[y][x])
		}
		left++
	}
	return result
}
