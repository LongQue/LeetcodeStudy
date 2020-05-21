package main

/*
[5]最长回文子串
//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
// 示例 1：
//
// 输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//
//
// 示例 2：
//
// 输入: "cbbd"
//输出: "bb"
//
// Related Topics 字符串 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	//中心扩展法，选定一个中心向两边扩展
	if len(s) <= 0 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		//中心单独一个
		len1 := expandAroundCenter(s, i, i)
		//中心两个相同
		len2 := expandAroundCenter(s, i, i+1)
		slen := max(len1, len2)
		if slen > end-start {
			start = i - (slen-1)/2
			end = i + slen/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	l, r := left, right
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
