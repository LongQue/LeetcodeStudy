package main

/*
[437]路径总和 III
//给定一个二叉树，它的每个结点都存放着一个整数值。
//
// 找出路径和等于给定数值的路径总数。
//
// 路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//
// 二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。
//
// 示例：
//
// root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
//
//      10
//     /  \
//    5   -3
//   / \    \
//  3   2   11
// / \   \
//3  -2   1
//
//返回 3。和等于 8 的路径有:
//
//1.  5 -> 3
//2.  5 -> 2 -> 1
//3.  -3 -> 11
//
// Related Topics 树
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	//只需要求出个数，用前缀和，
	//先序遍历树，将每条路径上的值相加放入map，map[当前值-目标值]的结果说明可以有多少条路径
	preSumCount := make(map[int]int)
	preSumCount[0] = 1
	return pathSumHelper(root, preSumCount, sum, 0)
}
func pathSumHelper(root *TreeNode, preSumCount map[int]int, sum, cur int) int {
	if root == nil {
		return 0
	}
	res := 0
	cur += root.Val

	res += preSumCount[cur-sum]
	preSumCount[cur]++

	res += pathSumHelper(root.Left, preSumCount, sum, cur)
	res += pathSumHelper(root.Right, preSumCount, sum, cur)
	preSumCount[cur]--
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
