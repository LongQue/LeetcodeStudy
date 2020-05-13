package main

import "sort"

/*
[56]合并区间
//给出一个区间的集合，请合并所有重叠的区间。
//
// 示例 1:
//
// 输入: [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
//
// 示例 2:
//
// 输入: [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
// Related Topics 排序 数组
*/

//leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//记录合并结果
	result := make([][]int, 0, len(intervals))
	//合并分片的下标
	index := -1
	for _, v := range intervals {
		// -1代表没有合并，可直接放入。后面判断代表没有重合，可直接放入
		if index == -1 || v[0] > result[index][1] {
			result = append(result, v)
			index++
		} else {
			//与前一个存在重合，比较两者[1],去最大值
			result[index][1] = max(result[index][1], v[1])
		}
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
