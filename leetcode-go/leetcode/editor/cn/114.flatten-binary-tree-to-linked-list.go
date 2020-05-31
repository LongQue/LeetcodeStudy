package main

/*
[114]二叉树展开为链表
//给定一个二叉树，原地将它展开为一个单链表。
//
//
//
// 例如，给定二叉树
//
//     1
//   / \
//  2   5
// / \   \
//3   4   6
//
// 将其展开为：
//
// 1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6
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
func flatten(root *TreeNode) {
	for root != nil {
		if root.Left != nil {
			pre := root.Left
			//右子树拼接到左子树最右的节点
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = root.Right
			//root 左子树覆盖右子树
			root.Right = root.Left
			//root 左子树nil
			root.Left = nil
		}
		root = root.Right
	}
}

//leetcode submit region end(Prohibit modification and deletion)
