package main

/*
[74]搜索二维矩阵
//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//
// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。
//
//
// 示例 1:
//
// 输入:
//matrix = [
//  [1,   3,  5,  7],
//  [10, 11, 16, 20],
//  [23, 30, 34, 50]
//]
//target = 3
//输出: true
//
//
// 示例 2:
//
// 输入:
//matrix = [
//  [1,   3,  5,  7],
//  [10, 11, 16, 20],
//  [23, 30, 34, 50]
//]
//target = 13
//输出: false
// Related Topics 数组 二分查找
// 👍 203 👎 0
*/

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	rows := len(matrix)
	cols := len(matrix[0])
	if rows > 0 && cols > 0 {
		row := 0
		col := cols - 1
		for row < rows && col >= 0 {
			if matrix[row][col] == target {
				return true
			} else if matrix[row][col] > target {
				col--
			} else {
				row++
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
