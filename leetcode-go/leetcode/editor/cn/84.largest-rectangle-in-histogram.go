package main

/*
[84]柱状图中最大的矩形
//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//
//
//
//
// 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
//
//
//
//
//
// 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
//
//
//
// 示例:
//
// 输入: [2,1,5,6,2,3]
//输出: 10
// Related Topics 栈 数组
*/

//leetcode submit region begin(Prohibit modification and deletion)
//https://leetcode-cn.com/problems/largest-rectangle-in-histogram/solution/84-by-ikaruga/
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
