package main

/*
[300]最长上升子序列
//给定一个无序的整数数组，找到其中最长上升子序列的长度。
//
// 示例:
//
// 输入: [10,9,2,5,3,7,101,18]
//输出: 4
//解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
//
// 说明:
//
//
// 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
// 你算法的时间复杂度应该为 O(n2) 。
//
//
// 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
// Related Topics 二分查找 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	//临时数组，长度即为最长子序列长度
	temp := make([]int, len(nums))
	tempLen := 0
	//遍历原数组
	for _, v := range nums {
		//二分查找，遍历临时数组
		//i是下标，j是长度
		i, j := 0, tempLen
		for i < j {
			mid := (i + j) / 2
			if temp[mid] < v {
				i = mid + 1
			} else {
				j = mid
			}
		}
		temp[i] = v
		//若j没有改动，说明i=mid+1 = j，temp[i]的位置存入数据，会使长度+1
		if tempLen == j {
			tempLen++
		}
	}
	return tempLen
}

//leetcode submit region end(Prohibit modification and deletion)
