package search

// 35. 搜索插入位置
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。
func searchInsert(nums []int, target int) int {
	right, left := len(nums), 0
	idx := (right + left) / 2

	for right > left {
		if nums[idx] == target {
			return idx
		} else if nums[idx] > target {
			right = idx
		} else {
			left = idx + 1
		}
		idx = (right + left) / 2
	}
	return idx
}
