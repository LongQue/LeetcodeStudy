package main

/*
[1371]每个元音包含偶数次的最长子字符串
//给你一个字符串 s ，请你返回满足以下条件的最长子字符串的长度：每个元音字母，即 'a'，'e'，'i'，'o'，'u' ，在子字符串中都恰好出现了偶数次。
//
//
//
//
// 示例 1：
//
//
//输入：s = "eleetminicoworoep"
//输出：13
//解释：最长子字符串是 "leetminicowor" ，它包含 e，i，o 各 2 个，以及 0 个 a，u 。
//
//
// 示例 2：
//
//
//输入：s = "leetcodeisgreat"
//输出：5
//解释：最长子字符串是 "leetc" ，其中包含 2 个 e 。
//
//
// 示例 3：
//
//
//输入：s = "bcbcbc"
//输出：6
//解释：这个示例中，字符串 "bcbcbc" 本身就是最长的，因为所有的元音 a，e，i，o，u 都出现了 0 次。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 5 x 10^5
// s 只包含小写英文字母。
//
// Related Topics 字符串
*/

//leetcode submit region begin(Prohibit modification and deletion)
func findTheLongestSubstring(s string) int {
	//前缀和，如果刚好落入到一个访问过的下标，说明状态重复，出现符合条件的字符串
	m := make([]int, 32)
	for i := range m {
		//意为未访问过
		m[i] = -2
	}
	// index 0特殊处理
	m[0] = -1
	count := 0
	ret := 0
	for i := range s {
		// 5个字母共计32种状态
		switch s[i] {
		case 'a':
			count ^= 1
			break
		case 'e':
			count ^= 2
			break
		case 'i':
			count ^= 4
			break
		case 'u':
			count ^= 8
			break
		case 'o':
			count ^= 16
			break
		default:
			break
		}
		if m[count] == -2 {
			//存每个状态第一次出现的下标，第二次出现的时候就循环了
			m[count] = i
		} else {
			ret = max(ret, i-m[count])
		}
	}
	return ret
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
