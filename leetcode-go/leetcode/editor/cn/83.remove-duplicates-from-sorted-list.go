package main

/*
[83]删除排序链表中的重复元素
//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
//
// 示例 1:
//
// 输入: 1->1->2
//输出: 1->2
//
//
// 示例 2:
//
// 输入: 1->1->2->3->3
//输出: 1->2->3
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
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			temp := cur.Next
			for temp != nil && temp.Val == cur.Val {
				temp = temp.Next
			}
			cur.Next = temp
			//如果为nil则外层for就无法执行
			if temp == nil {
				break
			}
		}
		cur = cur.Next
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
