package main

/*
[143]重排链表
//给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
//将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
// 示例 1:
//
// 给定链表 1->2->3->4, 重新排列为 1->4->2->3.
//
// 示例 2:
//
// 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
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
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	//快慢指针，断成两条，后一条翻转
	slow := head
	fast := head
	//长度相等，或者慢指针多1
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	newHead := slow.Next
	slow.Next = nil
	newHead = reorderListHelper(newHead)
	//拼接
	for newHead != nil {
		temp := newHead.Next
		newHead.Next = head.Next
		head.Next = newHead
		head = newHead.Next
		newHead = temp
	}
}

func reorderListHelper(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{
		Next: head,
	}
	cur := dummy.Next
	for cur.Next != nil {
		temp := dummy.Next
		dummy.Next = cur.Next
		cur.Next = cur.Next.Next
		dummy.Next.Next = temp
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
