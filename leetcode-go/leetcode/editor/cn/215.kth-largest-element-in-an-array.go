package main

/*
[215]数组中的第K个最大元素
//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 示例 1:
//
// 输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//
//
// 示例 2:
//
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
//
// 说明:
//
// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
// Related Topics 堆 分治算法
*/

//leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	left := 0
	right := len(nums) - 1
	for {
		if left >= right {
			return nums[left]
		}
		p := findKtyLargestHelp(nums, left, right)
		if p+1 == k {
			return nums[p]
		}
		if p+1 < k {
			left = p + 1
		} else {
			right = p - 1
		}

	}
}
func findKtyLargestHelp(nums []int, left, right int) int {
	pivot := nums[right]
	for i := left; i < right; i++ {
		//大的在左，小的在右，那么第K大即为下标K-1
		if nums[i] > pivot {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
	nums[left], nums[right] = nums[right], nums[left]
	return left
}

//leetcode submit region end(Prohibit modification and deletion)
