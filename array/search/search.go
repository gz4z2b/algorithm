package search

// 704. 二分查找
// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
// 你可以假设 nums 中的所有元素是不重复的。
// n 将在 [1, 10000]之间。
// nums 的每个元素都将在 [-9999, 9999]之间。

func search(nums []int, target int) int {
	begin, end := 0, len(nums)-1
	idx := begin + (end-begin)/2
	for begin <= end {
		if nums[idx] == target {
			return idx
		} else if nums[idx] < target {
			begin = idx + 1
		} else {
			end = idx - 1
		}
		idx = begin + (end-begin)/2
	}
	return -1
}
