package main

/*
[329]矩阵中的最长递增路径
//给定一个整数矩阵，找出最长递增路径的长度。
//
// 对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外（即不允许环绕）。
//
// 示例 1:
//
// 输入: nums =
//[
//  [9,9,4],
//  [6,6,8],
//  [2,1,1]
//]
//输出: 4
//解释: 最长递增路径为 [1, 2, 6, 9]。
//
// 示例 2:
//
// 输入: nums =
//[
//  [3,4,5],
//  [3,2,6],
//  [2,2,1]
//]
//输出: 4
//解释: 最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。
//
// Related Topics 深度优先搜索 拓扑排序 记忆化
*/

//leetcode submit region begin(Prohibit modification and deletion)
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	//记录长度，至少1，初始为0代表没有被遍历过
	visited := make([][]int, len(matrix))
	for i := range visited {
		visited[i] = make([]int, len(matrix[0]))
	}
	maxValue := 0
	for i := range matrix {
		for j := range matrix[i] {
			if visited[i][j] == 0 {
				maxValue = max(maxValue, dfs(i, j, matrix, visited))
			}
		}
	}
	return maxValue
}
func dfs(i, j int, matrix, visited [][]int) int {
	if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) {
		return 0
	}
	//已遍历则直接返回
	if visited[i][j] > 0 {
		return visited[i][j]
	}

	maxValue := 0
	//查看四周是否有小于当前节点的点，如果有则递归遍历该节点
	if i-1 >= 0 && matrix[i-1][j] < matrix[i][j] {
		maxValue = max(maxValue, dfs(i-1, j, matrix, visited))
	}
	if j-1 >= 0 && matrix[i][j-1] < matrix[i][j] {
		maxValue = max(maxValue, dfs(i, j-1, matrix, visited))
	}
	if i+1 < len(matrix) && matrix[i+1][j] < matrix[i][j] {
		maxValue = max(maxValue, dfs(i+1, j, matrix, visited))
	}
	if j+1 < len(matrix[0]) && matrix[i][j+1] < matrix[i][j] {
		maxValue = max(maxValue, dfs(i, j+1, matrix, visited))
	}
	visited[i][j] = maxValue + 1
	return maxValue + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
