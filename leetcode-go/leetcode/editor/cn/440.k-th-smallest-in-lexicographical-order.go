package main

/*
[440]字典序的第K小数字
//给定整数 n 和 k，找到 1 到 n 中字典序第 k 小的数字。
//
// 注意：1 ≤ k ≤ n ≤ 109。
//
// 示例 :
//
//
//输入:
//n: 13   k: 2
//
//输出:
//10
//
//解释:
//字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。
//
//
*/

//leetcode submit region begin(Prohibit modification and deletion)
func findKthNumber(n int, k int) int {
	cur := 1
	k -= 1
	for k > 0 {
		steps := findKthNumberHelper(n, cur, cur+1)
		if steps <= k {
			k -= steps
			cur += 1
		} else {
			k -= 1
			cur *= 10
		}
	}
	return cur
}
func findKthNumberHelper(n, i, j int) int {
	step := 0
	for i <= n {
		step += min(j, n+1) - i
		i *= 10
		j *= 10
	}
	return step
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
