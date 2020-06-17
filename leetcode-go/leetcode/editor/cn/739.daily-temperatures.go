package main

/*
[739]每日温度
//请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
//
// 例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2
//, 1, 1, 0, 0]。
//
// 提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
// Related Topics 栈 哈希表
*/

//leetcode submit region begin(Prohibit modification and deletion)
func dailyTemperatures(T []int) []int {
	length := len(T)
	//模拟栈，存储下标，确保后面的大于前面
	var array []int
	res := make([]int, length)
	for i := 0; i < length; i++ {
		//如果栈不空，且后面的小于前面，则把前面的取出来，计算差值
		for len(array) != 0 && T[array[len(array)-1]] < T[i] {
			res[array[len(array)-1]] = i - array[len(array)-1]
			array = array[:len(array)-1]
		}
		array = append(array, i)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
