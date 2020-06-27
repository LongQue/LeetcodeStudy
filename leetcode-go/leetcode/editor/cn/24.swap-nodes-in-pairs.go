package main

/*
[24]两两交换链表中的节点
//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例:
//
// 给定 1->2->3->4, 你应该返回 2->1->4->3.
//
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
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	for head != nil && head.Next != nil {
		/**
		交换BC  pre=A  B=head
		A->B->C->D
		1. A->C->D  B->C->D
		2.B->D  A->C->D
		3.A->C->B->D
		4. pre=B head=D
		*/
		pre.Next = head.Next
		head.Next = head.Next.Next
		pre.Next.Next = head
		pre = head
		head = head.Next
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
