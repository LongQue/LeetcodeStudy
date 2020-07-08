package main

import "sort"

/*
[386]字典序排数
//给定一个整数 n, 返回从 1 到 n 的字典顺序。
//
// 例如，
//
// 给定 n =1 3，返回 [1,10,11,12,13,2,3,4,5,6,7,8,9] 。
//
// 请尽可能的优化算法的时间复杂度和空间复杂度。 输入的数据 n 小于等于 5,000,000。
// 👍 103 👎 0
*/

//leetcode submit region begin(Prohibit modification and deletion)
func lexicalOrder(n int) []int {
	result := make([]int, n)
	dfs(&result, 0, n)
	return result
}
func dfs(result *[]int, start, end int) {
	for i := 0; i < 10 && start+i <= end; i++ {
		if start+i == 0 {
			continue
		}
		*result = append(*result, start+i)
		if (start+i)*10 <= end {
			dfs(result, (start+i)*10, end)
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
