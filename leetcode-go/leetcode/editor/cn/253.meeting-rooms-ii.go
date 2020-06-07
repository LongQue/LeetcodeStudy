package main

import "sort"

/*
[253]会议室 II
//给定一个会议时间安排的数组，每个会议时间都会包括开始和结束的时间 [[s1,e1],[s2,e2],...] (si < ei)，为避免会议冲突，同时要考虑
//充分利用会议室资源，请你计算至少需要多少间会议室，才能满足这些会议安排。
//
// 示例 1:
//
// 输入: [[0, 30],[5, 10],[15, 20]]
//输出: 2
//
// 示例 2:
//
// 输入: [[7,10],[2,4]]
//输出: 1
// Related Topics 堆 贪心算法 排序
*/

//leetcode submit region begin(Prohibit modification and deletion)
func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	var start []int
	var end []int
	//开始结束时间分组&排序
	for _, v := range intervals {
		start = append(start, v[0])
		end = append(end, v[1])
	}
	sort.Ints(start)
	sort.Ints(end)
	//如果开始时间小于结束时间，说明要多加一间会议室
	//从0开始遍历开始结束
	startIndex, endIndex := 0, 0
	length := len(intervals)
	useRoom := 0
	for startIndex < length {
		if start[startIndex] < end[endIndex] {
			useRoom++
		} else {
			//不小于时，结束时间下标+1
			endIndex++
		}
		startIndex++
	}
	return useRoom
}

//leetcode submit region end(Prohibit modification and deletion)
