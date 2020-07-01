package main

/*
[209]长度最小的子数组
//给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组，并返回其长度。如果不存在符合条件的连续子数组，返回
// 0。
//
//
//
// 示例：
//
// 输入：s = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的连续子数组。
//
//
//
//
// 进阶：
//
//
// 如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
//
// Related Topics 数组 双指针 二分查找
*/

//leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLen(s int, nums []int) int {
	//模拟滑动窗口
	var queue []int
	//记录窗口内的值
	cur := 0
	//设置最大值
	res := len(nums) + 1
	for _, v := range nums {
		queue = append(queue, v)
		cur += v
		//如果大于等于s，从左边移除
		for cur >= s {
			res = min(len(queue), res)
			cur -= queue[0]
			queue = queue[1:]
		}
	}
	//可能存在所有总和达不到s的情况
	if res == len(nums)+1 {
		return 0
	}
	return res
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
