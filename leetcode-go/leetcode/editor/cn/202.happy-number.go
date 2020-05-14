package main

/*
[202]快乐数
//编写一个算法来判断一个数 n 是不是快乐数。
//
// 「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
//如果 可以变为 1，那么这个数就是快乐数。
//
// 如果 n 是快乐数就返回 True ；不是，则返回 False 。
//
//
//
// 示例：
//
// 输入：19
//输出：true
//解释：
//12 + 92 = 82
//82 + 22 = 68
//62 + 82 = 100
//12 + 02 + 02 = 1
//
// Related Topics 哈希表 数学
*/

//leetcode submit region begin(Prohibit modification and deletion)
var arrays []int

func isHappy(n int) bool {
	arrays = []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}
	//快慢指针找环，快慢指针重复时，判断fast==1，如果true，1是唯一重复的值，如果false代表存在环，无限循环
	//不必规定快慢指针直接走多少步
	slow, fast := n, bitSquareSum(n)
	for fast != 1 && slow != fast {
		slow = bitSquareSum(slow)
		fast = bitSquareSum(bitSquareSum(fast))
	}
	return fast == 1
}
func bitSquareSum(n int) int {
	sum := 0
	for n > 0 {
		bit := n % 10
		sum += arrays[bit]
		n = n / 10
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
