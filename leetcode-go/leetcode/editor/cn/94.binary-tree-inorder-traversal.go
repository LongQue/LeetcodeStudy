package main

/*
[94]二叉树的中序遍历
//给定一个二叉树，返回它的中序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [1,3,2]
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 哈希表
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
func inorderTraversal(root *TreeNode) []int {
	var ret []int
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		//把当前节点及其所有左节点入栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		//取出栈顶，值放入队列，切换到其的右节点继续外层for循环寻找左节点
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, root.Val)
		root = root.Right
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
