package main

/*
[面试题 08.11]硬币
//硬币。给定数量不限的硬币，币值为25分、10分、5分和1分，编写代码计算n分有几种表示法。(结果可能会很大，你需要将结果模上1000000007)
//
// 示例1:
//
//
// 输入: n = 5
// 输出：2
// 解释: 有两种方式可以凑成总金额:
//5=5
//5=1+1+1+1+1
//
//
// 示例2:
//
//
// 输入: n = 10
// 输出：4
// 解释: 有四种方式可以凑成总金额:
//10=10
//10=5+5
//10=5+1+1+1+1+1
//10=1+1+1+1+1+1+1+1+1+1
//
//
// 说明：
//
// 注意:
//
// 你可以假设：
//
//
// 0 <= n (总金额) <= 1000000
//
// Related Topics 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func waysToChange(n int) int {
	//https://leetcode-cn.com/problems/coin-lcci/solution/ying-bi-by-leetcode-solution/ 官方数学方法
	if n == 0 {
		return 0
	}
	temp := make([]int, n+1)
	temp[0] = 1
	arrays := []int{1, 5, 10, 25}
	//计算只用某种硬币时有多少种用法，然后在用其他硬币来算
	for _, v := range arrays {
		for i := v; i < n+1; i++ {
			temp[i] = (temp[i] + temp[i-v]) % 1000000007
		}
	}
	return temp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
