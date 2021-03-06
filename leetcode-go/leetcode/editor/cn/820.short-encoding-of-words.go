package main

import (
	"sort"
	"strings"
)

/*
[820]单词的压缩编码
//给定一个单词列表，我们将这个列表编码成一个索引字符串 S 与一个索引列表 A。
//
// 例如，如果这个列表是 ["time", "me", "bell"]，我们就可以将其表示为 S = "time#bell#" 和 indexes = [0,
// 2, 5]。
//
// 对于每一个索引，我们可以通过从字符串 S 中索引的位置开始读取字符串，直到 "#" 结束，来恢复我们之前的单词列表。
//
// 那么成功对给定单词列表进行编码的最小字符串长度是多少呢？
//
//
//
// 示例：
//
// 输入: words = ["time", "me", "bell"]
//输出: 10
//说明: S = "time#bell#" ， indexes = [0, 2, 5] 。
//
//
//
//
// 提示：
//
//
// 1 <= words.length <= 2000
// 1 <= words[i].length <= 7
// 每个单词都是小写字母 。
//
//
*/

//leetcode submit region begin(Prohibit modification and deletion)
func minimumLengthEncoding(words []string) int {
	if words == nil || len(words) == 0 {
		return 0
	}
	//1.翻转string
	for i := range words {
		words[i] = reverse(words[i])
	}
	//2.排序
	sort.Strings(words)
	//3.防越界
	words = append(words, "*")
	resLen := 0
	for i := 0; i < len(words)-1; i++ {
		// a前缀含有b
		if strings.HasPrefix(words[i+1], words[i]) {
			continue
		}
		resLen += len(words[i]) + 1

	}
	return resLen
}
func reverse(str string) string {
	bytes := []byte(str)
	l, r := 0, len(bytes)-1
	for l < r {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}
	return string(bytes)
}

//leetcode submit region end(Prohibit modification and deletion)
