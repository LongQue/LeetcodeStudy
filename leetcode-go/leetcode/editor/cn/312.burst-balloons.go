package main

/*
[312]戳气球
//有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
//
// 现在要求你戳破所有的气球。每当你戳破一个气球 i 时，你可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的
//left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。
//
// 求所能获得硬币的最大数量。
//
// 说明:
//
//
// 你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
// 0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
//
//
// 示例:
//
// 输入: [3,1,5,8]
//输出: 167
//解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
//     coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
//
// Related Topics 分治算法 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func maxCoins(nums []int) int {
	// Store the input's length.
	n := len(nums)
	// Padding 1s to head and tail of nums.
	nums = append(nums, 1)
	nums = append([]int{1}, nums...)

	// Create 2D-DP with n+2 width and height.
	// c[i][j] represents the max coins from i to j.
	c := make([][]int, n+2)
	for i := range c {
		c[i] = make([]int, n+2)
	}

	// l is the length of subarray. We start with l= 1, end with l = n.
	for l := 1; l <= n; l++ {
		// i is the start point in this subarray.
		for i := 1; i <= n-l+1; i++ {
			// j is the subarray's end.
			j := i + l - 1
			// k is the break point to separate.
			for k := i; k <= j; k++ {
				c[i][j] = max(c[i][j], c[i][k-1]+nums[i-1]*nums[k]*nums[j+1]+c[k+1][j])
			}
		}
	}

	return c[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
