package main

/*
[199]二叉树的右视图
//给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//
// 示例:
//
// 输入: [1,2,3,null,5,null,4]
//输出: [1, 3, 4]
//解释:
//
//   1            <---
// /   \
//2     3         <---
// \     \
//  5     4       <---
//
// Related Topics 树 深度优先搜索 广度优先搜索
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
//广度遍历，每次找队列中最右侧
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var temp []*TreeNode
	temp = append(temp, root)
	var result []int
	for len(temp) > 0 {
		n := len(temp)
		result = append(result, temp[n-1].Val)
		for i := 0; i < n; i++ {
			if temp[i].Left != nil {
				temp = append(temp, temp[i].Left)
			}
			if temp[i].Right != nil {
				temp = append(temp, temp[i].Right)
			}
		}
		temp = temp[n:]
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
