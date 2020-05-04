package main

/*
[750]角矩形的数量
//给定一个只包含 0 和 1 的网格，找出其中角矩形的数量。
//
// 一个「角矩形」是由四个不同的在网格上的 1 形成的轴对称的矩形。注意只有角的位置才需要为 1。并且，4 个 1 需要是不同的。
//
//
//
// 示例 1：
//
// 输入：grid =
//[[1, 0, 0, 1, 0],
// [0, 0, 1, 0, 1],
// [0, 0, 0, 1, 0],
// [1, 0, 1, 0, 1]]
//输出：1
//解释：只有一个角矩形，角的位置为 grid[1][2], grid[1][4], grid[3][2], grid[3][4]。
//
//
// 示例 2：
//
// 输入：grid =
//[[1, 1, 1],
// [1, 1, 1],
// [1, 1, 1]]
//输出：9
//解释：这里有 4 个 2x2 的矩形，4 个 2x3 和 3x2 的矩形和 1 个 3x3 的矩形。
//
//
// 示例 3：
//
// 输入：grid =
//[[1, 1, 1, 1]]
//输出：0
//解释：矩形必须有 4 个不同的角。
//
//
//
//
// 提示：
//
//
// 网格 grid 中行和列的数目范围为 [1, 200]。
// 每个网格 grid[i][j] 中的值不是 0 就是 1 。
// 网格中 1 的个数不会超过 6000。
//
//
//
// Related Topics 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func countCornerRectangles(grid [][]int) int {
	temp := make([][]int, 200)
	for i := range temp {
		temp[i] = make([]int, 200)
	}
	count := 0
	//遍历每个1
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				for k := j + 1; k < len(grid[i]); k++ {
					//查找同一行的下一个是否为1
					if grid[i][k] == 1 {
						//temp初始为0
						count += temp[j][k]
						//j,k代表这两列有多少对1
						temp[j][k]++
					}
				}
			}
		}
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
