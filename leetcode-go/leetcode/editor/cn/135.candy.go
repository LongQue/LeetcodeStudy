package main

/*
[135]分发糖果
//老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
//
// 你需要按照以下要求，帮助老师给这些孩子分发糖果：
//
//
// 每个孩子至少分配到 1 个糖果。
// 相邻的孩子中，评分高的孩子必须获得更多的糖果。
//
//
// 那么这样下来，老师至少需要准备多少颗糖果呢？
//
// 示例 1:
//
// 输入: [1,0,2]
//输出: 5
//解释: 你可以分别给这三个孩子分发 2、1、2 颗糖果。
//
//
// 示例 2:
//
// 输入: [1,2,2]
//输出: 4
//解释: 你可以分别给这三个孩子分发 1、2、1 颗糖果。
//     第三个孩子只得到 1 颗糖果，这已满足上述两个条件。
// Related Topics 贪心算法
*/

//leetcode submit region begin(Prohibit modification and deletion)
func candy(ratings []int) int {
	//题目可分为，左大于右&&右大于左时至少需要多少，在满足这两个条件的情况下需最大值
	//构造左右两分片，初始值为1
	left := make([]int, len(ratings))
	for i := range left {
		left[i] = 1
	}
	right := make([]int, len(left))
	copy(right, left)
	//从左侧开始，计算满足右大于左的+1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	count := left[len(left)-1]
	//从右侧开始，计算左大于右的+1
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
		//取满足这两个条件所需糖果至少要多少
		count += max(right[i], left[i])
	}
	return count
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
