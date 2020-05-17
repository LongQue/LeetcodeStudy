package main

/*
[560]和为K的子数组
//给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
//
// 示例 1 :
//
//
//输入:nums = [1,1,1], k = 2
//输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
//
//
// 说明 :
//
//
// 数组的长度为 [1, 20,000]。
// 数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
//
// Related Topics 数组 哈希表
*/

//leetcode submit region begin(Prohibit modification and deletion)
func subarraySum(nums []int, k int) int {
	//前缀和，把和最为key，value作为次数来记录某个前缀和出现过多少次，然后通过新的前缀减k，取value
	m := make(map[int]int)
	//0距离
	m[0] = 1
	preSum := 0
	count := 0
	for i := range nums {
		preSum += nums[i]
		if temp, ok := m[preSum-k]; ok {
			count += temp
		}
		m[preSum] = m[preSum] + 1
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
