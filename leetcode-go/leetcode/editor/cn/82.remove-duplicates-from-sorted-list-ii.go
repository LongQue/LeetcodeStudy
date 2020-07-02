package main

/*
[82]删除排序链表中的重复元素 II
//给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
//
// 示例 1:
//
// 输入: 1->2->3->3->4->4->5
//输出: 1->2->5
//
//
// 示例 2:
//
// 输入: 1->1->1->2->3
//输出: 2->3
// Related Topics 链表
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {

	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	cur := head
	//pre保持在cur前
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			//有相同的情况下移动的cur，找到cur.Next.Val不同或者cur.Next==nil
			cur = cur.Next
			for cur.Next != nil && cur.Next.Val == cur.Val {
				cur = cur.Next
			}
			//跨过cur这段，pre.Next直接指向cur.Next
			pre.Next = cur.Next
			//保持pre在cur前
			cur = pre.Next
		} else {
			//保持pre在cur前
			pre = cur
			cur = cur.Next
		}
	}
	return dummy.Next

}

//leetcode submit region end(Prohibit modification and deletion)
