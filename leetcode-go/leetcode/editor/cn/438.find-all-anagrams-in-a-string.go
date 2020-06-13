package main

/*
[438]找到字符串中所有字母异位词
//给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
//
// 字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
//
// 说明：
//
//
// 字母异位词指字母相同，但排列不同的字符串。
// 不考虑答案输出的顺序。
//
//
// 示例 1:
//
//
//输入:
//s: "cbaebabacd" p: "abc"
//
//输出:
//[0, 6]
//
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
//
//
// 示例 2:
//
//
//输入:
//s: "abab" p: "ab"
//
//输出:
//[0, 1, 2]
//
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
//
// Related Topics 哈希表
*/

//leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	//记录需要的字符及其次数
	need := make(map[uint8]int)
	//记录滑动窗口内的字符及其次数
	windows := make(map[uint8]int)
	//统计
	for i := range p {
		u := p[i] - 'a'
		need[u]++
	}
	left := 0
	right := 0
	valid := 0
	var res []int
	for right < len(s) {
		//获取右指针对应的字符
		bt := s[right] - 'a'
		right++
		//如果该值是需要的
		if need[bt] != 0 {
			//窗口map记录
			windows[bt]++
			//如果窗口等于需要的
			if windows[bt] == need[bt] {
				//计数
				valid++
			}
		}
		//去报长度足够
		for right-left >= len(p) {
			//成功的计数等于不同字符出现的次数
			if valid == len(need) {
				//存入起始索引
				res = append(res, left)
			}
			//获取起始位置
			d := s[left] - 'a'
			//到达下一个
			left++
			//如果d是p中的字符
			if need[d] != 0 {
				//且相等，不相等就不需要减valid
				if need[d] == windows[d] {
					valid--
				}
				//移动之后 windows当然减一
				windows[d]--
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
