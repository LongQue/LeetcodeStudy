package main

import "sort"

/*
[15]三数之和
//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复
//的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例：
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
//
// Related Topics 数组 双指针0
*/

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	//排序+双指针
	sort.Ints(nums)
	var list [][]int
	for i := range nums {
		//固定当前值nums[i]，双指针 第一个从i+1开始 第二个从后面正数的开始
		//如果当前值大于0，则不可能还有
		if nums[i] > 0 {
			return list
		}
		//相同的数值直接跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		L := i + 1
		R := len(nums) - 1
		for L < R {
			temp := nums[i] + nums[L] + nums[R]
			if temp == 0 {
				list = append(list, []int{nums[i], nums[L], nums[R]})
				//相同的数值直接跳过
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				//相同的数值直接跳过
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if temp < 0 {
				L++
			} else {
				R--
			}
		}
	}
	return list
}

//leetcode submit region end(Prohibit modification and deletion)
