package main

/*
[234]回文链表
//请判断一个链表是否为回文链表。
//
// 示例 1:
//
// 输入: 1->2
//输出: false
//
// 示例 2:
//
// 输入: 1->2->2->1
//输出: true
//
//
// 进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
// Related Topics 链表 双指针
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast, slow := head, head
	var p *ListNode
	var pre *ListNode
	//快慢指针，慢指针部分翻转链表，前半部以p为链表头
	for fast != nil && fast.Next != nil {
		p = slow
		slow = slow.Next
		fast = fast.Next.Next

		p.Next = pre
		pre = p
	}
	if fast != nil {
		slow = slow.Next
	}
	//从p 和后半部开头slow开始遍历回文
	for p != nil {
		if p.Val != slow.Val {
			return false
		}
		p = p.Next
		slow = slow.Next
	}
	return true

}

//leetcode submit region end(Prohibit modification and deletion)
