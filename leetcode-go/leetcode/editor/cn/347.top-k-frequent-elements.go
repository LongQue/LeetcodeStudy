package main

/*
[347]前 K 个高频元素
//给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
//
//
//
// 示例 1:
//
// 输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
//
//
// 示例 2:
//
// 输入: nums = [1], k = 1
//输出: [1]
//
//
//
// 提示：
//
//
// 你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
// 你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
// 你可以按任意顺序返回答案。
//
// Related Topics 堆 哈希表
*/

//leetcode submit region begin(Prohibit modification and deletion)
func topKFrequent(nums []int, k int) []int {
	//list 作为桶排序， 出现次数从 0到len(nums)次，所以len(nums)+1
	//但由于下面的遍历方式所以没出现过的不会放入0次
	// 第一层为频率，第二层为出现的数字
	list := make([][]int, len(nums)+1)
	//用于统计某个数字出现的次数
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	//k出现的数字， v出现的次数
	for k, v := range m {
		//用次数v做桶排序，数字k作为值
		list[v] = append(list[v], k)
	}
	var res []int
	//从后面开始遍历，将次数高的一次加入到结果中，直到结果的长度不小于k
	//最多len(nums)次，最小1次， 0次没意义
	for i := len(nums); i >= 1 && len(res) < k; i-- {
		res = append(res, list[i]...)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
