package main

/*
[503]下一个更大元素 II
//给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第
//一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。
//
// 示例 1:
//
//
//输入: [1,2,1]
//输出: [2,-1,2]
//解释: 第一个 1 的下一个更大的数是 2；
//数字 2 找不到下一个更大的数；
//第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
//
//
// 注意: 输入数组的长度不会超过 10000。
// Related Topics 栈
// 👍 181 👎 0
*/

//leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElements(nums []int) []int {
	result := make([]int, len(nums)) //执行结果
	var arrays []int                 //模拟单调栈

	//从后往前遍历，记录单调栈(前一个元素可看到的更大元素),增加一倍
	n := len(nums)
	for i := n*2 - 1; i >= 0; i-- {
		//不为0且前一个元素大于栈顶时，出栈，<=去重
		for len(arrays) != 0 && arrays[len(arrays)-1] <= nums[i%n] {
			arrays = arrays[:len(arrays)-1]
		}
		if len(arrays) == 0 {
			result[i%n] = -1
		} else {
			result[i%n] = arrays[len(arrays)-1]
		}
		arrays = append(arrays, nums[i%n])
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
