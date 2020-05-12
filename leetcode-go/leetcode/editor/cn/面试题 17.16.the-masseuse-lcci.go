package main

/*
[面试题 17.16]按摩师
//一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。给定一个预约请求序列，替按摩
//师找到最优的预约集合（总预约时间最长），返回总的分钟数。
//
// 注意：本题相对原题稍作改动
//
//
//
// 示例 1：
//
// 输入： [1,2,3,1]
//输出： 4
//解释： 选择 1 号预约和 3 号预约，总时长 = 1 + 3 = 4。
//
//
// 示例 2：
//
// 输入： [2,7,9,3,1]
//输出： 12
//解释： 选择 1 号预约、 3 号预约和 5 号预约，总时长 = 2 + 9 + 1 = 12。
//
//
// 示例 3：
//
// 输入： [2,1,4,5,3,1,1,3]
//输出： 12
//解释： 选择 1 号预约、 3 号预约、 5 号预约和 8 号预约，总时长 = 2 + 4 + 3 + 3 = 12。
//
// Related Topics 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func massage(nums []int) int {
	if nums == nil {
		return 0
	}
	//记录每多一个元素时的最优，多出的2存放+的操作，i是遍历nums f(i+2)=max(f(i)+nums[i],f(i+1))
	arrays := make([]int, len(nums)+2)
	//动态规划
	for i := range nums {
		arrays[i+2] = max(arrays[i]+nums[i], arrays[i+1])
	}
	return arrays[len(arrays)-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
