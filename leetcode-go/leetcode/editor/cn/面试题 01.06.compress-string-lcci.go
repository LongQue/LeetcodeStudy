package main

import (
	"bytes"
	"strconv"
)

/*
[面试题 01.06]字符串压缩
//字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没
//有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。
//
// 示例1:
//
//
// 输入："aabcccccaaa"
// 输出："a2b1c5a3"
//
//
// 示例2:
//
//
// 输入："abbccd"
// 输出："abbccd"
// 解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。
//
//
// 提示：
//
//
// 字符串长度在[0, 50000]范围内。
//
// Related Topics 字符串
*/

//leetcode submit region begin(Prohibit modification and deletion)
func compressString(S string) string {
	if len(S) < 2 {
		return S
	}
	bs := []byte(S)
	count := 0
	var stringBuffer bytes.Buffer
	//遍历 n-1个
	for i := 1; i < len(S); i++ {
		//记录有多少个相同
		count++
		if S[i-1] != S[i] {
			stringBuffer.WriteByte(S[i-1])
			stringBuffer.WriteString(strconv.Itoa(count))

			count = 0
		}
	}
	stringBuffer.WriteByte(S[len(S)-1])
	//+1为了补充最后一个
	stringBuffer.WriteString(strconv.Itoa(count + 1))
	s := stringBuffer.String()
	if len(bs) <= len(s) {
		return S
	}
	return s
}

//leetcode submit region end(Prohibit modification and deletion)
