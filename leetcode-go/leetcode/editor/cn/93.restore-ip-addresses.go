package main

import (
	"strconv"
	"strings"
)

/*
[93]复原IP地址
//给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
//
// 有效的 IP 地址正好由四个整数（每个整数位于 0 到 255 之间组成），整数之间用 '.' 分隔。
//
//
//
// 示例:
//
// 输入: "25525511135"
//输出: ["255.255.11.135", "255.255.111.35"]
// Related Topics 字符串 回溯算法
*/

//leetcode submit region begin(Prohibit modification and deletion)

func restoreIpAddresses(s string) []string {
	length := len(s)
	var res []string
	if length > 12 || length < 4 {
		return res
	}
	//记录每层的切割
	var temp []string
	restoreIpAddressesHelp(s, length, 0, 4, &temp, &res)
	return res
}
func restoreIpAddressesHelp(s string, length, begin, residue int, temp, res *[]string) {
	//最后一个已分配
	if length == begin {
		//且分了四次
		if residue == 0 {
			//串联起来
			*res = append(*res, strings.Join(*temp, "."))
		}
		return
	}

	for i := begin; i < begin+3; i++ {
		//不可超过最大长度
		if i >= length {
			break
		}
		//确保residue层可以分完
		if length-i > residue*3 {
			continue
		}
		if judgeIpSegment(s, begin, i) {
			str := s[begin : i+1]
			*temp = append(*temp, str)
			restoreIpAddressesHelp(s, length, i+1, residue-1, temp, res)
			*temp = (*temp)[:len(*temp)-1]
		}
	}
}
func judgeIpSegment(s string, left, right int) bool {
	length := right - left + 1
	//对于长度大于1的，不可以0开头
	if length > 1 && s[left] == '0' {
		return false
	}
	res := 0
	for left <= right {
		//直接以s[left]获取到的是 uint8
		i, _ := strconv.Atoi(s[left : left+1])
		res = res*10 + i
		left++
	}
	return res >= 0 && res <= 255
}

//leetcode submit region end(Prohibit modification and deletion)
