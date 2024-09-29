package hash

import "math"

// 202. 快乐数
// 编写一个算法来判断一个数 n 是不是快乐数。

// 「快乐数」 定义为：

// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 如果这个过程 结果为 1，那么这个数就是快乐数。
// 如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

func isHappy(n int) bool {
	haveExist := make(map[int]bool)
	haveExist[n] = true
	for {
		n = caculateEachDigit(n)
		if n == 1 {
			return true
		}
		if _, exist := haveExist[n]; exist {
			return false
		}
		haveExist[n] = true
	}
}

func caculateEachDigit(n int) int {
	result := 0
	for n > 0 {
		digt := float64(n % 10)
		result += int(math.Pow(digt, 2))
		n /= 10
	}
	return result
}
