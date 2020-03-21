package main

import "strings"

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
func findNumberIn2DArray(matrix [][]int,target int) bool {

	if len(matrix)==0{
		return false
	}
	row,col := 0,len(matrix[0]) - 1
	for col > -1 && row < len(matrix) {
		if matrix[row][col]==target {
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
	if head== nil {
		return []int{}
	}
	var res []int
	for head!= nil {
		res=append(res,head.Val)
		head=head.Next
	}
	for i,j:=0,len(res)-1;i<j; {
		res[i],res[j]=res[j],res[i]
		i++
		j--
	}
	return res
}

func main(){

}


