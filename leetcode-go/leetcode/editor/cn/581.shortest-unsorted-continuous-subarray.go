package main

/*
[581]最短无序连续子数组
//给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
//
// 你找到的子数组应是最短的，请输出它的长度。
//
// 示例 1:
//
//
//输入: [2, 6, 4, 8, 10, 9, 15]
//输出: 5
//解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
//
//
// 说明 :
//
//
// 输入的数组长度范围在 [1, 10,000]。
// 输入的数组可能包含重复元素 ，所以升序的意思是<=。
//
// Related Topics 数组
*/

//leetcode submit region begin(Prohibit modification and deletion)
func findUnsortedSubarray(nums []int) int {
	length := len(nums)
	max := nums[0]
	min := nums[length-1]
	l, r := 0, -1
	for i := 0; i < length; i++ {
		//从左往右遍历找到右边界，小于当前左侧最大值则说明右边界是该位置，不断修改
		if max > nums[i] {
			r = i
		} else {
			max = nums[i]
		}
		//从右往左遍历找到左边界，大于当前右侧最小值则说明左边界是该位置，不断修改
		if min < nums[length-i-1] {
			l = length - i - 1
		} else {
			min = nums[length-i-1]
		}
	}
	return r - l + 1
}

//leetcode submit region end(Prohibit modification and deletion)
