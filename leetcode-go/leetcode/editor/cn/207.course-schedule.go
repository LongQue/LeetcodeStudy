package main

/*
[207]课程表
//你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
//
// 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
//
// 给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
//
//
//
// 示例 1:
//
// 输入: 2, [[1,0]]
//输出: true
//解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
//
// 示例 2:
//
// 输入: 2, [[1,0],[0,1]]
//输出: false
//解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
//
//
//
// 提示：
//
//
// 输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
// 你可以假定输入的先决条件中没有重复的边。
// 1 <= numCourses <= 10^5
//
// Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序
*/

//leetcode submit region begin(Prohibit modification and deletion)
func canFinish(numCourses int, prerequisites [][]int) bool {
	//统计放入queue中的次数
	count := 0
	//记录每门课程的入度
	rudu := make([]int, numCourses)
	//记录每门课程的出度指向
	chudu := make([][]int, numCourses)
	for _, v := range prerequisites {
		rudu[v[0]]++
		if len(chudu[v[1]]) == 0 {
			chudu[v[1]] = []int{v[0]}
		} else {
			chudu[v[1]] = append(chudu[v[1]], v[0])
		}
	}
	var queue []int
	//入度为0表名不需要前置课程，直接入列
	for i, v := range rudu {
		if v == 0 {
			queue = append(queue, i)
			//入列个数加1
			count++
		}
	}
	for len(queue) != 0 {
		//取出第一个
		temp := queue[0]
		queue = queue[1:]
		//查找出度指向列表
		arrays := chudu[temp]
		//遍历列表，列表中的元素的入度减一，如果为0则入列，入列数+1
		for _, v := range arrays {
			rudu[v]--
			if rudu[v] == 0 {
				queue = append(queue, v)
				count++
			}
		}
	}
	//判断入列数等于课程数
	return count == numCourses
}

//leetcode submit region end(Prohibit modification and deletion)
