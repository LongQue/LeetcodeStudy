package main

import (
	"sort"
	"strconv"
	"strings"
)

/*
[179]最大数
//给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。
//
// 示例 1:
//
// 输入: [10,2]
//输出: 210
//
// 示例 2:
//
// 输入: [3,30,34,5,9]
//输出: 9534330
//
// 说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
// Related Topics 排序
// 👍 317 👎 0
*/

//leetcode submit region begin(Prohibit modification and deletion)
func largestNumber(nums []int) string {
	var temp []string
	for _, v := range nums {
		temp = append(temp, strconv.Itoa(v))
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i]+temp[j] > temp[j]+temp[i]
	})
	result := strings.Join(temp, "")
	if result[0] == '0' {
		return "0"
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
