package main

/*
[25]K 个一组翻转链表
//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
//
//
// 示例：
//
// 给你这个链表：1->2->3->4->5
//
// 当 k = 2 时，应当返回: 2->1->4->3->5
//
// 当 k = 3 时，应当返回: 3->2->1->4->5
//
//
//
// 说明：
//
//
// 你的算法只能使用常数的额外空间。
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	//开始pre和end在同一起点
	pre := dummy
	end := dummy
	for end.Next != nil {
		//end移动k次，代表翻转k个，此时end指向该组翻转的最后一个
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		//如果为nil则说明没有够k个，直接返回结果即可
		if end == nil {
			break
		}
		//start指向该组翻转的第一个
		start := pre.Next
		//记录下一组翻转的第一个
		next := end.Next
		//nil是为了下边翻转的判断条件，使两组翻转断掉
		end.Next = nil
		//翻转前 pre->start->...end 断 next
		//翻转后 pre->start   end->...->start 断 next ，修改pre指向end
		pre.Next = reverse(start)
		//接上start next    pre->end->...start->next
		start.Next = next
		//pre end移动至start，变回循环开始状态
		pre = start
		end = pre
	}
	return dummy.Next
}
func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	//pre 断 cur
	//指针不停的改向左指，pre cur不断地向左移，知道cur为nil，pre移到尽头
	for cur != nil {
		// pre 断 cur->next
		next := cur.Next
		// pre<-cur 断 next
		cur.Next = pre
		// pre(cur) 断 next
		pre = cur
		// pre 断 cur
		cur = next
	}
	//返回pre
	return pre
}

//leetcode submit region end(Prohibit modification and deletion)
