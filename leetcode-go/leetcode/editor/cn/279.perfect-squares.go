package main

/*
[279]完全平方数
//给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
//
// 示例 1:
//
// 输入: n = 12
//输出: 3
//解释: 12 = 4 + 4 + 4.
//
// 示例 2:
//
// 输入: n = 13
//输出: 2
//解释: 13 = 4 + 9.
// Related Topics 广度优先搜索 数学 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func numSquares(n int) int {
	//提前预置平方数
	square := []int{0, 1, 4}
	squIndex := 1
	//n+1为了方便从 temp[0]=0
	temp := make([]int, n+1)
	temp[0] = 0
	//temp 的索引
	i := 1
	for i <= n {
		//用上面的举例，如果i>=4这说明之后的i比如大于square的最大值，需要加入新的平方数
		if i >= square[squIndex+1] {
			//squere增加
			squIndex++
			square = append(square, (squIndex+1)*(squIndex+1))
		}
		//设置为最大值n
		temp[i] = n
		//遍历平方数数组，直到1
		for j := squIndex; j >= 1; j-- {
			temp[i] = min(temp[i], temp[i-square[j]]+1)
		}
		i++
	}
	return temp[n]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
