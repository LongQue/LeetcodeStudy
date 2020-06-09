package main

/*
[990]等式方程的可满足性
//给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 equations[i] 的长度为 4，并采用两种不同的形式之一："a==b" 或 "a!
//=b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。
//
// 只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回 true，否则返回 false。
//
//
//
//
//
//
// 示例 1：
//
// 输入：["a==b","b!=a"]
//输出：false
//解释：如果我们指定，a = 1 且 b = 1，那么可以满足第一个方程，但无法满足第二个方程。没有办法分配变量同时满足这两个方程。
//
//
// 示例 2：
//
// 输入：["b==a","a==b"]
//输出：true
//解释：我们可以指定 a = 1 且 b = 1 以满足满足这两个方程。
//
//
// 示例 3：
//
// 输入：["a==b","b==c","a==c"]
//输出：true
//
//
// 示例 4：
//
// 输入：["a==b","b!=c","c==a"]
//输出：false
//
//
// 示例 5：
//
// 输入：["c==c","b==d","x!=z"]
//输出：true
//
//
//
//
// 提示：
//
//
// 1 <= equations.length <= 500
// equations[i].length == 4
// equations[i][0] 和 equations[i][3] 是小写字母
// equations[i][1] 要么是 '='，要么是 '!'
// equations[i][2] 是 '='
//
// Related Topics 并查集 图
*/

//leetcode submit region begin(Prohibit modification and deletion)
func equationsPossible(equations []string) bool {
	//等于的传递性，说明a==b b==c 即 a==b，把相等的串起来，他们拥有相同的根节点

	//将数据看成一个个节点，记录26字母，通过下标指向下一个相等的节点，这样就可以串起来
	parent := make([]int, 26)
	//赋予初始值，需要与findRoot函数配合，
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}
	//把相等的字母连起来
	for _, v := range equations {
		if v[1] == '=' {
			//加入新的根
			updateRoot(parent, int(v[0]-'a'), int(v[3]-'a'))
		}
	}
	for _, v := range equations {
		if v[1] == '!' {
			//判断两者的跟相同
			if findRoot(parent, int(v[0]-'a')) == findRoot(parent, int(v[3]-'a')) {
				return false
			}
		}
	}
	return true
}
func updateRoot(parent []int, pre, last int) {
	//举例子第一步 a==b  则开始的时候 数组中都是自身的下标，所欲 parent[0]=1
	//第二步 b==c  则 parent[1]=2
	//第三步 a==c   a 根据parent[0]=1 查找parent[1]=2 再到parent[2]
	parent[findRoot(parent, pre)] = findRoot(parent, last)
}
func findRoot(parent []int, index int) int {
	//查找根节点
	for parent[index] != index {
		parent[index] = parent[parent[index]]
		index = parent[index]
	}
	return index
}

//leetcode submit region end(Prohibit modification and deletion)
