package main

/*
[138]复制带随机指针的链表
//给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。
//
// 要求返回这个链表的 深拷贝。
//
// 我们用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：
//
//
// val：一个表示 Node.val 的整数。
// random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为 null 。
//
//
//
//
// 示例 1：
//
//
//
// 输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
//输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
//
//
// 示例 2：
//
//
//
// 输入：head = [[1,1],[2,1]]
//输出：[[1,1],[2,1]]
//
//
// 示例 3：
//
//
//
// 输入：head = [[3,null],[3,0],[3,null]]
//输出：[[3,null],[3,0],[3,null]]
//
//
// 示例 4：
//
// 输入：head = []
//输出：[]
//解释：给定的链表为空（空指针），因此返回 null。
//
//
//
//
// 提示：
//
//
// -10000 <= Node.val <= 10000
// Node.random 为空（null）或指向链表中的节点。
// 节点数目不超过 1000 。
//
// Related Topics 哈希表 链表
// 👍 311 👎 0
*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	ptr := head
	//将每个节点拷贝一份新串和旧串练一次，注意nil是不会拷贝的
	// A-A'-B-B'-C-C'-nil
	for ptr != nil {
		temp := &Node{
			Val:  ptr.Val,
			Next: ptr.Next,
		}
		ptr.Next = temp
		ptr = temp.Next
	}

	ptr = head
	//连接旧串的随机节点
	for ptr != nil {
		//小心nil
		if ptr.Random != nil {
			ptr.Next.Random = ptr.Random.Next
		} else {
			ptr.Next.Random = nil
		}
		ptr = ptr.Next.Next
	}

	old := head
	new := head.Next
	temp := new
	//分离新串旧串
	for old != nil {
		old.Next = old.Next.Next
		if new.Next != nil {
			new.Next = new.Next.Next
		} else {
			new.Next = nil
		}
		old = old.Next
		new = new.Next
	}
	return temp
}

//leetcode submit region end(Prohibit modification and deletion)
