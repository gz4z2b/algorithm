package removeelement

// 283. 移动零

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。

func moveZeroes(nums []int) {
	zeroIdx := 0

	for idx := zeroIdx; idx < len(nums); idx++ {
		if nums[idx] != 0 {
			nums[idx], nums[zeroIdx] = nums[zeroIdx], nums[idx]
			zeroIdx++
		}
	}
}
