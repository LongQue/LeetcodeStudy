package main

import "strings"

/*
[402]移掉K位数字
//给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。
//
// 注意:
//
//
// num 的长度小于 10002 且 ≥ k。
// num 不会包含任何前导零。
//
//
// 示例 1 :
//
//
//输入: num = "1432219", k = 3
//输出: "1219"
//解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
//
//
// 示例 2 :
//
//
//输入: num = "10200", k = 1
//输出: "200"
//解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
//
//
// 示例 3 :
//
//
//输入: num = "10", k = 2
//输出: "0"
//解释: 从原数字移除所有的数字，剩余为空就是0。
//
// Related Topics 栈 贪心算法
// 👍 250 👎 0
*/

//leetcode submit region begin(Prohibit modification and deletion)
func removeKdigits(num string, k int) string {
	//1432219    1a2219   1b2219   a>b时 1a2219>1b2219则，确保前面部分递增，直到移除次数达到k。又或者全体递增，k剩余多少，最后一次从末尾移除k
	if len(num) <= k {
		return "0"
	}
	//记录结果
	var stack []byte
	for i := range num {
		c := num[i]
		//k>0且stack不为空且末尾大于c，则需要移除末尾
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > c {
			stack = stack[:len(stack)-1]

			k--
		}
		//每个都加入
		stack = append(stack, c)
	}
	//预防k>0的情况
	stack = stack[:len(stack)-k]
	s := string(stack)
	//删除开头为0
	s = strings.TrimLeft(s, "0")
	if s == "" {
		return "0"
	}
	return s
}

//leetcode submit region end(Prohibit modification and deletion)
