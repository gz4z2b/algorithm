package removeelement

import "math"

// 977. 有序数组的平方
// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

func sortedSquares(nums []int) []int {
	// 先找出中间数
	idx := len(nums) - 1
	for ; idx > 0; idx-- {
		if nums[idx] < 0 {
			break
		}
	}
	idx++
	result := []int{
		int(math.Pow(float64(nums[idx]), 2)),
	}

	addOffset, minOffset := 1, 1
	for idx+addOffset <= len(nums) && idx-minOffset >= 0 {
		minNum := math.Abs(float64(nums[idx+addOffset]))
		addNum := math.Abs(float64(nums[idx-minOffset]))
		if minNum < addNum {
			result = append(result, int(math.Pow(minNum, 2)))
			minOffset++
		} else {
			result = append(result, int(math.Pow(addNum, 2)))
			addOffset++
		}
	}
	if idx+addOffset <= len(nums) {
		for ; idx+addOffset <= len(nums); addOffset++ {
			result = append(result, int(math.Pow(float64(nums[idx+addOffset]), 2)))
		}
	} else {
		for ; idx-minOffset >= 0; minOffset++ {
			result = append(result, int(math.Pow(float64(nums[idx-minOffset]), 2)))
		}
	}
	return result
}
