package main

import "sort"

/*
[47]全排列 II
//给定一个可包含重复数字的序列，返回所有不重复的全排列。
//
// 示例:
//
// 输入: [1,1,2]
//输出:
//[
//  [1,1,2],
//  [1,2,1],
//  [2,1,1]
//]
// Related Topics 回溯算法
// 👍 339 👎 0
*/

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	var result [][]int
	sort.Ints(nums)
	permuteUniqueHelper(&result, nums, 0)
	return result
}
func permuteUniqueHelper(result *[][]int, nums []int, i int) {
	n := len(nums)
	if i == n-1 {
		temp := make([]int, n)
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}
	//从k开始调换顺序达到全排列效果

	for k := i; k < n; k++ {
		if k != i && nums[k] == nums[i] {
			continue
		}
		nums[k], nums[i] = nums[i], nums[k]
		permuteUniqueHelper(result, nums, i+1)
	}
	//顺序还原
	for k := n - 1; k > i; k-- {
		nums[k], nums[i] = nums[i], nums[k]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
