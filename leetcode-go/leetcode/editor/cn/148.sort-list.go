package main

/*
[148]排序链表
//在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
//
// 示例 1:
//
// 输入: 4->2->1->3
//输出: 1->2->3->4
//
//
// 示例 2:
//
// 输入: -1->5->3->4->0
//输出: -1->0->3->4->5
// Related Topics 排序 链表
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	//归并排序
	if head == nil || head.Next == nil {
		return head
	}
	//快慢指针找中间值
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	temp := slow.Next
	//断开链表，分而治之
	slow.Next = nil
	//反复断开链表之后
	left := sortList(head)
	right := sortList(temp)
	h := &ListNode{
		Val: 0,
	}
	//记录虚构 便于返回
	res := h
	for left != nil && right != nil {
		//比较连接
		if left.Val < right.Val {
			h.Next = left
			left = left.Next
		} else {
			h.Next = right
			right = right.Next
		}
		h = h.Next
	}
	//剩余的一次接上
	if left != nil {
		h.Next = left
	} else {
		h.Next = right
	}
	//返回拼接好的头指针
	return res.Next
}

//leetcode submit region end(Prohibit modification and deletion)
