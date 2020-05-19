package main

/*
[466]统计重复个数
//由 n 个连接的字符串 s 组成字符串 S，记作 S = [s,n]。例如，["abc",3]=“abcabcabc”。
//
// 如果我们可以从 s2 中删除某些字符使其变为 s1，则称字符串 s1 可以从字符串 s2 获得。例如，根据定义，"abc" 可以从 “abdbec” 获得
//，但不能从 “acbbe” 获得。
//
// 现在给你两个非空字符串 s1 和 s2（每个最多 100 个字符长）和两个整数 0 ≤ n1 ≤ 106 和 1 ≤ n2 ≤ 106。现在考虑字符串 S
//1 和 S2，其中 S1=[s1,n1] 、S2=[s2,n2] 。
//
// 请你找出一个可以满足使[S2,M] 从 S1 获得的最大整数 M 。
//
//
//
// 示例：
//
// 输入：
//s1 ="acb",n1 = 4
//s2 ="ab",n2 = 2
//
//返回：
//2
//
// Related Topics 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	len1, len2 := len(s1), len(s2)
	index1, index2 := 0, 0 // 注意此处直接使用 Ra Rb 的下标，不取模

	if len1 == 0 || len2 == 0 || len1*n1 < len2*n2 {
		return 0
	}

	map1, map2 := make(map[int]int), make(map[int]int)
	ans := 0 // 注意，此处存储的是 Ra 中 Sb 的个数，而非 Ra 中 Rb 的个数

	for index1/len1 < n1 { // 遍历整个 Ra
		if index1%len1 == len1-1 { //在 Sa 末尾
			if val, ok := map1[index2%len2]; ok { // 出现了循环，进行快进
				cycleLen := index1/len1 - val/len1                 // 每个循环占多少个 Sa
				cycleNum := (n1 - 1 - index1/len1) / cycleLen      // 还有多少个循环
				cycleS2Num := index2/len2 - map2[index2%len2]/len2 // 每个循环含有多少个 Sb

				index1 += cycleNum * cycleLen * len1 // 将 Ra 快进到相应的位置
				ans += cycleNum * cycleS2Num         // 把快进部分的答案数量加上
			} else { // 第一次，注意存储的是未取模的
				map1[index2%len2] = index1
				map2[index2%len2] = index2
			}

		}

		if s1[index1%len1] == s2[index2%len2] {
			if index2%len2 == len2-1 {
				ans += 1
			}
			index2 += 1
		}
		index1 += 1
	}
	return ans / n2
}

//leetcode submit region end(Prohibit modification and deletion)
