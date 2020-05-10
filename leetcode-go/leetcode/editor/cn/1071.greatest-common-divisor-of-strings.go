package main

/*
[1071]字符串的最大公因子
//对于字符串 S 和 T，只有在 S = T + ... + T（T 与自身连接 1 次或多次）时，我们才认定 “T 能除尽 S”。
//
// 返回最长字符串 X，要求满足 X 能除尽 str1 且 X 能除尽 str2。
//
//
//
// 示例 1：
//
// 输入：str1 = "ABCABC", str2 = "ABC"
//输出："ABC"
//
//
// 示例 2：
//
// 输入：str1 = "ABABAB", str2 = "ABAB"
//输出："AB"
//
//
// 示例 3：
//
// 输入：str1 = "LEET", str2 = "CODE"
//输出：""
//
//
//
//
// 提示：
//
//
// 1 <= str1.length <= 1000
// 1 <= str2.length <= 1000
// str1[i] 和 str2[i] 为大写英文字母
//
// Related Topics 字符串
*/

//leetcode submit region begin(Prohibit modification and deletion)
func gcdOfStrings(str1 string, str2 string) string {
	//str1==n个str   str2==m个str，所以相加理应相等
	if str1+str2 != str2+str1 {
		return ""
	}
	//辗转相除法求最大公因数
	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if b != 0 {
			return gcd(b, a%b)
		} else {
			return a
		}
	}
	//用长度计算公因数
	res := str1[:gcd(len(str1), len(str2))]
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
