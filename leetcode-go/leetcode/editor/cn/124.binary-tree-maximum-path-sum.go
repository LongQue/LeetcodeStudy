package main

/*
[124]二叉树中的最大路径和
//给定一个非空二叉树，返回其最大路径和。
//
// 本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
//
// 示例 1:
//
// 输入: [1,2,3]
//
//       1
//      / \
//     2   3
//
//输出: 6
//
//
// 示例 2:
//
// 输入: [-10,9,20,null,null,15,7]
//
//   -10
//   / \
//  9  20
//    /  \
//   15   7
//
//输出: 42
// Related Topics 树 深度优先搜索
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
var maxValue int

func maxPathSum(root *TreeNode) int {
	// 最小值，无字符右移一位是最大值，再取反是最小值
	maxValue = ^int(^uint(0) >> 1)
	maxPathSumHelper(root)
	return maxValue
}
func maxPathSumHelper(root *TreeNode) int {
	//nil节点返回0
	if root == nil {
		return 0
	}
	//左子树最大和
	left := maxPathSumHelper(root.Left)
	//右子树最大和
	right := maxPathSumHelper(root.Right)
	//当前树最大和  左+中+右
	temp := root.Val + max(0, left) + max(0, right)
	//只算一遍路径的最大和  中+左/右
	ret := root.Val + max(0, max(left, right))
	//最大值
	maxValue = max(maxValue, max(temp, ret))
	return ret
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
