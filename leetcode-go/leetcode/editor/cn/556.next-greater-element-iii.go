package main

import (
	"sort"
	"strconv"
	"strings"
)

/*
[556]下一个更大元素 III
//给定一个32位正整数 n，你需要找到最小的32位整数，其与 n 中存在的位数完全相同，并且其值大于n。如果不存在这样的32位整数，则返回-1。
//
// 示例 1:
//
//
//输入: 12
//输出: 21
//
//
// 示例 2:
//
//
//输入: 21
//输出: -1
//
// Related Topics 字符串
// 👍 93 👎 0
*/

//leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElement(n int) int {
	//从后面开始找，找出第一个打破递增规律的数得到位置a
	//再找一遍，找一个比他大一点的数得到位置b
	//交换位置ab，在从小到大排序a之后的所有数
	if n < 10 {
		return -1
	}
	//转string，再分割
	str := strings.Split(strconv.Itoa(n), "")
	i := len(str) - 2
	//找到位置a
	for i >= 0 && str[i+1] <= str[i] {
		i--
	}
	//说明从后面数起，完全递增
	if i < 0 {
		return -1
	}
	j := len(str) - 1
	//找到位置b
	for j > i && str[j] <= str[i] {
		j--
	}
	//交换
	str[j], str[i] = str[i], str[j]
	//不确定Strings是否会new一个新的改变原有
	temp := str[i+1:]
	//默认从小到大拍
	sort.Strings(temp)
	//拼接
	str2 := append(str[0:i+1], temp...)
	//合并成string，转int(go int可能自适应成int64)
	atoi, _ := strconv.Atoi(strings.Join(str2, ""))
	//判断不可大于32位
	if atoi > (1 << 31) {
		return -1
	}
	return atoi
}

//leetcode submit region end(Prohibit modification and deletion)
