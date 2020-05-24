package main

/*
[20]有效的括号
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
// 注意空字符串可被认为是有效字符串。
//
// 示例 1:
//
// 输入: "()"
//输出: true
//
//
// 示例 2:
//
// 输入: "()[]{}"
//输出: true
//
//
// 示例 3:
//
// 输入: "(]"
//输出: false
//
//
// 示例 4:
//
// 输入: "([)]"
//输出: false
//
//
// 示例 5:
//
// 输入: "{[]}"
//输出: true
// Related Topics 栈 字符串
*/

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	left := map[byte]int{
		'[': 1,
		'{': 2,
		'(': 3,
	}
	right := map[byte]int{
		']': 1,
		'}': 2,
		')': 3,
	}
	var temp []int
	for i := range s {
		if left[s[i]] != 0 {
			temp = append(temp, left[s[i]])
		} else if len(temp) != 0 && right[s[i]] == temp[len(temp)-1] {
			temp = temp[:len(temp)-1]
		} else {
			return false
		}
	}
	if len(temp) != 0 {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
