package main

/*
[72]编辑距离
//给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
//
// 你可以对一个单词进行如下三种操作：
//
//
// 插入一个字符
// 删除一个字符
// 替换一个字符
//
//
//
//
// 示例 1：
//
// 输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
//
//
// 示例 2：
//
// 输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
//
// Related Topics 字符串 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func minDistance(word1 string, word2 string) int {
	row := len(word1)
	col := len(word2)
	//构建二维数组，长宽都+1,第0列第0行为了应付字符串为空的情况

	//下一步的变化等于左上[i-1,j-1]，左[i,j-1],上[i-1,j] 最小值+1
	dp := make([][]int, row+1)
	for i := range dp {
		if i == 0 {
			var temp []int
			for i := 0; i <= col; i++ {
				//初始化第0行
				temp = append(temp, i)
			}
			dp[i] = temp
		} else {
			dp[i] = make([]int, col+1)
		}
		//初始化第0列
		dp[i][0] = i
	}
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			//二维数组从0开始，所以要-1
			if word1[i-1] == word2[j-1] {
				//相等的情况下不需要任何操作，即为左上
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			}
		}
	}
	return dp[row][col]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
