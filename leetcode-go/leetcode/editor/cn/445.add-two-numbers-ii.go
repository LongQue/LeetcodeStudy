package main

/*
[445]两数相加 II
//给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
//
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
//
//
// 进阶：
//
// 如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
//
//
//
// 示例：
//
// 输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 8 -> 0 -> 7
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//栈，先遍历出来
	var queue1 []int
	var queue2 []int
	for l1 != nil {
		queue1 = append(queue1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		queue2 = append(queue2, l2.Val)
		l2 = l2.Next
	}
	//记录进位
	carry := 0
	var head *ListNode
	//carry防止 两对列都为空，且刚好进位的情况
	for len(queue1) != 0 || len(queue2) != 0 || carry != 0 {
		sum := carry
		if len(queue1) != 0 {
			sum += queue1[len(queue1)-1]
			queue1 = queue1[:len(queue1)-1]
		}
		if len(queue2) != 0 {
			sum += queue2[len(queue2)-1]
			queue2 = queue2[:len(queue2)-1]
		}
		node := &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		node.Next = head
		head = node
		carry = sum / 10
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
