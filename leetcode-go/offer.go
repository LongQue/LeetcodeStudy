package main

import (
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
	words := []byte (word)
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
	result :=make([]int,0,sum)

	for i := 1; i < sum; i++ {
		result =append(result,i)
	}
	return result
}
func main() {
	println(cuttingRope2(120))

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
