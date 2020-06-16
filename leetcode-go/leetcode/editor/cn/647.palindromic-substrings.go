package main

/*
[647]回文子串
//给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
//
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被计为是不同的子串。
//
// 示例 1:
//
//
//输入: "abc"
//输出: 3
//解释: 三个回文子串: "a", "b", "c".
//
//
// 示例 2:
//
//
//输入: "aaa"
//输出: 6
//说明: 6个回文子串: "a", "a", "a", "aa", "aa", "aaa".
//
//
// 注意:
//
//
// 输入的字符串长度不会超过1000。
//
// Related Topics 字符串 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func countSubstrings(s string) int {
	n := len(s)
	//count回文数，至少n个
	count := n
	//动态规划，行代表左指针，列代表右指针
	temp := make([][]int, n)
	for i := range temp {
		temp[i] = make([]int, n)
	}
	//i左指针，j右指针
	for i := n - 1; i >= 0; i-- {
		//对角线位置全部置为1，即单个
		temp[i][i] = 1
		//因为如果s[i]=s[j],则s[i~j]是不是回文，取决于s[(i+1)~(j-1)],所以从下往上，从左往右遍历，即贴近对角线出发
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				//可能存在 "*aa*"的情况
				if j-i == 1 {
					temp[i][j] = 1
				} else {
					temp[i][j] = temp[i+1][j-1]
				}
			}
			//统计个数
			if temp[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
