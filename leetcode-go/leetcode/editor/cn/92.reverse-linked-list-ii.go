package main

/*
[92]反转链表 II
//反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
//
// 说明:
//1 ≤ m ≤ n ≤ 链表长度。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
//输出: 1->4->3->2->5->NULL
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
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := dummy
	i := 1
	for i < m {
		pre = pre.Next
		i++
	}
	cur := pre.Next
	for i < n {
		temp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = temp
		i++
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
