package main

/*
[200]岛屿数量
//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
//
//
// 示例 1:
//
// 输入:
//11110
//11010
//11000
//00000
//输出: 1
//
//
// 示例 2:
//
// 输入:
//11000
//11000
//00100
//00011
//输出: 3
//解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
//
// Related Topics 深度优先搜索 广度优先搜索 并查集
*/

//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	directions := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	n := len(grid)
	m := len(grid[0])
	temp := make([][]int, n)
	var queue [][]int
	count := 0
	for i := range temp {
		temp[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' && temp[i][j] != 1 {
				temp[i][j] = 1
				count++
				queue = append(queue, []int{i, j})
				for len(queue) != 0 {
					y := queue[0][0]
					x := queue[0][1]
					queue = queue[1:]
					for k := 0; k < 4; k++ {
						newY := y + directions[k][0]
						newX := x + directions[k][1]
						if inArea(newX, newY, m, n) && grid[newY][newX] == '1' && temp[newY][newX] != 1 {
							temp[newY][newX] = 1
							queue = append(queue, []int{newY, newX})
						}
					}
				}
			}
		}
	}
	return count
}
func inArea(x, y, m, n int) bool {
	return x < m && x >= 0 && y < n && y >= 0

}

//leetcode submit region end(Prohibit modification and deletion)
