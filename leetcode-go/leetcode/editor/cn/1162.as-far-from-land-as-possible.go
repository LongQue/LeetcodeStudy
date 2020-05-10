package main

/*
[1162]地图分析
//你现在手里有一份大小为 N x N 的「地图」（网格） grid，上面的每个「区域」（单元格）都用 0 和 1 标记好了。其中 0 代表海洋，1 代表陆地，
//请你找出一个海洋区域，这个海洋区域到离它最近的陆地区域的距离是最大的。
//
// 我们这里说的距离是「曼哈顿距离」（ Manhattan Distance）：(x0, y0) 和 (x1, y1) 这两个区域之间的距离是 |x0 - x
//1| + |y0 - y1| 。
//
// 如果我们的地图上只有陆地或者海洋，请返回 -1。
//
//
//
// 示例 1：
//
//
//
// 输入：[[1,0,1],[0,0,0],[1,0,1]]
//输出：2
//解释：
//海洋区域 (1, 1) 和所有陆地区域之间的距离都达到最大，最大距离为 2。
//
//
// 示例 2：
//
//
//
// 输入：[[1,0,0],[0,0,0],[0,0,0]]
//输出：4
//解释：
//海洋区域 (2, 2) 和所有陆地区域之间的距离都达到最大，最大距离为 4。
//
//
//
//
// 提示：
//
//
// 1 <= grid.length == grid[0].length <= 100
// grid[i][j] 不是 0 就是 1
//
// Related Topics 广度优先搜索 图
*/

//leetcode submit region begin(Prohibit modification and deletion)
func maxDistance(grid [][]int) int {
	var queue [][]int
	M := len(grid)
	N := len(grid[0])
	//BFS 先记录陆地的点，然后查看四周
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	//必须是-1
	res := -1
	//排除全是陆地
	if len(queue) == M*N {
		return res
	}
	//有陆地查看陆地的四周
	for len(queue) > 0 {
		res++
		temp := len(queue)
		for i := 0; i < temp; i++ {
			a, b := queue[i][0], queue[i][1]
			if a-1 >= 0 && grid[a-1][b] == 0 {
				//标记以遍历过
				grid[a-1][b] = 1
				queue = append(queue, []int{a - 1, b})
			}
			if b-1 >= 0 && grid[a][b-1] == 0 {
				grid[a][b-1] = 1
				queue = append(queue, []int{a, b - 1})
			}
			if a+1 < M && grid[a+1][b] == 0 {
				grid[a+1][b] = 1
				queue = append(queue, []int{a + 1, b})
			}
			if b+1 < N && grid[a][b+1] == 0 {
				grid[a][b+1] = 1
				queue = append(queue, []int{a, b + 1})
			}
		}
		queue = queue[temp:]
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
