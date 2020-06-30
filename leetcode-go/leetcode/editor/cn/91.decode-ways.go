package main

/*
[91]解码方法
//一条包含字母 A-Z 的消息通过以下方式进行了编码：
//
// 'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
//
//
// 给定一个只包含数字的非空字符串，请计算解码方法的总数。
//
// 示例 1:
//
// 输入: "12"
//输出: 2
//解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
//
//
// 示例 2:
//
// 输入: "226"
//输出: 3
//解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
//
// Related Topics 字符串 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func numDecodings(s string) int {
	length := len(s)
	array := make([]int, length+1)
	array[0] = 1
	if length == 0 || s[0] == '0' {
		return 0
	}
	array[1] = 1
	//由于已经得到了array[1]，array和s的下标能同步，所以i从2开始，通过i<=length,s[i-2],s[i-1]来同步，
	//动态规划
	//每次新加入s[i-1]会出现
	//1.s[i-2],s[i-1]都是0则无法加入，返回0
	//2.s[i-2]==0,s[i-1]不为0,则array[i]=array[i-1]
	//3.s[i-2]!=0,s[i-1]==0  :
	//	3.1 temp>26,由于0无法单带所以返回0
	//	3.2  temp<=26,说明s[i-2]=1/2 可以加入0  array[i]= array[i-2]
	//4. temp大于26，单带，值不变；小于等于26 ，等于前两个值相加
	for i := 2; i <= length; i++ {
		temp := (s[i-2]-'0')*10 + (s[i-1] - '0')
		if temp == 0 {
			return 0
		} else if s[i-2] == '0' {
			array[i] = array[i-1]
		} else if s[i-1] == '0' {
			if temp > 26 {
				return 0
			}
			array[i] = array[i-2]
		} else if temp > 26 {
			array[i] = array[i-1]
		} else {
			array[i] = array[i-1] + array[i-2]
		}
	}
	return array[len(array)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
