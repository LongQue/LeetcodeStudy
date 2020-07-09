package main

/*
[718]最长重复子数组
//给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。
//
//
//
// 示例：
//
// 输入：
//A: [1,2,3,2,1]
//B: [3,2,1,4,7]
//输出：3
//解释：
//长度最长的公共子数组是 [3, 2, 1] 。
//
//
//
//
// 提示：
//
//
// 1 <= len(A), len(B) <= 1000
// 0 <= A[i], B[i] < 100
//
// Related Topics 数组 哈希表 二分查找 动态规划
// 👍 281 👎 0
*/

//leetcode submit region begin(Prohibit modification and deletion)
func findLength(A []int, B []int) int {
	n, m := len(A), len(B)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	ans := 0
	//从后面开始算，dp[i][j] = dp[i + 1][j + 1] + 1，所以需要n+1,m+1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if A[i] == B[j] {
				//
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			if ans < dp[i][j] {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
