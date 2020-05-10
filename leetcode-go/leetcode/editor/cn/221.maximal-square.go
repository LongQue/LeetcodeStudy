package main

/*
[221]最大正方形
//在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
//
// 示例:
//
// 输入:
//
//1 0 1 0 0
//1 0 1 1 1
//1 1 1 1 1
//1 0 0 1 0
//
//输出: 4
// Related Topics 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	//构造比原数组长宽大1的临时数组
	temp := make([][]int, len(matrix)+1)
	for i := range temp {
		temp[i] = make([]int, len(matrix[0])+1)
	}
	//最大正方形的边长
	res := 0
	for i := 1; i < len(temp); i++ {
		for j := 1; j < len(temp[0]); j++ {
			//元数组为byte类型
			if matrix[i-1][j-1]&1 == 1 {
				//当前节点边长，是另外三个节点最小值+1
				temp[i][j] = min(min(temp[i-1][j], temp[i][j-1]), temp[i-1][j-1]) + 1
				if temp[i][j] > res {
					res = temp[i][j]
				}
			}
		}
	}
	return res * res
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
