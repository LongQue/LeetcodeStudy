package main

import "strings"

/*
[67]二进制求和
//给你两个二进制字符串，返回它们的和（用二进制表示）。
//
// 输入为 非空 字符串且只包含数字 1 和 0。
//
//
//
// 示例 1:
//
// 输入: a = "11", b = "1"
//输出: "100"
//
// 示例 2:
//
// 输入: a = "1010", b = "1011"
//输出: "10101"
//
//
//
// 提示：
//
//
// 每个字符串仅由字符 '0' 或 '1' 组成。
// 1 <= a.length, b.length <= 10^4
// 字符串如果不是 "0" ，就都不含前导零。
//
// Related Topics 数学 字符串
// 👍 431 👎 0
*/

//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	var result []string
	//进位标志
	flag := 0
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		temp := flag
		if i >= 0 {
			temp += int(a[i] - '0')
		}
		if j >= 0 {
			temp += int(b[j] - '0')
		}
		flag = temp / 2

		result = append([]string{string(temp%2 + '0')}, result...)
	}
	if flag == 1 {
		result = append([]string{string(flag + '0')}, result...)
	}
	return strings.Join(result, "")
}

//leetcode submit region end(Prohibit modification and deletion)
