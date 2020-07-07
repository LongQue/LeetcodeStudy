package main

/*
[145]二叉树的后序遍历
//给定一个二叉树，返回它的 后序 遍历。
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
//输出: [3,2,1]
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树
// 👍 341 👎 0
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
func postorderTraversal(root *TreeNode) []int {
	//从下到上  从左到右
	//stack反过来遍历，从上到下，从右到左
	var stack []*TreeNode
	//从头添加结果
	var result []int
	if root == nil {
		return result
	}

	stack = append(stack, root)
	for len(stack) != 0 {
		//取末尾实现从右到左
		end := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append([]int{end.Val}, result...)
		if end.Left != nil {
			stack = append(stack, end.Left)
		}
		if end.Right != nil {
			stack = append(stack, end.Right)
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
