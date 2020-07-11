package main

import "math"

/*
[188]买卖股票的最佳时机 IV
//给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
//
// 注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
// 示例 1:
//
// 输入: [2,4,1], k = 2
//输出: 2
//解释: 在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
//
//
// 示例 2:
//
// 输入: [3,2,6,5,0,3], k = 2
//输出: 7
//解释: 在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
//     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3
//。
//
// Related Topics 动态规划
// 👍 248 👎 0
*/

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(k int, prices []int) int {
	if k > len(prices)/2 {
		return maxProfitInf(prices)
	}
	dp_i_0, dp_i_1 := make([]int, k+1), make([]int, k+1)
	for i := 0; i < k+1; i++ {
		dp_i_0[i] = 0
		dp_i_1[i] = math.MinInt64
	}
	for _, v := range prices {
		for i := k; i >= 1; i-- {
			dp_i_0[i] = max(dp_i_0[i], dp_i_1[i]+v)
			dp_i_1[i] = max(dp_i_1[i], dp_i_0[i-1]-v)
		}
	}
	return dp_i_0[k]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProfitInf(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	sum := 0
	plen := len(prices)
	for i := 1; i < plen; i++ {
		if prices[i]-prices[i-1] > 0 {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
