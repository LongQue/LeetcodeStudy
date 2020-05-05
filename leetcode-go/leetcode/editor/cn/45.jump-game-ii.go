package main

/*
[45]跳跃游戏 II
//给定一个非负整数数组，你最初位于数组的第一个位置。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//
// 示例:
//
// 输入: [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//
//
// 说明:
//
// 假设你总是可以到达数组的最后一个位置。
// Related Topics 贪心算法 数组
*/

//leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) int {
	step := 0
	if len(nums) == 0 {
		return step
	}
	position := 0
	for position < len(nums)-1 {
		//len(num)-1 由于数组的每个元素都是正整数，所以达到 nums[len(nums)-2]=1时必定能达到终点
		if nums[position]+position >= len(nums)-1 {
			step++
			break
		} else {
			max, maxi := 0, 0
			for i := 1; i <= nums[position]; i++ {
				//找出 position ~ position+i之间的最大值
				if max < i+nums[position+i] {
					max = i + nums[position+i]
					maxi = i
				}
			}
			position += maxi
			step++
		}
	}
	return step
}

//leetcode submit region end(Prohibit modification and deletion)
