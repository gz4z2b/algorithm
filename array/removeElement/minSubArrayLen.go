package removeelement

import "math"

// 209. 长度最小的子数组
// 给定一个含有 n 个正整数的数组和一个正整数 target 。

// 找出该数组中满足其总和大于等于 target 的长度最小的
// 子数组
//  [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

func minSubArrayLen(target int, nums []int) int {

	sum, result := 0, math.MaxInt
	for fast, slow := 0, 0; fast <= len(nums)-1; fast++ {
		sum += nums[fast]
		for sum >= target {
			subLenth := fast - slow + 1
			if subLenth < result {
				result = subLenth
			}
			sum -= nums[slow]
			slow++
		}

	}
	if result == math.MaxInt {
		return 0
	} else {
		return result
	}
}
