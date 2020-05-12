package main

/*
[42]接雨水
//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Mar
//cos 贡献此图。
//
// 示例:
//
// 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
//输出: 6
// Related Topics 栈 数组 双指针
*/

//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	//总和
	sum := 0
	//左最大值
	lMax := 0
	//右最大值
	rMax := 0
	//当前左索引
	l := 1
	//当前右索引
	r := len(height) - 2
	for l <= r {

		if height[l-1] < height[r+1] {
			lMax = max(lMax, height[l-1])
			if lMax > height[l] {
				sum += lMax - height[l]
			}
			l++
		} else {
			rMax = max(rMax, height[r+1])
			if rMax > height[r] {
				sum += rMax - height[r]
			}
			r--
		}
	}
	return sum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
