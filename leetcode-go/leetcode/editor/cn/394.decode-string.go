package main

/*
[394]字符串解码
//给定一个经过编码的字符串，返回它解码后的字符串。
//
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
//
// 示例:
//
//
//s = "3[a]2[bc]", 返回 "aaabcbc".
//s = "3[a2[c]]", 返回 "accaccacc".
//s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
//
// Related Topics 栈 深度优先搜索
*/

//leetcode submit region begin(Prohibit modification and deletion)
func decodeString(s string) string {
	//存字母
	alphaStack := make([]byte, 0)
	//存重复次数
	countStack := make([]int, 0)
	//存[下一个开始的下标 指alphaStack下的
	alphaIndex := make([]int, 0)
	var count int
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			// 可能出现 两位数
			count = count*10 + int(s[i]-'0')
		} else if s[i] == '[' { // [ 后面肯定是字母
			// 可以把当前的数字放入次数数组
			countStack = append(countStack, count)
			count = 0

			alphaIndex = append(alphaIndex, len(alphaStack))
		} else if s[i] == ']' { // 对当前字母段出站
			// 获取当前字母段的次数 然后次数出栈
			c := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			// 获取当前字母段开始的索引 然后索引出栈
			index := alphaIndex[len(alphaIndex)-1]
			alphaIndex = alphaIndex[:len(alphaIndex)-1]

			//获取当前的字母段
			str := string(alphaStack[index:])
			// 字母数组出站
			alphaStack = alphaStack[:index]

			for j := 0; j < c; j++ {
				alphaStack = append(alphaStack, []byte(str)...)
			}
		} else {
			alphaStack = append(alphaStack, s[i])
		}
	}
	return string(alphaStack)
}

//leetcode submit region end(Prohibit modification and deletion)
