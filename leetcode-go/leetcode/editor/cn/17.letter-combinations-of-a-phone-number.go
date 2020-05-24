package main

/*
[17]电话号码的字母组合
//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
// 示例:
//
// 输入："23"
//输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//
//
// 说明:
//尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
// Related Topics 字符串 回溯算法
*/

//leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	//老老实实三层套娃吧
	if digits == "" {
		return nil
	}
	num2String := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	var res []string
	res = append(res, "")
	for i := range digits {
		nums := num2String[digits[i]]
		pre := res
		res = []string{}
		for _, num := range nums {
			for _, v := range pre {
				res = append(res, v+num)
			}
		}
	}
	return res

}

//leetcode submit region end(Prohibit modification and deletion)
