package main

/*
[128]最长连续序列
//给定一个未排序的整数数组，找出最长连续序列的长度。
//
// 要求算法的时间复杂度为 O(n)。
//
// 示例:
//
// 输入: [100, 4, 200, 1, 3, 2]
//输出: 4
//解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
// Related Topics 并查集 数组
*/

//leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	// k nums中的元素   v 以该元素作为边界的连续子序列有多长
	m := make(map[int]int)
	ret := 0
	for _, v := range nums {
		//没有出现过的节点才需要计算
		if m[v] == 0 {
			//查看以左右元素作为边界的连续子序列有多长
			left, right := m[v-1], m[v+1]
			//将左右连接起来的长度
			cur := left + right + 1
			//比较取最大值
			ret = max(cur, ret)
			//左边有left长，  v-left定位到左边边界，修改包含当前cur之后的长度
			m[v-left] = cur
			//同上
			m[v+right] = cur
			//修改自己的长度
			m[v] = cur
		}
	}
	return ret
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
