package main

/*
[210]课程表 II
//现在你总共有 n 门课需要选，记为 0 到 n-1。
//
// 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]
//
// 给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。
//
// 可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。
//
// 示例 1:
//
// 输入: 2, [[1,0]]
//输出: [0,1]
//解释: 总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
//
// 示例 2:
//
// 输入: 4, [[1,0],[2,0],[3,1],[3,2]]
//输出: [0,1,2,3] or [0,2,1,3]
//解释: 总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
//     因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
//
//
// 说明:
//
//
// 输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
// 你可以假定输入的先决条件中没有重复的边。
//
//
// 提示:
//
//
// 这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
// 通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
//
// 拓扑排序也可以通过 BFS 完成。
//
//
// Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序
*/

//leetcode submit region begin(Prohibit modification and deletion)
func findOrder(numCourses int, prerequisites [][]int) []int {
	//记录每个元素的入度
	degrees := make([]int, numCourses)
	//记录指向，如1->0 1->2 则为 1->[0,2]
	temps := make([][]int, numCourses)
	//[[1,0]]先学0再学1
	for _, v := range prerequisites {
		degrees[v[0]]++
		temps[v[1]] = append(temps[v[1]], v[0])
	}
	var queue []int
	//放入入度为0的点
	for i, v := range degrees {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	//需要额外变量记录长度
	length := len(queue)
	var res []int
	for length > 0 {
		//获取结果目前的顺序
		res = append(res, queue...)
		//遍历queue
		for i := 0; i < length; i++ {
			//找出queue中元素的指向
			for j := 0; j < len(temps[queue[i]]); j++ {
				//找出queue元素指向的列表的元素，使其入度减1
				degrees[temps[queue[i]][j]]--
				//入度为0，入列
				if degrees[temps[queue[i]][j]] == 0 {
					queue = append(queue, temps[queue[i]][j])
				}
			}
		}
		queue = queue[length:]
		length = len(queue)
	}
	if len(res) < numCourses {
		return nil
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
