package main

import "math"

/*
[152]乘积最大子数组
//给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
//
//
// 示例 1:
//
// 输入: [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
//
//
// 示例 2:
//
// 输入: [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
// Related Topics 数组 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {
	//结果
	Max := math.MinInt64
	//记录当前最大最小值
	curMax := 1
	curMin := 1
	for _, v := range nums {
		if v < 0 {
			//如果遇到负数，最大最小值调换
			curMax, curMin = curMin, curMax
		}
		//求最大值，max 防止正*负
		curMax = max(curMax*v, v)
		curMin = min(curMin*v, v)

		Max = max(curMax, Max)
	}
	return Max
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
