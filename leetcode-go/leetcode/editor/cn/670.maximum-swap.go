package main

import (
	"strconv"
)

/*
[670]最大交换
//给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。
//
// 示例 1 :
//
//
//输入: 2736
//输出: 7236
//解释: 交换数字2和数字7。
//
//
// 示例 2 :
//
//
//输入: 9973
//输出: 9973
//解释: 不需要交换。
//
//
// 注意:
//
//
// 给定数字的范围是 [0, 108]
//
// Related Topics 数组 数学
// 👍 85 👎 0
*/

//leetcode submit region begin(Prohibit modification and deletion)
func maximumSwap(num int) int {
	//思路是计数排序
	if num <= 11 {
		return num
	}

	str := []byte(strconv.Itoa(num))
	//统计9到0 每个数字有多少，然后按从大到小排序
	//比如 99788，9有2个，8有2个，7一个，排序后应该为99887
	//在计数排序的同时，比较str中的数是否和排序的一样，不一样就证明因该将str的该位交换掉，
	//例如99788中 和99887中，第一个不一样的为第2位7，因此要将7和str中最后一个8交换
	counts := make([]byte, 10)
	for i := 0; i < len(str); i++ {
		counts[str[i]-'0']++
	}

	breakFlag := false
	i := 0
	for j := 9; j >= 0; j-- {
		if breakFlag {
			break
		}
		for ; i < len(str); i++ {
			if counts[j] == 0 {
				break
			}

			if int(str[i]-'0') == j {
				counts[j]--
				continue
			}

			//发现第一个没按顺序排序的，找到最后一个等于j的位置，并和i位置的元素交换
			index := i
			for k := i + 1; k < len(str); k++ {
				if int(str[k]-'0') == j {
					index = k
				}
			}

			//swap i,max
			tmp := str[i]
			str[i] = str[index]
			str[index] = tmp

			breakFlag = true
			break
		}
	}

	res, _ := strconv.Atoi(string(str))
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
