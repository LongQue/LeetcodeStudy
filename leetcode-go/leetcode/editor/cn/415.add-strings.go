package main

/*
[415]字符串相加
//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
//
// 注意：
//
//
// num1 和num2 的长度都小于 5100.
// num1 和num2 都只包含数字 0-9.
// num1 和num2 都不包含任何前导零。
// 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
//
// Related Topics 字符串
*/

//leetcode submit region begin(Prohibit modification and deletion)
func addStrings(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	//保证len1>=len2
	if len1 < len2 {
		len1, len2 = len2, len1
		num1, num2 = num2, num1
	}
	ans := make([]byte, len1+1)

	//进位 默认0
	var ten byte

	for i := len1 - 1; i >= 0; i-- {
		var temp byte
		j := i - (len1 - len2)
		if j >= 0 {
			ten, temp = byteAdd(num1[i]+ten, num2[j])
		} else {
			ten, temp = byteAdd(num1[i]+ten, '0')
		}
		ans[i+1] = temp
	}
	//可能长度一样，加完最后一位还进位
	if ten > 0 {
		ans[0] = '1'
		return string(ans)
	} else {
		return string(ans[1:])
	}

}

//两个byte相加
//返回值 进位值 int (1 or 0)
//acsii byte
func byteAdd(a, b byte) (byte, byte) {
	ans := a + b - '0' - '0'
	if ans >= 10 {
		return 1, ans + '0' - 10
	}
	return 0, ans + '0'
}

//leetcode submit region end(Prohibit modification and deletion)
