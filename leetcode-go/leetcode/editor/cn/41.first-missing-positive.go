package main

/*
[41]缺失的第一个正数

//给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。
//
//
//
// 示例 1:
//
// 输入: [1,2,0]
//输出: 3
//
//
// 示例 2:
//
// 输入: [3,4,-1,1]
//输出: 2
//
//
// 示例 3:
//
// 输入: [7,8,9,11,12]
//输出: 1
//
//
//
//
// 提示：
//
// 你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。
// Related Topics 数组
*/

//leetcode submit region begin(Prohibit modification and deletion)
func firstMissingPositive(nums []int) int {
	//答案最小正整数，必然在1~length+1之间，这个包括了length+1个数字
	length := len(nums)
	//原地换位置
	//1. >0   <=length的说明存在 下标+1=nums[i],将其与下标互换，之后for nums[i]寻找是否还有符合的，如果没有则i++
	for i := range nums {
		for nums[i] > 0 && nums[i] <= length && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	//最后所有符合的值必然是下标+1

	for i, v := range nums {
		//如果不符合则后面返回下标+1
		if v != i+1 {
			return i + 1
		}
	}
	//全部符合 则长度+1
	return length + 1
}

//leetcode submit region end(Prohibit modification and deletion)
