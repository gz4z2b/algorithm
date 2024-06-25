package search

// 69. x 的平方根
// 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

// 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

// 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。

func mySqrt(x int) int {
	left, right := 0, x
	result := 0
	for left <= right {
		result = (left + right) / 2
		if result*result < x {
			left = result + 1
		} else if result*result > x {
			right = result - 1
		} else {
			return result
		}
	}
	if result*result > x {
		return result - 1
	} else {
		return result
	}
}
