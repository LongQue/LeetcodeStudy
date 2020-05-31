package main

/*
[142]环形链表 II
//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
//
// 说明：不允许修改给定的链表。
//
//
//
// 示例 1：
//
// 输入：head = [3,2,0,-4], pos = 1
//输出：tail connects to node index 1
//解释：链表中有一个环，其尾部连接到第二个节点。
//
//
//
//
// 示例 2：
//
// 输入：head = [1,2], pos = 0
//输出：tail connects to node index 0
//解释：链表中有一个环，其尾部连接到第一个节点。
//
//
//
//
// 示例 3：
//
// 输入：head = [1], pos = -1
//输出：no cycle
//解释：链表中没有环。
//
//
//
//
//
//
// 进阶：
//你是否可以不用额外空间解决此题？
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
func detectCycle(head *ListNode) *ListNode {
	pre, last := head, head
	//快慢指针找到节点
	for {
		if last == nil || last.Next == nil {
			return nil
		}
		last = last.Next.Next
		pre = pre.Next
		if pre == last {
			break
		}
	}
	//快指针从头开始，一步步走 和慢指针汇合即为结果
	// 例如  1->2->3->4->2   第一次快走两步，慢走一步，会在4相遇，
	//  第二次，last=head 大家都走一步步走，会在2相遇。
	//原理，pre当前的位置，往开始数，达到2的部分，是第一次中共同走过的部分，如果快从头走到这个位置，刚好是慢在环上走到这个位置的距离
	last = head
	for last != pre {
		last = last.Next
		pre = pre.Next
	}
	return pre
}

//leetcode submit region end(Prohibit modification and deletion)
