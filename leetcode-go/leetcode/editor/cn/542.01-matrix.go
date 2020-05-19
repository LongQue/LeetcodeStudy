package main

/*
[542]01 矩阵
//给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。
//
// 两个相邻元素间的距离为 1 。
//
// 示例 1:
//输入:
//
//
//0 0 0
//0 1 0
//0 0 0
//
//
// 输出:
//
//
//0 0 0
//0 1 0
//0 0 0
//
//
// 示例 2:
//输入:
//
//
//0 0 0
//0 1 0
//1 1 1
//
//
// 输出:
//
//
//0 0 0
//0 1 0
//1 2 1
//
//
// 注意:
//
//
// 给定矩阵的元素个数不超过 10000。
// 给定矩阵中至少有一个元素是 0。
// 矩阵中的元素只在四个方向上相邻: 上、下、左、右。
//
// Related Topics 深度优先搜索 广度优先搜索
*/

//leetcode submit region begin(Prohibit modification and deletion)
func updateMatrix(matrix [][]int) [][]int {
	var queue [][]int
	m := len(matrix)
	n := len(matrix[0])
	//所有0加入，将所有1改为-1，标记为 未访问节点
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				matrix[i][j] = -1
			}
		}
	}
	d := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	length := len(queue)
	for length > 0 {
		for i := 0; i < length; i++ {
			x := queue[i][0]
			y := queue[i][1]
			for j := 0; j < 4; j++ {
				newX := x + d[j][0]
				newY := y + d[j][1]
				if newX >= 0 && newX < m && newY >= 0 && newY < n && matrix[newX][newY] == -1 {
					//和最近0的距离 相当于来源节点+1
					matrix[newX][newY] = matrix[x][y] + 1
					queue = append(queue, []int{newX, newY})
				}
			}
		}
		queue = queue[length:]
		length = len(queue)
	}
	return matrix
}

//leetcode submit region end(Prohibit modification and deletion)
