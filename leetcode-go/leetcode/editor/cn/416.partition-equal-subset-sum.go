package main

/*
[416]分割等和子集
//给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
// 注意:
//
//
// 每个数组中的元素不会超过 100
// 数组的大小不会超过 200
//
//
// 示例 1:
//
// 输入: [1, 5, 11, 5]
//
//输出: true
//
//解释: 数组可以分割成 [1, 5, 5] 和 [11].
//
//
//
//
// 示例 2:
//
// 输入: [1, 2, 3, 5]
//
//输出: false
//
//解释: 数组不能分割成两个元素和相等的子集.
//
//
//
// Related Topics 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	//统计总数，如果为单则返回false
	if sum&1 == 1 {
		return false
	}
	//除2
	sum = sum >> 1
	n := len(nums)
	//temp[n+1][sum+1]   temp，使用前n+1个时可以满足sum+1的容量，n+1代表从0-n个，sum+1代表从0-sum容量。+1是为了应付 0的特殊值
	temp := make([][]bool, n+1)
	for i := range temp {
		temp[i] = make([]bool, sum+1)
	}
	// 使用任意个都可以满足 0容量
	for i := range temp {
		temp[i][0] = true
	}

	//外层从1遍历n，代表使用nums数组前n个
	for i := 1; i <= n; i++ {
		//内层从1遍历sum，代表需要满足的容量
		for j := 1; j <= sum; j++ {
			//小于0则说明第i个不足以满足剩余容量
			if j-nums[i-1] < 0 {
				//第i个不装入背包，所以等于上层，
				temp[i][j] = temp[i-1][j]
			} else {
				//大于等于0，说明不放入或者放入，或优先用true
				//j-nums[i-1] 查看前n-1个是否可满足 j-nums[i-1]容量，如果可以则说明放入nums[i-1]时，可满足j
				temp[i][j] = temp[i-1][j] || temp[i-1][j-nums[i-1]]
			}
		}
	}

	return temp[n][sum]
}

//leetcode submit region end(Prohibit modification and deletion)
