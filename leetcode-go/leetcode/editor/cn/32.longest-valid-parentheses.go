package main

/*
[32]最长有效括号
//给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
//
// 示例 1:
//
// 输入: "(()"
//输出: 2
//解释: 最长有效括号子串为 "()"
//
//
// 示例 2:
//
// 输入: ")()())"
//输出: 4
//解释: 最长有效括号子串为 "()()"
//
// Related Topics 字符串 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {
	maxValue := 0
	//动态规划，因为到)才计算长度，所以(都为0
	dp := make([]int, len(s))
	//从1开始遍历字符串
	for i := 1; i < len(s); i++ {
		//只看) 分两种情况 第一种() ,第二种))
		if s[i] == ')' {
			// ()时增加2个长度，此时需要加上()前面的长度即dp[i-2]，注意越界
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
				// ))增加的长度为 第二个）的长度等于第一个）的长度，还要算上与第二个）对应的（长度即2，如不是（则跳过，最后还要看第二个（对应的）前面的长度
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				dp[i] = dp[i-1] + 2
				if i-dp[i-1]-2 >= 0 {
					dp[i] += dp[i-dp[i-1]-2]
				}
			}
		}
		maxValue = max(maxValue, dp[i])
	}
	return maxValue
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
