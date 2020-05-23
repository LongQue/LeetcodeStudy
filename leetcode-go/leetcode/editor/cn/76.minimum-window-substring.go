package main

/*
[76]最小覆盖子串
//给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字符的最小子串。
//
// 示例：
//
// 输入: S = "ADOBECODEBANC", T = "ABC"
//输出: "BANC"
//
// 说明：
//
//
// 如果 S 中不存这样的子串，则返回空字符串 ""。
// 如果 S 中存在这样的子串，我们保证它是唯一的答案。
//
// Related Topics 哈希表 双指针 字符串 Sliding Window
*/

//leetcode submit region begin(Prohibit modification and deletion)

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	// t是字符串，但不限定重复次数 大小写%^^&*((
	//字符长度
	minLen := len(s) + 1
	//最小字符开头
	minStart := 0
	//记录t中的字符及出现次数
	m := make(map[int32]int)
	//记录是否为t中的字符，类似于list.contains()
	flags := make([]int, 128)
	//记录t字符的种类
	stringType := 0
	//计算t字符的数据
	for _, v := range t {
		if m[v] == 0 {
			flags[v] = 1
			stringType++
		}
		m[v]++
	}
	//滑动指针 ，left为最左，i为右边扩展
	left := 0
	for i, v := range s {
		//如果是t记录的字符
		if flags[v] == 1 {
			//m次数减一，到负了也不会影响
			m[v]--
			if m[v] == 0 {
				//符合t中的次数时，字符种类减1
				stringType--
			}
			//字符种类为0，即t的字符在left i之间已全部出现
			for stringType == 0 {
				//判断字符长度最小
				if i-left+1 < minLen {
					minLen = i - left + 1
					minStart = left
				}
				//移动左指针
				temp := s[left]
				if flags[temp] == 1 {
					//更改数量
					m[int32(temp)]++
				}
				//如果有缺失则终止循环，可能遇到连续都说left对应字符的情况
				if m[int32(temp)] > 0 {
					stringType++
				}
				left++
			}
		}
	}
	if minLen > len(s) {
		return ""
	}
	return s[minStart : minStart+minLen]
}

//leetcode submit region end(Prohibit modification and deletion)
