package removeelement

// 977. 有序数组的平方
// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

func sortedSquares(nums []int) []int {
	// 先找出中间数
	result := make([]int, 0, len(nums))
	idx := -1
	for i := 0; i < len(nums) && nums[i] < 0; i++ {
		idx = i
	}

	for slow, fast := idx, idx+1; slow >= 0 || fast < len(nums); {
		if slow < 0 {
			result = append(result, nums[fast]*nums[fast])
			fast++
		} else if fast >= len(nums) {
			result = append(result, nums[slow]*nums[slow])
			slow--
		} else if nums[fast]*nums[fast] < nums[slow]*nums[slow] {
			result = append(result, nums[fast]*nums[fast])
			fast++
		} else {
			result = append(result, nums[slow]*nums[slow])
			slow--
		}
	}

	return result
}
