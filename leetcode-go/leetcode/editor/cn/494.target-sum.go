package main

import (
	"math"
)

/*
[494]目标和
//给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选
//择一个符号添加在前面。
//
// 返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
//
//
//
// 示例：
//
// 输入：nums: [1, 1, 1, 1, 1], S: 3
//输出：5
//解释：
//
//-1+1+1+1+1 = 3
//+1-1+1+1+1 = 3
//+1+1-1+1+1 = 3
//+1+1+1-1+1 = 3
//+1+1+1+1-1 = 3
//
//一共有5种方法让最终目标和为3。
//
//
//
//
// 提示：
//
//
// 数组非空，且长度不会超过 20 。
// 初始的数组的和不会超过 1000 。
// 保证返回的最终结果能被 32 位整数存下。
//
// Related Topics 深度优先搜索 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func findTargetSumWays(nums []int, S int) int {
	//整数总和
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if math.Abs(float64(sum)) < math.Abs(float64(S)) {
		return 0
	}
	//  正负0
	t := sum*2 + 1
	row := len(nums)
	arrays := make([][]int, row)
	for i := range arrays {
		arrays[i] = make([]int, t)
	}
	if nums[0] == 0 {
		arrays[0][sum] = 2
	} else {
		arrays[0][sum+nums[0]] = 1
		arrays[0][sum-nums[0]] = 1
	}

	for i := 1; i < row; i++ {
		for j := 0; j < t; j++ {
			var l int
			var r int
			if j-nums[i] >= 0 {
				l = j - nums[i]
			} else {
				l = 0
			}
			if j+nums[i] < t {
				r = j + nums[i]
			} else {
				r = 0
			}
			arrays[i][j] = arrays[i-1][l] + arrays[i-1][r]
		}
	}
	return arrays[row-1][sum+S]
}

//leetcode submit region end(Prohibit modification and deletion)
