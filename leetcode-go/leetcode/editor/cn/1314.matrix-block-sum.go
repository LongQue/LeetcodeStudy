package main

import "math"

/*
[1314]矩阵区域和
//给你一个 m * n 的矩阵 mat 和一个整数 K ，请你返回一个矩阵 answer ，其中每个 answer[i][j] 是所有满足下述条件的元素 ma
//t[r][c] 的和：
//
//
// i - K <= r <= i + K, j - K <= c <= j + K
// (r, c) 在矩阵内。
//
//
//
//
// 示例 1：
//
// 输入：mat = [[1,2,3],[4,5,6],[7,8,9]], K = 1
//输出：[[12,21,16],[27,45,33],[24,39,28]]
//
//
// 示例 2：
//
// 输入：mat = [[1,2,3],[4,5,6],[7,8,9]], K = 2
//输出：[[45,45,45],[45,45,45],[45,45,45]]
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n, K <= 100
// 1 <= mat[i][j] <= 100
//
// Related Topics 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func matrixBlockSum(mat [][]int, K int) [][]int {
	rows, cols := len(mat), len(mat[0])

	preSum := make([][]int, rows+1)
	for i := range preSum {
		preSum[i] = make([]int, cols+1)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			preSum[i+1][j+1] = preSum[i+1][j] + preSum[i][j+1] - preSum[i][j] + mat[i][j]
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			row1 := math.Max(float64(i-K), 0)
			col1 := math.Max(float64(j-K), 0)

			row2 := math.Min(float64(i+K), float64(rows-1))
			col2 := math.Min(float64(j+K), float64(cols-1))
			mat[i][j] = sumRegion(int(row1), int(col1), int(row2), int(col2), preSum)
		}
	}
	return mat
}

func sumRegion(row1, col1, row2, col2 int, preSum [][]int) int {
	return preSum[row2+1][col2+1] - preSum[row1][col2+1] - preSum[row2+1][col1] + preSum[row1][col1]
}

//leetcode submit region end(Prohibit modification and deletion)
