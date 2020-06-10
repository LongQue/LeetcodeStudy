package main

/*
[301]删除无效的括号
//删除最小数量的无效括号，使得输入的字符串有效，返回所有可能的结果。
//
// 说明: 输入可能包含了除 ( 和 ) 以外的字符。
//
// 示例 1:
//
// 输入: "()())()"
//输出: ["()()()", "(())()"]
//
//
// 示例 2:
//
// 输入: "(a)())()"
//输出: ["(a)()()", "(a())()"]
//
//
// 示例 3:
//
// 输入: ")("
//输出: [""]
// Related Topics 深度优先搜索 广度优先搜索
*/

//leetcode submit region begin(Prohibit modification and deletion)
func removeInvalidParentheses(s string) []string {
	//广度遍历，校验队列是否有正确的(),如果有拿到正确的返回，如果没有将每一种删除的可能都加入队列（go没set，用map）

	//队列
	queue := make(map[string]int)
	queue[s] = 1
	var validQueue []string
	for {
		validQueue = []string{}
		for i, _ := range queue {
			//遍历队列，校验是否有正确的
			if removeInvalidParenthesesHelper(i) {
				validQueue = append(validQueue, i)
			}
		}
		//如果有，则返回
		if len(validQueue) != 0 {
			return validQueue
		}
		//临时队列
		temp := make(map[string]int)
		//遍历原队列
		for i := range queue {
			for j := range i {
				//删除，加入临时队列
				if i[j] == '(' || i[j] == ')' {
					temp[i[:j]+i[j+1:]] = 1
				}
			}
		}
		//替换队列
		queue = temp
	}
}
func removeInvalidParenthesesHelper(s string) bool {
	count := 0
	for i := range s {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}

//leetcode submit region end(Prohibit modification and deletion)
