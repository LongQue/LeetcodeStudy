package main

import "sort"

/*
[406]根据身高重建队列
//假设有打乱顺序的一群人站成一个队列。 每个人由一个整数对(h, k)表示，其中h是这个人的身高，k是排在这个人前面且身高大于或等于h的人数。 编写一个算法来
//重建这个队列。
//
// 注意：
//总人数少于1100人。
//
// 示例
//
//
//输入:
//[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
//
//输出:
//[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
//
// Related Topics 贪心算法
*/

//leetcode submit region begin(Prohibit modification and deletion)
func reconstructQueue(people [][]int) [][]int {
	//[i][j] 身高i相同按j排序，j小在前，身高不同按i排序，i大在前
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	var res [][]int
	for _, v := range people {
		//如果j==res长度，则意味着直接放入即可
		if v[1] == len(res) {
			res = append(res, v)
		} else {
			//不等于长度的，则需要插入到j对于的下标
			//因为[:]会改变原数组，所以append可以new 不受原数组影响，获取右侧
			right := append([][]int{}, res[v[1]:]...)
			//v加入左侧
			left := append(res[:v[1]], v)
			//拼接
			res = append(left, right...)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
