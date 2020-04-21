package main

import (
	"container/heap"
	"errors"
	"math"
	"sort"
	"strconv"
	"strings"
)

/**
3.数组中的重复数字
在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
*/
func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i]] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
		if nums[i] != i {
			return nums[i]
		}
	}
	return -1
}

/**
4.二维数组中的查找
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findNumberIn2DArray(matrix [][]int, target int) bool {

	if len(matrix) == 0 {
		return false
	}
	row, col := 0, len(matrix[0])-1
	for col > -1 && row < len(matrix) {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}

/**
5. 替换空格
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
示例 1：
输入：s = "We are happy."
输出："We%20are%20happy."

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func replaceSpace(s string) string {
	replace := strings.Replace(s, " ", "%20", -1)
	return replace
}

/**
6、从尾到头打印链表
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
示例 1：
输入：head = [1,3,2]
输出：[2,3,1]
限制：
0 <= 链表长度 <= 10000
先顺序读取出来O(n)，在翻转数组O(n/2)=O(n)
*/
func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

/**
7.重建二叉树
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
例如，给出
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
限制：
0 <= 节点个数 <= 5000
题解：以前序遍历为顺序，每次取前序第一个节点，在中序遍历中找相等的值，该值左边即为当前节点的左子树，右边为右子树
*/
func buildTree(preorder []int, inorder []int) (Head *TreeNode) {
	if len(inorder) == 0 {
		return nil
	}
	vp := preorder[0]
	for ii, vi := range inorder {
		if vp == vi {
			Head := &TreeNode{
				Val:   vp,
				Left:  buildTree(preorder[1:], inorder[:ii]),
				Right: buildTree(preorder[1:], inorder[ii+1:]),
			}
			return Head
		}
	}

	return
}

/**
09.用两个栈实现队列
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
示例 1：
输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
示例 2：
输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
提示：
1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用
*/
type CQueue struct {
	in  []int
	out []int
}

func Constructor() CQueue {
	return CQueue{
		make([]int, 0, 5),
		make([]int, 0, 5),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) == 0 {
		if len(this.in) == 0 {
			return -1
		} else {
			for len(this.in) != 0 {
				i := this.in[len(this.in)-1]
				this.in = this.in[:len(this.in)-1]
				this.out = append(this.out, i)
			}
		}
	}
	i := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return i
}

/**
10-1斐波那契数列
写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项。斐波那契数列的定义如下：
F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
示例 1：
输入：n = 2
输出：1

示例 2：
输入：n = 5
输出：5
提示：
0 <= n <= 100

暂存取代递归
*/
func fib(n int) int {
	if n < 2 {
		return n
	}
	a, b, c := 0, 1, 0
	for i := 2; i <= n; i++ {
		c = (a + b) % 1000000007
		a = b
		b = c
	}
	return c
}

/**
面试题10- II. 青蛙跳台阶问题
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
示例 1：
输入：n = 2
输出：2
示例 2：

输入：n = 7
输出：21
提示：
0 <= n <= 100

n=0也算跳一次，这是最骚的
*/
func numWays(n int) int {
	if n < 2 {
		return 1
	}
	a, b, c := 1, 1, 0
	for i := 2; i <= n; i++ {
		c = (a + b) % 1000000007
		a = b
		b = c
	}
	return c
}

/**
面试题11. 旋转数组的最小数字
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。
示例 1：
输入：[3,4,5,1,2]
输出：1

示例 2：
输入：[2,2,2,0,1]
输出：0
注意：本题与主站 154 题相同：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/

注意排除三者left mid right都相同的情况
*/
func minArray(nums []int) int {
	left, right := 0, len(nums)-1
	mid := left
	for nums[left] >= nums[right] {
		if right-left == 1 {
			mid = right
			break
		}
		mid := (right-left)/2 + left
		//三者相同，遍历
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			result := nums[left]
			for left = left + 1; left < right; left++ {
				if result > nums[left] {
					result = nums[left]
				}
			}
			return result
		}
		if nums[left] <= nums[mid] {
			left = mid
		} else {
			right = mid
		}
	}
	return nums[mid]
}

/**
面试题12. 矩阵中的路径
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在
矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。
[["a","b","c","e"],
["s","f","c","s"],
["a","d","e","e"]]
但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
示例 1：
输入：board = [["A","B","C","E"],
              ["S","F","C","S"],
		      ["A","D","E","E"]], word = "ABCCED"
输出：true

示例 2：
输入：board = [["a","b"],
              ["c","d"]], word = "abcd"
输出：false
提示：
1 <= board.length <= 200
1 <= board[i].length <= 200

判断不为空，建同样大小的临时表(bool）记录每个走过的位置（标记位true),如果为true则不走该节点。
进入下一个节点先标记位true，如果回退，在改回false
*/
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	words := []byte(word)
	tempMap := make([][]bool, len(board))
	for k, v := range board {
		tempMap[k] = make([]bool, len(v))
	}
	curLen := 0
	getPath := false

	var walk func(r, c int)
	walk = func(r, c int) {
		if getPath == true || board[r][c] != words[curLen] {
			return
		}
		tempMap[r][c] = true
		curLen++
		if curLen == len(words) {
			getPath = true
			return
		}
		if r-1 >= 0 && tempMap[r-1][c] == false {
			walk(r-1, c)
		}
		if c-1 >= 0 && tempMap[r][c-1] == false {
			walk(r, c-1)
		}
		if r+1 < len(board) && tempMap[r+1][c] == false {
			walk(r+1, c)
		}
		if c+1 < len(board[0]) && tempMap[r][c+1] == false {
			walk(r, c+1)
		}
		tempMap[r][c] = false
		curLen--
	}
	for r, v := range board {
		for c := range v {
			if getPath == true {
				return true
			}
			walk(r, c)
		}
	}
	return getPath
}

/**
面试题13. 机器人的运动范围
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格
（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
示例 1：
输入：m = 2, n = 3, k = 1
输出：3

示例 2：
输入：m = 3, n = 1, k = 0
输出：1
提示：
1 <= n,m <= 100
0 <= k <= 20
*/

func movingCount(m int, n int, k int) int {
	tempMap := make([][]bool, m)
	for k, _ := range tempMap {
		tempMap[k] = make([]bool, n)
	}

	//计算位数之和
	var getDigitSum func(num int) int
	getDigitSum = func(num int) int {
		var sum int
		for num > 0 {
			sum += n % 10
			n = n / 10
		}
		return sum
	}
	//检查是否可用
	var check func(curM, curN int) bool
	check = func(curM, curN int) bool {
		if curM < m && curM >= 0 && curN < n && curN >= 0 && (getDigitSum(curN)+getDigitSum(curM)) <= k && tempMap[curM][curN] != true {
			return true
		}
		return false
	}
	//移动到当前
	var walk func(m, n, count, curM, curN int, tempMap [][]bool) int
	walk = func(m, n, count, curM, curN int, tempMap [][]bool) int {
		if check(curM, curN) {
			tempMap[curM][curN] = true
			count = 1 + walk(m, n, count, curM-1, curN, tempMap) + walk(m, n, count, curM+1, curN, tempMap) + walk(m, n, count, curM, curN-1, tempMap) + walk(m, n, count, curM, curN+1, tempMap)
		}
		return count
	}
	count := walk(m, n, 0, 0, 0, tempMap)
	return count

}

/**
面试题14- I. 剪绳子
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]
...k[m] 。请问 k[0]*k[1]*...*k[m] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3
的三段，此时得到的最大乘积是18。

示例 1：
输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1

示例 2:
输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
提示：
2 <= n <= 58
*/
func cuttingRope(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	result := []int{0, 1, 2, 3}
	var max int
	for i := 4; i <= n; i++ {
		max = 0
		for j := 1; j <= i/2; j++ {
			temp := result[j] * result[i-j]
			if max < temp {
				max = temp
			}
		}
		result = append(result, max)
	}
	return result[n]
}

/**
面试题14- II. 剪绳子 II

给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m] 。
请问 k[0]*k[1]*...*k[m] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

示例 1：
输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1

示例 2:
输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
提示：
2 <= n <= 1000
*/
func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}
	temp, a, b, c := 1, n/3, n%3, 1000000007
	var num func(a int) int
	num = func(a int) int {
		for i := 0; i < a; i++ {
			temp = (temp * 3) % c
		}
		return temp
	}
	if b == 2 {
		temp = num(a)
		return (temp * 2) % c
	} else if b == 1 {
		temp = num(a - 1)
		return (temp * 4) % c
	} else {
		return num(a) % c
	}
}

/**
面试题15. 二进制中1的个数
请实现一个函数，输入一个整数，输出该数二进制表示中 1 的个数。例如，把 9 表示成二进制是 1001，有 2 位是 1。
因此，如果输入 9，则该函数输出 2。

示例 1：
输入：00000000000000000000000000001011
输出：3
解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。

示例 2：
输入：00000000000000000000000010000000
输出：1
解释：输入的二进制串 00000000000000000000000010000000 中，共有一位为 '1'。

示例 3：
输入：11111111111111111111111111111101
输出：31
解释：输入的二进制串 11111111111111111111111111111101 中，共有 31 位为 '1'。
*/
func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}

/**
面试题16. 数值的整数次方
实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大
数问题。

示例 1:
输入: 2.00000, 10
输出: 1024.00000

示例 2:
输入: 2.10000, 3
输出: 9.26100

示例 3:
输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25
说明:
-100.0 < x < 100.0
n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。
*/
func myPow(x float64, n int) float64 {
	if n < 0 {
		return myPow16(1.0/x, -n)
	} else if n == 0 {
		return 1
	} else {
		return myPow16(x, n)
	}

}
func myPow16(x float64, n int) float64 {
	if n == 1 {
		return x
	} else if n&1 == 0 {
		return myPow16(x*x, n>>1)
	} else {
		return x * myPow16(x, n-1)
	}
}

/**
面试题17. 打印从1到最大的n位数
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

示例 1:
输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]

说明：
用返回一个整数列表来代替打印
n 为正整数
*/
func printNumbers(n int) []int {
	if n == 0 {
		return []int{}
	}
	sum := 1
	for n > 0 {
		sum = sum * 10
		n--
	}
	result := make([]int, 0, sum)

	for i := 1; i < sum; i++ {
		result = append(result, i)
	}
	return result
}

/**
面试题18. 删除链表的节点
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
返回删除后的链表的头节点。
注意：此题对比原题有改动

示例 1:
输入: head = [4,5,1,9], val = 5
输出: [4,1,9]
解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.

示例 2:
输入: head = [4,5,1,9], val = 1
输出: [4,5,9]
解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.

说明：
题目保证链表中节点的值互不相同
若使用 C 或 C++ 语言，你不需要 free 或 delete 被删除的节点
*/
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		head = head.Next
		return head
	}
	var temp = head
	for temp.Next != nil {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
			break
		}
		temp = temp.Next
	}
	return head
}

/**
面试题19. 正则表达式匹配（此题不会
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
说明:
s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。

示例 1:
输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。

示例 2:
输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。

示例 3:
输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。

示例 4:
输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。

示例 5:
输入:
s = "mississippi"
p = "mis*is*p*."
输出: false
*/
func isMatch(s string, p string) bool {
	slen, plen := len(s), len(p)
	if slen == 0 {
		if plen%2 != 0 {
			return false
		}
		for j := 1; j < plen; j += 2 {
			if p[j] != '*' {
				return false
			}
		}
		return true
	}
	if plen == 0 {
		return slen == 0
	}
	dp := make([][]bool, slen, slen)
	for i := 0; i < slen; i++ {
		dp[i] = make([]bool, plen, plen)
	}
	// init: dp(0, j)
	dp[0][0] = p[0] == s[0] || p[0] == '.'
	for canPreEmpty, j := true, 1; j < plen; j += 1 {
		if canPreEmpty && j%2 == 1 && p[j] != '*' {
			canPreEmpty = false
		}
		switch p[j] {
		case '.':
			dp[0][j] = canPreEmpty
		case '*':
			if p[j-1] == s[0] || p[j-1] == '.' {
				dp[0][j] = canPreEmpty || (j >= 2 && dp[0][j-2])
			} else {
				dp[0][j] = j-2 >= 0 && dp[0][j-2]
			}
		default:
			dp[0][j] = s[0] == p[j] && canPreEmpty
		}
	}
	// cal: dp(i, j)
	for i := 1; i < slen; i++ {
		for j := 1; j < plen; j++ {
			switch p[j] {
			case '.':
				dp[i][j] = dp[i-1][j-1]
			case '*':
				if p[j-1] == s[i] || p[j-1] == '.' {
					dp[i][j] = (j-2 >= 0 && dp[i][j-2]) || dp[i-1][j]
				} else {
					dp[i][j] = j-2 >= 0 && dp[i][j-2]
				}
			default:
				dp[i][j] = s[i] == p[j] && dp[i-1][j-1]
			}
		}
	}
	return dp[slen-1][plen-1]
}

/**
面试题20. 表示数值的字符串（此题不会
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"0123"及"
-1E-16"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。
注意：本题与主站 65 题相同：https://leetcode-cn.com/problems/valid-number/
*/
func isNumber(s string) bool {
	_, err := strconv.ParseFloat(strings.Trim(s, " "), 64)
	return err == nil || errors.Is((interface{}(err).(*strconv.NumError)).Err, strconv.ErrRange)
}

/**
面试题21. 调整数组顺序使奇数位于偶数前面
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

示例：
输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。

提示：
1 <= nums.length <= 50000
1 <= nums[i] <= 10000
*/
func exchange(nums []int) []int {

	for a, b := 0, len(nums)-1; ; {
		for a < b && (nums[a]&1 == 1) {
			a++
		}
		for a < b && (nums[b]&1 == 0) {
			b--
		}
		if a < b {
			nums[a], nums[b] = nums[b], nums[a]
		} else {
			return nums
		}
	}
}

/**
面试题22. 链表中倒数第k个节点
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1
个节点。例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是
值为4的节点。

示例：

给定一个链表: 1->2->3->4->5, 和 k = 2.

返回链表 4->5.
*/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	first, second := head, head
	for i := 0; i < k; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	return second
}

/**
面试题24. 反转链表
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

示例:
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

限制：
0 <= 节点个数 <= 5000
*/
func reverseList(head *ListNode) *ListNode {
	var pre, next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

/**
面试题25. 合并两个排序的链表
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

限制：
0 <= 链表长度 <= 1000
注意：本题与主站 21 题相同：https://leetcode-cn.com/problems/merge-two-sorted-lists/
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = &ListNode{
				Val:  l1.Val,
				Next: nil,
			}
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{
				Val:  l2.Val,
				Next: nil,
			}
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}

/**
面试题26. 树的子结构
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:
     3
    / \
   4   5
  / \
 1   2
给定的树 B：

   4
  /
 1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：

输入：A = [1,2,3], B = [3,1]
输出：false
示例 2：

输入：A = [3,4,5,1,2], B = [4,1]
输出：true
限制：

0 <= 节点个数 <= 10000
*/
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	result := false
	if A != nil && B != nil {
		if A.Val == B.Val {
			result = isSubStructure2(A, B)
		}
		if !result {
			result = isSubStructure(A.Left, B)
		}
		if !result {
			result = isSubStructure(A.Right, B)
		}
	}
	return result
}

/* 从该节点开始判断是否是子结构*/
func isSubStructure2(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return isSubStructure2(A.Left, B.Left) && isSubStructure2(A.Right, B.Right)
}

/**
面试题27. 二叉树的镜像
请完成一个函数，输入一个二叉树，该函数输出它的镜像。
例如输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

示例 1：
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]
限制：
0 <= 节点个数 <= 1000
*/
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	if root.Left != nil {
		root.Left = mirrorTree(root.Left)
	}
	if root.Right != nil {
		root.Right = mirrorTree(root.Right)
	}
	return root
}

/**
面试题28. 对称的二叉树
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3
示例 1：
输入：root = [1,2,2,3,4,4,3]
输出：true

示例 2：
输入：root = [1,2,2,null,3,null,3]
输出：false
限制：
0 <= 节点个数 <= 1000
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetric2(root.Left, root.Right)
}

func isSymmetric2(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right != nil && left.Val == right.Val {
		return isSymmetric2(left.Left, right.Right) && isSymmetric2(left.Right, right.Left)
	}
	return false
}

/**
面试题29. 顺时针打印矩阵
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

示例 2：
输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

限制：
0 <= matrix.length <= 100
0 <= matrix[i].length <= 100
*/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	step := 0
	size := len(matrix) * len(matrix[0])
	top, bottom, right, left := 0, len(matrix)-1, len(matrix[0])-1, 0
	nums := make([]int, size)
	for step < size {
		for i := left; i <= right && step < size; i++ {
			nums[step] = matrix[top][i]
			step++
		}
		top++
		for i := top; i <= bottom && step < size; i++ {
			nums[step] = matrix[i][right]
			step++
		}
		right--
		for i := right; i >= left && step < size; i-- {
			nums[step] = matrix[bottom][i]
			step++
		}
		bottom--
		for i := bottom; i >= top && step < size; i-- {
			nums[step] = matrix[i][left]
			step++
		}
		left++
	}
	return nums
}

/**
面试题30. 包含min函数的栈
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

示例:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.

提示：
各函数的调用总次数不超过 20000 次

用辅助栈记录主栈每个元素时的最小值
*/
type MinStack struct {
	slave  []int
	master []int
}

/** initialize your data structure here. */
func Constructor1() MinStack {
	return MinStack{
		slave:  make([]int, 0, 10),
		master: make([]int, 0, 10),
	}
}

func (this *MinStack) Push(x int) {
	this.master = append(this.master, x)
	if len(this.slave) == 0 {
		this.slave = append(this.slave, x)
	} else if this.slave[len(this.slave)-1] < x {
		this.slave = append(this.slave, this.slave[len(this.slave)-1])
	} else {
		this.slave = append(this.slave, x)
	}

}

func (this *MinStack) Pop() {
	this.master = this.master[:len(this.master)-1]
	this.slave = this.slave[:len(this.slave)-1]
}

func (this *MinStack) Top() int {
	return this.master[len(this.master)-1]
}

func (this *MinStack) Min() int {
	return this.slave[len(this.slave)-1]

}

/**
面试题31. 栈的压入、弹出序列
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列
{1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。

示例 1：
输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1

示例 2：
输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。

提示：
0 <= pushed.length == popped.length <= 1000
0 <= pushed[i], popped[i] < 1000
pushed 是 popped 的排列。
*/
func validateStackSequences(pushed []int, popped []int) bool {
	temp := make([]int, 0, len(pushed))

	popIndex := 0
	for _, v := range pushed {
		if v == popped[popIndex] {
			popIndex++
			for len(temp) != 0 {
				if temp[len(temp)-1] == popped[popIndex] {
					temp = temp[:len(temp)-1]
					popIndex++
				} else {
					break
				}
			}
		} else {
			temp = append(temp, v)
		}
	}

	return len(temp) == 0
}

/**
面试题32 - I. 从上到下打印二叉树
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：
[3,9,20,15,7]
提示：
节点总数 <= 1000
*/
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		var temp []*TreeNode
		for _, v := range queue {
			result = append(result, v.Val)
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		queue = temp
	}
	return result
}

/**
面试题32 - II. 从上到下打印二叉树 II
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

例如:
给定二叉树: [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：
[
  [3],
  [9,20],
  [15,7]
]
提示：
节点总数 <= 1000
*/
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var tempInt []int
		var tempQueue []*TreeNode
		for _, v := range queue {
			tempInt = append(tempInt, v.Val)
			if v.Left != nil {
				tempQueue = append(tempQueue, v.Left)
			}
			if v.Right != nil {
				tempQueue = append(tempQueue, v.Right)
			}
		}
		queue = tempQueue
		result = append(result, tempInt)
	}
	return result

}

/**
面试题32 - III. 从上到下打印二叉树 III
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

例如:
给定二叉树: [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7

返回其层次遍历结果：
[
  [3],
  [20,9],
  [15,7]
]
提示：
节点总数 <= 1000
*/
func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		var tempInt []int
		var tempQueue []*TreeNode
		for _, v := range queue {
			tempInt = append(tempInt, v.Val)
			if v.Left != nil {
				tempQueue = append(tempQueue, v.Left)
			}
			if v.Right != nil {
				tempQueue = append(tempQueue, v.Right)
			}
		}
		queue = tempQueue
		if level&1 == 1 {
			for i := 0; i < len(tempInt)/2; i++ {
				tempInt[i], tempInt[len(tempInt)-1-i] = tempInt[len(tempInt)-1-i], tempInt[i]
			}
		}
		result = append(result, tempInt)
		level++
	}
	return result
}

/**
面试题33. 二叉搜索树的后序遍历序列
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。

参考以下这颗二叉搜索树：
     5
    / \
   2   6
  / \
 1   3

示例 1：
输入: [1,6,3,2,5]
输出: false

示例 2：
输入: [1,3,2,6,5]
输出: true

提示：
数组长度 <= 1000
*/
func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}

	root := postorder[len(postorder)-1]
	i := 0
	for postorder[i] < root {
		i++
	}
	mid := i
	for ; i < len(postorder)-1; i++ {
		if postorder[i] < root {
			return false
		}
	}
	return verifyPostorder(postorder[:mid]) && verifyPostorder(postorder[mid:len(postorder)-1])

}

/**
面试题34 二叉树中和为某一值的路径
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
示例:
给定如下二叉树，以及目标和 sum = 22，
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:
[
   [5,4,11,2],
   [5,8,4,5]
]
深度遍历,注意sum未必为正，和节点值之间没有大于或小于关系，相加操作的结果可能在sum上下徘徊
*/
func pathSum(root *TreeNode, sum int) [][]int {
	var result [][]int
	pathSumWork(root, sum, 0, []int{}, &result)
	return result
}
func pathSumWork(root *TreeNode, sum, now int, temp []int, result *[][]int) {
	if root == nil || sum < now+root.Val {
		return
	}
	now = now + root.Val
	temp = append(temp, root.Val)
	if now == sum && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		*result = append(*result, tmp)
		return
	}
	pathSumWork(root.Left, sum, now, temp, result)
	pathSumWork(root.Right, sum, now, temp, result)
}

/**
面试题35. 复杂链表的复制
无 GO
*/

/**
面试题38. 字符串的排列
输入一个字符串，打印出该字符串中字符的所有排列。
你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
示例:
输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]

限制：
1 <= s 的长度 <= 8
*/
func permutation(s string) []string {
	var res []string
	permutationWork([]byte(s), 0, &res)
	return res
}
func permutationWork(s []byte, start int, res *[]string) {
	if start == len(s) {
		*res = append(*res, string(s))
	}
	m := make(map[byte]int)
	for i := start; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			continue
		}
		s[i], s[start] = s[start], s[i]
		permutationWork(s, start+1, res)
		s[i], s[start] = s[start], s[i]
		m[s[i]] = 1
	}

}

/**
面试题39. 数组中出现次数超过一半的数字
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
示例 1:
输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2

result 记录nums[0]，num记录次数，从nums[1]开始遍历，遇到result相同的num+1，否则减一，num==0时，result换成当前nums[i]
限制：
1 <= 数组长度 <= 50000
*/
func majorityElement(nums []int) int {
	result, num := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if result == nums[i] {
			num++
		} else {
			num--
			if num == 0 {
				result = nums[i]
				num = 1
			}
		}
	}
	return result
}

/**
面试题40. 最小的k个数
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

示例 1：
输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]

示例 2：
输入：arr = [0,1,2,1], k = 1
输出：[0]

限制：
0 <= k <= arr.length <= 10000
0 <= arr[i] <= 10000
*/
func getLeastNumbers(arr []int, k int) []int {
	if k >= len(arr) {
		return arr
	}

	var recur func(arr []int, k int)
	recur = func(arr []int, k int) {

		s, e := 0, len(arr)-1
		c := arr[s]
		for s < e {
			for s < e {
				if arr[s] > c {
					break
				}
				s++
			}

			for e > s {
				if arr[e] <= c {
					break
				}
				e--
			}
			if e > s {
				arr[s], arr[e] = arr[e], arr[s]
			}
		}
		if arr[e] < arr[0] {
			arr[0], arr[e] = arr[e], arr[0]
		}
		if e > k {
			recur(arr[0:e], k)
		} else if e < k {
			recur(arr[e:], k-e)
		} else {
			return
		}
	}
	recur(arr, k)
	return arr[0:k]
}

/**
面试题41. 数据流中的中位数
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

例如，
[2,3,4] 的中位数是 3
[2,3] 的中位数是 (2 + 3) / 2 = 2.5
设计一个支持以下两种操作的数据结构：
void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。

示例 1：
输入：
["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
[[],[1],[2],[],[3],[]]
输出：[null,null,null,1.50000,null,2.00000]

示例 2：
输入：
["MedianFinder","addNum","findMedian","addNum","findMedian"]
[[],[2],[],[3],[]]
输出：[null,null,2.00000,null,2.50000]

限制：
最多会对 addNum、findMedia进行 50000 次调用。

最大堆和最小堆
*/

type MaxHeap []int
type MinHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	lower MaxHeap
	upper MinHeap
}

/** initialize your data structure here. */
func Constructor3() MedianFinder {
	return MedianFinder{make([]int, 0), make([]int, 0)}
}

func (this *MedianFinder) AddNum(num int) { // len(lower) >= len(upper)
	if len(this.lower) == 0 || this.lower[0] >= num {
		heap.Push(&this.lower, num)
	} else {
		heap.Push(&this.upper, num)
	}

	if len(this.upper) > len(this.lower) {
		v := heap.Pop(&this.upper)
		heap.Push(&this.lower, v)
	} else if len(this.lower) > len(this.upper)+1 {
		v := heap.Pop(&this.lower)
		heap.Push(&this.upper, v)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.lower) > len(this.upper) {
		return float64(this.lower[0])
	} else {
		return float64(this.lower[0]+this.upper[0]) / 2.0
	}
}

/**
面试题42. 连续子数组的最大和
输入一个整型数组，数组里有正数也有负数。数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
要求时间复杂度为O(n)。

示例1:
输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

提示：
1 <= arr.length <= 10^5
-100 <= arr[i] <= 100

找出以每个元素结尾时的和

if i==0||f(i-1)<=0
   f(i)=nums[i]
else if i>0&&f(i-1)>0
	f(i)=f(i-1)+nums[i]
*/
func maxSubArray(nums []int) int {
	max := -101
	temp := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || temp <= 0 {
			temp = nums[i]
		} else if i > 0 && temp > 0 {
			temp = temp + nums[i]
		}
		if temp > max {
			max = temp
		}
	}
	return max
}

/**
面试题43. 1～n整数中1出现的次数
输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。
例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。

示例 1：
输入：n = 12
输出：5

示例 2：
输入：n = 13
输出：6

限制：
1 <= n < 2^31

某位数分三种，
1、 若百位等于0，如12045，则看高位(12)，在所有数中百位是1的情况有 100-199，1100-1199,2100-2199...共计12*100
2、 若百位等于1，如12145，则看高低位(12,45)，在所有数中，百位出现1的情况 高位有12*100， 低位有45+1
3、 若百位大于1，如12345，则看高位（12），在所有数中，百位出现1的情况再上述的基础上加上 12100-12199，即（12+1）*100个
*/
func countDigitOne(n int) int {
	if n == 0 {
		return 0
	}
	i, count := 1, 0
	for n/i != 0 {
		left, cur, right := n/(10*i), (n/i)%10, n-(n/i)*i
		if cur == 0 {
			count += left * i
		} else if cur == 1 {
			count += left*i + right + 1
		} else {
			count += (left + 1) * i
		}
		i *= 10
	}
	return count
}

/**
面试题44. 数字序列中某一位的数字
数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，
第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。
请写一个函数，求任意第n位对应的数字。

示例 1：
输入：n = 3
输出：3

示例 2：
输入：n = 11
输出：0

限制：
0 <= n < 2^31
去掉特例0
判断输入的n大约是落在几位数的范围，直接跳过前面部分
1-9      9位      10^(位数-1)*位数*9
10-99    180位
100-999  2700位
....
若 271，先去掉0得269，
>9,减9 得260
>180，减180 得80
<小于2700，说明271是三位数
计算是三位数第几个 80/位数=26
计算是数的第几位   80%位数=2
3位数第一个 100 加 26得到对于的数字 126
["1","2","6"] s[2]转换成int
*/
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	//去掉0
	n--
	i := 1
	for {
		temp := int(math.Pow(10, float64(i-1))) * i * 9
		if n <= temp {
			break
		}
		i++
		n -= temp
	}
	a, b := n/i, n%i
	f := int(math.Pow(10, float64(i-1))) + a
	s := strconv.Itoa(f)
	result := string(s[b])
	i, _ = strconv.Atoi(result)
	return i
}

/**
面试题45. 把数组排成最小的数
输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

示例 1:
输入: [10,2]
输出: "102"

示例 2:
输入: [3,30,34,5,9]
输出: "3033459"

提示:
0 < nums.length <= 100
说明:
输出结果可能非常大，所以你需要返回一个字符串而不是整数
拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0
*/
func minNumber(nums []int) string {
	str := make([]string, len(nums))
	for i := range nums {
		str[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(str, func(i, j int) bool {
		s12 := str[i] + str[j]
		s21 := str[j] + str[i]
		return s12 < s21
	})
	res := ""
	for _, v := range str {
		res += v
	}
	return res
}

/**

面试题46. 把数字翻译成字符串
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。
请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

示例 1:
输入: 12258
输出: 5

解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
提示：
0 <= num < 2^31

类似于青蛙台阶，对于最后一个添加的数，if它和前一个相加超过25或者小于10则说明只能跳一阶，只有一种跳法  f(x)=f(x-1)
else 则说明可以跳两次一阶或者跳2阶 f(x)=f(x-1)+f(x-2）
需要注意处理f(2）的特殊情况
*/
func translateNum(num int) int {
	left, right := 0, -1
	cur := 1
	last := cur
	for num > 0 {
		left = num % 10
		num /= 10
		if right >= 0 && ((left*10 + right) >= 10) && ((left*10 + right) <= 25) {
			cur, last = cur+last, cur
		} else {
			last = cur
		}
		right = left
	}
	return cur
}

/**
面试题47. 礼物的最大价值
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移
动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

示例 1:
输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 12
解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物

提示：
0 < grid.length <= 200
0 < grid[0].length <= 200

动态规划
f(i,j)=max( f(i-1,j),f(i,j-1) ) + grid[i][j]
无法保证每步最优，需要另外建一个二维数组保存每个f(i,j)的值
*/
func maxValue(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	maxValues := make([][]int, row)
	for i := 0; i < len(grid); i++ {
		maxValues[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			left, up := 0, 0
			if i > 0 {
				up = maxValues[i-1][j]
			}
			if j > 0 {
				left = maxValues[i][j-1]
			}
			maxValues[i][j] = Max(up, left) + grid[i][j]
		}
	}
	return maxValues[row-1][col-1]
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	i := [][]int{{1, 2, 5}, {3, 2, 1}}
	print(maxValue(i))
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
