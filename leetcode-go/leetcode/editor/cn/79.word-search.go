package main

/*
[79]单词搜索
//给定一个二维网格和一个单词，找出该单词是否存在于网格中。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
// 示例:
//
// board =
//[
//  ['A','B','C','E'],
//  ['S','F','C','S'],
//  ['A','D','E','E']
//]
//
//给定 word = "ABCCED", 返回 true
//给定 word = "SEE", 返回 true
//给定 word = "ABCB", 返回 false
//
//
//
// 提示：
//
//
// board 和 word 中只包含大写和小写英文字母。
// 1 <= board.length <= 200
// 1 <= board[i].length <= 200
// 1 <= word.length <= 10^3
//
// Related Topics 数组 回溯算法
*/

//leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])
	//遍历每一个元素，不用管找不找得到第一个，在后面的函数再处理
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if existHelper(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}
func existHelper(board [][]byte, word string, i, j, index int) bool {
	//index是下标，如果下标等于长度说明已找全
	if index == len(word) {
		return true
	}
	//处理越界问题
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return false
	}
	//如果下一个不等于字符串下一个返回 false
	if board[i][j] != word[index] {
		return false
	}
	//先把值取出来
	temp := board[i][j]
	//设为空格，标记以遍历过
	board[i][j] = byte(' ')
	//遍历四个方向，下标+1
	if existHelper(board, word, i+1, j, index+1) {
		return true
	}
	if existHelper(board, word, i-1, j, index+1) {
		return true
	}
	if existHelper(board, word, i, j+1, index+1) {
		return true
	}
	if existHelper(board, word, i, j-1, index+1) {
		return true
	}
	//回溯把值设置回来
	board[i][j] = temp
	//四个方向都没有适合的返回false
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
