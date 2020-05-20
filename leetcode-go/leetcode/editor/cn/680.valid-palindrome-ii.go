package main

/*
[680]验证回文字符串 Ⅱ
//给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
//
// 示例 1:
//
//
//输入: "aba"
//输出: True
//
//
// 示例 2:
//
//
//输入: "abca"
//输出: True
//解释: 你可以删除c字符。
//
//
// 注意:
//
//
// 字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
//
// Related Topics 字符串
*/

//leetcode submit region begin(Prohibit modification and deletion)
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return validPalindromeHelper(s, l+1, r) || validPalindromeHelper(s, l, r-1)
		}
		l++
		r--
	}
	return true
}
func validPalindromeHelper(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
