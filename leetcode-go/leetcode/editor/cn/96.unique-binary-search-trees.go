package main

/*
[96]不同的二叉搜索树
//给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
//
// 示例:
//
// 输入: 3
//输出: 5
//解释:
//给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
//
//   1         3     3      2      1
//    \       /     /      / \      \
//     3     2     1      1   3      2
//    /     /       \                 \
//   2     1         2                 3
// Related Topics 树 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	//动态规划
	//n构成的数据是 1 2 3 4...n 大小已排序，对于二叉树而言，不同排序的左子树和右子树的个数，仅与节点个数有关
	// 1 2 3 4 5 当2作为root时，左边1个节点，右边3个节点，两者组成的树相乘等于2作为root的树的个数，若1-5分别做root并累加，即为n=5时的结果
	//遍历n个数据，当前数字作为root，左边的个数作为左子树，右边作右子树
	dp := make([]int, n+1)
	//节点数为0,1时可组成一个树
	dp[0] = 1
	dp[1] = 1
	//动态规划，从i=2算起，计算dp[2]的值
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			//j-1  i-j左右节点的树个数
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
