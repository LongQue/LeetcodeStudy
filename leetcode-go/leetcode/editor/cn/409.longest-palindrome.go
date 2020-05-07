package main

/*
[409]最长回文串
//给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
//
// 在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
//
// 注意:
//假设字符串的长度不会超过 1010。
//
// 示例 1:
//
//
//输入:
//"abccccdd"
//
//输出:
//7
//
//解释:
//我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
//
// Related Topics 哈希表
*/

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) int {
	//记录每种字符出现次数
	var arrays [128]int
	//回文长度
	length := 0
	for _, v := range s {
		//出现一次
		arrays[v]++
		//说明出现0次数是双数，回文+2
		if arrays[v]&1 == 0 {
			length += 2
		}
	}
	//如果回文长度小于s长度，则说明有落单字符可以作为中心，则回文长度+1
	if length < len(s) {
		return length + 1
	}
	return length
}

//leetcode submit region end(Prohibit modification and deletion)
