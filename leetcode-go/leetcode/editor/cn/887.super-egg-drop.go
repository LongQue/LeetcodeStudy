package main

/*
[887]鸡蛋掉落
//你将获得 K 个鸡蛋，并可以使用一栋从 1 到 N 共有 N 层楼的建筑。
//
// 每个蛋的功能都是一样的，如果一个蛋碎了，你就不能再把它掉下去。
//
// 你知道存在楼层 F ，满足 0 <= F <= N 任何从高于 F 的楼层落下的鸡蛋都会碎，从 F 楼层或比它低的楼层落下的鸡蛋都不会破。
//
// 每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 X 扔下（满足 1 <= X <= N）。
//
// 你的目标是确切地知道 F 的值是多少。
//
// 无论 F 的初始值如何，你确定 F 的值的最小移动次数是多少？
//
//
//
//
//
//
// 示例 1：
//
// 输入：K = 1, N = 2
//输出：2
//解释：
//鸡蛋从 1 楼掉落。如果它碎了，我们肯定知道 F = 0 。
//否则，鸡蛋从 2 楼掉落。如果它碎了，我们肯定知道 F = 1 。
//如果它没碎，那么我们肯定知道 F = 2 。
//因此，在最坏的情况下我们需要移动 2 次以确定 F 是多少。
//
//
// 示例 2：
//
// 输入：K = 2, N = 6
//输出：3
//
//
// 示例 3：
//
// 输入：K = 3, N = 14
//输出：4
//
//
//
//
// 提示：
//
//
// 1 <= K <= 100
// 1 <= N <= 10000
//
// Related Topics 数学 二分查找 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
//https://leetcode-cn.com/problems/super-egg-drop/solution/gobian-ma-jie-fa-shuang-1000ms-19mb-by-ilwwli/
/**
用10位二进制可以代表1024个数，即2^10次方排列组合可以唯一映射表示1024个数 A(10,10)=C(0,10)+C(1,10)+C(2,10)+...+C(10,10)

用time次扔鸡蛋可以表示多少个数（层楼F） 0-鸡蛋没碎，1-鸡蛋碎了，限制条件 鸡蛋全碎了 or 已达到time次
K 鸡蛋数
sum += C(K,time)  K in (0,K)，如果这个映射数小于题目给出的层数，说明扔的次数不够，需要time+1，注意由于0<=F<=N的情况，所以需要F的个数=N+1
*/
func superEggDrop(K int, N int) int {
	N++
	for time := 0; ; time++ {
		max := 0
		//i鸡蛋破碎数
		for i := 0; i <= K && i <= time; i++ {
			max += comb(i, time)
		}
		if max >= N {
			return time
		}
	}
}

// 组合 C(2,5)= 5! / (2!*(5-2)!)=3*4*5 / (5-2)!=3*4*5  / 2*3 = 4*5/2
func comb(a, b int) int {
	if a > b {
		a = b
	}
	sum := 1
	// 5-2+1 +1是为了防止i=0，从4开始
	for i := b - a + 1; i <= b; i++ {
		sum *= i
	}
	// 从2开始到a
	for i := 2; i <= a; i++ {
		sum /= i
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
