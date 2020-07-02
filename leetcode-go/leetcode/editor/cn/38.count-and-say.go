package main

import (
	"strconv"
	"strings"
)

/*
[38]外观数列
//给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。
//
// 注意：整数序列中的每一项将表示为一个字符串。
//
// 「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：
//
// 1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//
//
// 第一项是数字 1
//
// 描述前一项，这个数是 1 即 “一个 1 ”，记作 11
//
// 描述前一项，这个数是 11 即 “两个 1 ” ，记作 21
//
// 描述前一项，这个数是 21 即 “一个 2 一个 1 ” ，记作 1211
//
// 描述前一项，这个数是 1211 即 “一个 1 一个 2 两个 1 ” ，记作 111221
//
//
//
// 示例 1:
//
// 输入: 1
//输出: "1"
//解释：这是一个基本样例。
//
// 示例 2:
//
// 输入: 4
//输出: "1211"
//解释：当 n = 3 时，序列是 "21"，其中我们有 "2" 和 "1" 两组，"2" 可以读作 "12"，也就是出现频次 = 1 而 值 = 2；类似
//"1" 可以读作 "11"。所以答案是 "12" 和 "11" 组合在一起，也就是 "1211"。
// Related Topics 字符串
*/

//leetcode submit region begin(Prohibit modification and deletion)
func countAndSay(n int) string {
	//初始化 第一层时的string
	str := "1"
	//从第2层开始，直到n层
	for i := 2; i <= n; i++ {
		var temp strings.Builder
		//获取第一个
		preByte := str[0]
		//计算重复次数，最少1次
		count := 1
		//遍历str
		for j := 1; j < len(str); j++ {
			//相同，重复数+1
			if str[j] == preByte {
				count++
			} else {
				//不同时写入string，并把count置为1
				//写入string
				temp.WriteString(strconv.Itoa(count))
				//写入byte
				temp.WriteByte(preByte)
				//更新preByte为下一个不同的字符
				preByte = str[j]
				count = 1
			}
		}
		//最后一个无法比较，需要在for外面拼接
		temp.WriteString(strconv.Itoa(count))
		temp.WriteByte(preByte)
		//更新str
		str = temp.String()
	}
	return str
}

//leetcode submit region end(Prohibit modification and deletion)
