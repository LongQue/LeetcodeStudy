package main

/*
[974]和可被 K 整除的子数组
//给定一个整数数组 A，返回其中元素之和可被 K 整除的（连续、非空）子数组的数目。
//
//
//
// 示例：
//
// 输入：A = [4,5,0,-2,-3,1], K = 5
//输出：7
//解释：
//有 7 个子数组满足其元素之和可被 K = 5 整除：
//[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]
//
//
//
//
// 提示：
//
//
// 1 <= A.length <= 30000
// -10000 <= A[i] <= 10000
// 2 <= K <= 10000
//
// Related Topics 数组 哈希表
*/

//leetcode submit region begin(Prohibit modification and deletion)
func subarraysDivByK(A []int, K int) int {
	// m[0]=1 因为第一次整除直接得到结果，其他索引需要先++等下次在遇到相同索引的时候就知道重复了K的倍数
	//例如  5 2 1 4      K=5
	// 5 得到下标 0，ret+=map[0],之后map[0]++， 索引map[0]需要初始化1
	// 2  map[2]++,当 1 4 的时候 5+2+1+4又回到2，ret+=map[2]，如果有值，说明 5 2 1 4 - 5 2 刚好等于5的倍数，存在可以整除的连续数组 1 4
	m := map[int]int{0: 1}
	sum := 0
	ret := 0
	for _, v := range A {
		sum += v
		//避免负数的整除 -7%5=-2  (-2+5)%=3   (-7+10)%5=3
		modulus := (sum%K + K) % K
		ret += m[modulus]
		m[modulus]++
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
