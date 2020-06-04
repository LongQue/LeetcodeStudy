package main

/*
[239]滑动窗口最大值
//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
//
// 返回滑动窗口中的最大值。
//
//
//
// 进阶：
//
// 你能在线性时间复杂度内解决此题吗？
//
//
//
// 示例:
//
// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7]
//解释:
//
//  滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10^5
// -10^4 <= nums[i] <= 10^4
// 1 <= k <= nums.length
//
// Related Topics 堆 Sliding Window
*/

//leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	var Q = make([]int, 0, len(nums))
	res := make([]int, len(nums)-k+1)
	for i := 0; i < len(nums); i++ {

		for len(Q) != 0 && nums[i] >= nums[Q[len(Q)-1]] {
			Q = Q[:len(Q)-1]
		}
		// 当前元素入栈
		Q = append(Q, i)

		// 窗口离开first元素时，栈中第一个元素出栈
		if Q[0] == i-k {
			Q = Q[1:]
		}
		// 窗口中满了k个数
		if i+1-k >= 0 {
			res[i+1-k] = nums[Q[0]]
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
