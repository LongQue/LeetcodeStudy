package main

import "math"

/*
[6]Z 字形变换
//将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
//
// 比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
//
// L   C   I   R
//E T O E S I I G
//E   D   H   N
//
//
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
//
// 请你实现这个将字符串进行指定行数变换的函数：
//
// string convert(string s, int numRows);
//
// 示例 1:
//
// 输入: s = "LEETCODEISHIRING", numRows = 3
//输出: "LCIRETOESIIGEDHN"
//
//
// 示例 2:
//
// 输入: s = "LEETCODEISHIRING", numRows = 4
//输出: "LDREOEIIECIHNTSG"
//解释:
//
//L     D     R
//E   O E   I I
//E C   I H   N
//T     S     G
// Related Topics 字符串
*/

//leetcode submit region begin(Prohibit modification and deletion)
func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}
	//一组最多有多少个
	groupLen := numRows*2 - 2
	//有多少组
	groupNum := int(math.Ceil(float64(len(s)) / float64(groupLen)))
	var ansString []byte

	for i := 0; i < numRows; i++ {
		//计算第 i 行字符串
		for j := 0; j < groupNum; j++ {
			//计算第 j 组字符串
			charIndex := groupLen*j + i
			if charIndex >= len(s) {
				continue
			}
			ansString = append(ansString, s[charIndex])
			//非上下两边的情况
			/*
				0   6
				1 4 7 11
				2 5 8 10
				3   9
			*/
			if i != 0 && i != numRows-1 {
				//groupLen*(j+1)整组的数目，如果j=i，则到第二组时最多12
				//如果当前是(1,1)是7，则 j=1 groupLen*(j+1)=12,12-1=11
				charIndex = groupLen*(j+1) - i
				if charIndex < len(s) {
					ansString = append(ansString, s[charIndex])
				}
			}
		}

	}
	return string(ansString)
}

//leetcode submit region end(Prohibit modification and deletion)
