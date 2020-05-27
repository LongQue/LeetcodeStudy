package main

/*
[85]最大矩形
//给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
//
// 示例:
//
// 输入:
//[
//  ["1","0","1","0","0"],
//  ["1","0","1","1","1"],
//  ["1","1","1","1","1"],
//  ["1","0","0","1","0"]
//]
//输出: 6
// Related Topics 栈 数组 哈希表 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func maximalRectangle(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 {
		return 0
	}
	maxValue := 0
	height := make([]int, len(matrix[0]))
	//一行一行的叠加，然后找出最大矩形，如遇到0则该列矩形改为0
	for _, v := range matrix {
		for j, b := range v {
			if b == '0' {
				height[j] = 0
			} else {
				height[j] += 1
			}
		}
		maxValue = max(maxValue, largestRectangleArea(height))
	}
	return maxValue
}

//84 最大矩形的解
func largestRectangleArea(heights []int) int {
	ret := 0
	//柱形图专用，哨兵技巧，在左右两边放个0
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)
	//作为栈存放单调递增的数据的下标
	var temp []int
	for i := 0; i < len(heights); i++ {
		//如果栈不为空且柱形图后一个小于前一个，那么这时候就可以把不小于heights[i]的出栈计算这些的面积。
		for len(temp) != 0 && heights[temp[len(temp)-1]] > heights[i] {
			//栈顶的下标
			cur := temp[len(temp)-1]
			//出栈
			temp = temp[:len(temp)-1]
			//此时栈顶的下边，需要+1
			left := temp[len(temp)-1] + 1
			right := i - 1
			//计算面积
			ret = max(ret, (right-left+1)*heights[cur])
		}
		//每遍历一个入栈
		temp = append(temp, i)
	}
	return ret
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
