package main

import (
	"math"
	"sort"
)

/*
[16]最接近的三数之和
//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和
//。假定每组输入只存在唯一答案。
//
//
//
// 示例：
//
// 输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 10^3
// -10^3 <= nums[i] <= 10^3
// -10^4 <= target <= 10^4
//
// Related Topics 数组 双指针
*/

//leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := range nums {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-ans)) {
				ans = sum
			}
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				return ans
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
