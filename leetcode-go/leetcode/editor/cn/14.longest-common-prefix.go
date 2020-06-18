package main

/*
[14]最长公共前缀
//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
// 示例 1:
//
// 输入: ["flower","flow","flight"]
//输出: "fl"
//
//
// 示例 2:
//
// 输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//
//
// 说明:
//
// 所有输入只包含小写字母 a-z 。
// Related Topics 字符串
*/

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	first := strs[0]
	minValue := len(first)
	for i := 1; i < len(strs); i++ {
		minValue = min(longestCommonPrefixHelper(first, strs[i]), minValue)
	}
	return first[:minValue]
}
func longestCommonPrefixHelper(first, target string) int {
	count := 0
	for i := 0; i < len(first) && i < len(target); i++ {
		if first[i] == target[i] {
			count++
		} else {
			break
		}
	}
	return count
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
