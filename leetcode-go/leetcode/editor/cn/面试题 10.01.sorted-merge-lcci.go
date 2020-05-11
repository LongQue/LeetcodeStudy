package main

/*
[面试题 10.01]合并排序的数组
//给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。
//
// 初始化 A 和 B 的元素数量分别为 m 和 n。
//
// 示例:
//
// 输入:
//A = [1,2,3,0,0,0], m = 3
//B = [2,5,6],       n = 3
//
//输出: [1,2,2,3,5,6]
//
// 说明:
//
//
// A.length == n + m
//
// Related Topics 数组 双指针
*/

//leetcode submit region begin(Prohibit modification and deletion)
func merge(A []int, m int, B []int, n int) {
	//双指针指向AB数据结尾，比较大小排在A数组后面
	AIndex := m - 1
	BIndex := n - 1
	idx := len(A) - 1
	for AIndex >= 0 && BIndex >= 0 {
		if A[AIndex] > B[BIndex] {
			A[idx] = A[AIndex]
			AIndex--
		} else {
			A[idx] = B[BIndex]
			BIndex--
		}
		idx--
	}
	for BIndex >= 0 {
		A[idx] = B[BIndex]
		BIndex--
		idx--
	}
}

//leetcode submit region end(Prohibit modification and deletion)
