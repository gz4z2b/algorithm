package spiralorder

// 59. 螺旋矩阵 II
// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

func generateMatrix(n int) [][]int {
	xBegin, xEnd, yBegin, yEnd := 0, n-1, 0, n-1
	x, y, num := xBegin, yBegin, 1
	result := make([][]int, n)
	for k := range result {
		result[k] = make([]int, n)
	}
	loop := 1
	for loop <= n/2 {
		for x, y = xBegin, yBegin; x < xEnd; x++ {
			result[y][x] = num
			num++
		}
		for x, y = xEnd, yBegin; y < yEnd; y++ {
			result[y][x] = num
			num++
		}
		for x, y = xEnd, yEnd; x > xBegin; x-- {
			result[y][x] = num
			num++
		}
		for x, y = xBegin, yEnd; y > yBegin; y-- {
			result[y][x] = num
			num++
		}
		xBegin++
		xEnd--
		yBegin++
		yEnd--
		loop++
	}
	if n%2 == 1 {
		result[yBegin][xBegin] = num
	}
	return result

}
