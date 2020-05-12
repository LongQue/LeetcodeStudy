package main

import "container/heap"

/*
[23]合并K个排序链表
//合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
//
// 示例:
//
// 输入:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//输出: 1->1->2->3->4->4->5->6
// Related Topics 堆 链表 分治算法
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//最小堆 或者 链表两两合并
type PQ []*ListNode

func (p PQ) Len() int { return len(p) }
func (p PQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PQ) Less(i, j int) bool {
	return p[i].Val < p[j].Val
}

func (p *PQ) Push(x interface{}) {
	node := x.(*ListNode)
	*p = append(*p, node)
}

func (p *PQ) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &ListNode{
		Val:  -1,
		Next: nil,
	}
	t := h
	if len(lists) == 0 {
		return h.Next
	}

	pq := make(PQ, 0)
	for i, _ := range lists {
		if lists[i] != nil {
			pq = append(pq, lists[i])
		}
	}
	heap.Init(&pq)

	for len(pq) > 0 {
		item := heap.Pop(&pq).(*ListNode)

		t.Next = item
		t = t.Next

		if item.Next != nil {
			heap.Push(&pq, item.Next)
		}
	}

	return h.Next

}

//leetcode submit region end(Prohibit modification and deletion)
