package main

/*
[2]两数相加
//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例：
//
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//
// Related Topics 链表 数学
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
	b := 0
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	result := dummy
	for l1 != nil || l2 != nil {
		temp := 0
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}
		result.Next = &ListNode{
			Val:  (temp + b) % 10,
			Next: nil,
		}
		b = (temp + b) / 10
		result = result.Next
	}
	if b == 1 {
		result.Next = &ListNode{
			Val:  b,
			Next: nil,
		}
	}
	return dummy.Next
}
func pow(a int) int {
	sum := 1
	for i := 0; i < a; i++ {
		sum = sum * 10
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
