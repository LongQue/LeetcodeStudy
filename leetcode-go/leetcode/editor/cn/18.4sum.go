package main

import "sort"

/*
[18]四数之和
//给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c +
// d 的值与 target 相等？找出所有满足条件且不重复的四元组。
//
// 注意：
//
// 答案中不可以包含重复的四元组。
//
// 示例：
//
// 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
//
//满足要求的四元组集合为：
//[
//  [-1,  0, 0, 1],
//  [-2, -1, 1, 2],
//  [-2,  0, 0, 2]
//]
//
// Related Topics 数组 哈希表 双指针
// 👍 509 👎 0
*/

//leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for a := 0; a < len(nums)-3; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b < len(nums)-2; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			c := b + 1
			d := len(nums) - 1
			for c < d {
				if nums[a]+nums[b]+nums[c]+nums[d] < target {
					c++
				} else if nums[a]+nums[b]+nums[c]+nums[d] > target {
					d--
				} else {
					result = append(result, append([]int{nums[a]}, nums[b], nums[c], nums[d]))
					for c < d && nums[c] == nums[c+1] {
						c++
					}
					for c < d && nums[d] == nums[d-1] {
						d--
					}
					c++
					d--
				}
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
