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
	//十叉树先序遍历

	//先走第一步，走到1位置
	cur := 1
	k -= 1
	for k > 0 {
		//先计算1-2的区间
		steps := findKthNumberHelper(n, cur, cur+1)
		//小于步数，则减去，再计算2-3的区间
		if steps <= k {
			k -= steps
			cur += 1
		} else {
			//大于步数则需要细化，计算10-11的区间
			//k-=1 从节点1走到节点10需要一步
			k -= 1
			cur *= 10
		}
	}
	return cur
}

//用于计算两个同级相邻节点间相隔多少节点
func findKthNumberHelper(n, i, j int) int {
	//j用于比i大1，0-123456789-10,11,12,13
	step := 0
	//第二层 i=1 j=2 相隔1
	//第三层取n和j=20最大值 减去i=10获得相隔5，则原始i，j即1,2相隔5位
	//从节点1走到节点2需要5步，即第五个元素是2
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
