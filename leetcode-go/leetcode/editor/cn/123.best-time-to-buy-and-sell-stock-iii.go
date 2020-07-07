package main

/*
[123]买卖股票的最佳时机 III
//给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
//
// 注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
// 示例 1:
//
// 输入: [3,3,5,0,0,3,1,4]
//输出: 6
//解释: 在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
//     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
//
// 示例 2:
//
// 输入: [1,2,3,4,5]
//输出: 4
//解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
//     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
//
//
// 示例 3:
//
// 输入: [7,6,4,3,1]
//输出: 0
//解释: 在这个情况下, 没有交易完成, 所以最大利润为 0。
// Related Topics 数组 动态规划
// 👍 432 👎 0
*/

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	pLen := len(prices)
	if pLen == 0 {
		return 0
	}

	result := 0
	profit := make([][3][2]int, pLen)
	/**
	 *三维数组
	 *profit[ii][kk][jj]
	 *ii 第 ii 天， kk 股票操作了几次 ， jj 是否有股票
	 *最多可以完成 两笔 交易： kk 可以为 0 1 2 次操作 ， jj可以为 0 ，1
	 **/
	/**
	 *第一天 初始化
	 *第 1 天 操作 k 次 没有股票，所以初始值为 0
	 *第 1 天 操作i 次 有股票， 所以初始值为 - prices[0]
	 **/
	profit[0][0][0], profit[0][0][1] = 0, -prices[0]
	profit[0][1][0], profit[0][1][1] = 0, -prices[0]
	profit[0][2][0], profit[0][2][1] = 0, -prices[0]

	//注意 买 卖 都进行一次算一次操作 k + 1,单独 买入 不算完成一次操作
	for i := 1; i < pLen; i++ {
		//第 i 天 0 次交易 没有股票最大利润 = 第 i-1 天 0 次交易 没有股票最大利润
		profit[i][0][0] = profit[i-1][0][0]
		//第 i 天 0 次交易 有股票最大利润 = max(第 i-1 天 0 次交易 有股票最大利润 , 第 i-1 天 0 次交易 无股票最大利润 - 当天股票价格prices[i]（买入）)
		profit[i][0][1] = max(profit[i-1][0][1], profit[i-1][0][0]-prices[i])

		//# 第 i 天 1 次交易 无股票最大利润 = max(第 i-1 天 1次交易 无股票最大利润 , 第 i-1 天 0 次交易 有股票最大利润 + 当天股票价格prices[i]（卖出）)
		profit[i][1][0] = max(profit[i-1][1][0], profit[i-1][0][1]+prices[i])
		// # 第 i 天 1 次交易 有股票最大利润 = max(第 i-1 天 1 次交易 有股票最大利润 , 第 i-1 天 1 次交易 无股票最大利润 - 当天股票价格prices[i]（买入）)
		profit[i][1][1] = max(profit[i-1][1][1], profit[i-1][1][0]-prices[i])

		//# 第 i 天 2 次交易 无股票最大利润 = max(第 i-1 天 2次交易 无股票最大利润 , 第 i-1 天 1 次交易 有股票最大利润 + 当天股票价格prices[i]（卖出）)
		profit[i][2][0] = max(profit[i-1][2][0], profit[i-1][1][1]+prices[i])

	}
	result = max(profit[pLen-1][0][0], max(profit[pLen-1][1][0], profit[pLen-1][2][0]))
	return result
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

//leetcode submit region end(Prohibit modification and deletion)
