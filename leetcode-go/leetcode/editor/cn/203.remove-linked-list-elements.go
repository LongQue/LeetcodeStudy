package main

/*
[203]移除链表元素
//删除链表中等于给定值 val 的所有节点。
//
// 示例:
//
// 输入: 1->2->6->3->4->5->6, val = 6
//输出: 1->2->3->4->5
//
// Related Topics 链表
// 👍 405 👎 0
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	for head != nil {
		if head.Val == val {
			pre.Next = head.Next
			head = pre.Next
			continue
		}
		pre = head
		head = head.Next
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
