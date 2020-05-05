package main

/*
[322]零钱兑换
//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回
// -1。
//
//
//
// 示例 1:
//
// 输入: coins = [1, 2, 5], amount = 11
//输出: 3
//解释: 11 = 5 + 5 + 1
//
// 示例 2:
//
// 输入: coins = [2], amount = 3
//输出: -1
//
//
//
// 说明:
//你可以认为每种硬币的数量是无限的。
// Related Topics 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {
	//辅助队列，记录金额逐渐增加时，所需要的硬币数，amount+1，使下面的i即为当前金额数
	temp := make([]int, amount+1)
	temp[0] = 0

	for i := 1; i < len(temp); i++ {
		//失败标记
		temp[i] = amount + 1
		//遍历硬币面额，计算可从前面第[i-coin]个加上coin得到当前金额
		for _, v := range coins {
			//小于0说明不需要
			if i-v < 0 {
				continue
			}
			temp[i] = min(temp[i], temp[i-v]+1)
		}
	}
	if temp[amount] == amount+1 {
		return -1
	}
	return temp[amount]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
