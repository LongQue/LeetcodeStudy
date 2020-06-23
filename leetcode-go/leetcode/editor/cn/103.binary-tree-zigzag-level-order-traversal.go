package main

/*
[103]二叉树的锯齿形层次遍历
//给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
// 例如：
//给定二叉树 [3,9,20,null,null,15,7],
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回锯齿形层次遍历如下：
//
// [
//  [3],
//  [20,9],
//  [15,7]
//]
//
// Related Topics 栈 树 广度优先搜索
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	res = append(res, []int{root.Val})
	//记录每层遍历的节点，初始化为第一层
	treeArrays := []*TreeNode{root}
	length := len(treeArrays)
	for length > 0 {
		var temp []int
		//倒叙遍历
		for i := length - 1; i >= 0; i-- {
			//双数层先右后左
			if treeArrays[i].Right != nil {
				treeArrays = append(treeArrays, treeArrays[i].Right)
				temp = append(temp, treeArrays[i].Right.Val)
			}
			if treeArrays[i].Left != nil {
				treeArrays = append(treeArrays, treeArrays[i].Left)
				temp = append(temp, treeArrays[i].Left.Val)
			}
		}
		//移除已遍历节点
		treeArrays = treeArrays[length:]
		length = len(treeArrays)
		//检查是否有剩余，下同
		if length == 0 {
			return res
		}
		res = append(res, temp)

		temp = []int{}
		//一样倒叙
		for i := length - 1; i >= 0; i-- {
			//单数层先左后右
			if treeArrays[i].Left != nil {
				treeArrays = append(treeArrays, treeArrays[i].Left)
				temp = append(temp, treeArrays[i].Left.Val)
			}
			if treeArrays[i].Right != nil {
				treeArrays = append(treeArrays, treeArrays[i].Right)
				temp = append(temp, treeArrays[i].Right.Val)
			}
		}
		treeArrays = treeArrays[length:]
		length = len(treeArrays)
		if length == 0 {
			return res
		}
		res = append(res, temp)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
