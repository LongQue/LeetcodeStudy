package main

/*
[4]寻找两个正序数组的中位数
//给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
//
// 请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
// 你可以假设 nums1 和 nums2 不会同时为空。
//
//
//
// 示例 1:
//
// nums1 = [1, 3]
//nums2 = [2]
//
//则中位数是 2.0
//
//
// 示例 2:
//
// nums1 = [1, 2]
//nums2 = [3, 4]
//
//则中位数是 (2 + 3)/2 = 2.5
//
// Related Topics 数组 二分查找 分治算法
*/

//leetcode submit region begin(Prohibit modification and deletion)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//如果总长度为7，则需要第4位，  7+1/2  7+2/2 都是4
	//如果总长度为8，需要第4,5位   8+1/2  8+2/2  是 4和5
	left := (len(nums1) + len(nums2) + 1) / 2
	right := (len(nums1) + len(nums2) + 2) / 2
	// 为单数 计算两次4， 为双数计算4和5
	return float64(getMid(nums1, nums2, left)+getMid(nums1, nums2, right)) / 2
}

func getMid(nums1, nums2 []int, k int) int {
	n := len(nums1)
	m := len(nums2)
	//确保nums1长度小于nums2
	if n > m {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}
	//处理num1为空的情况
	if n == 0 {
		return nums2[k-1]
	}
	//找到第k个数
	if k == 1 {
		return min(nums1[0], nums2[0])
	}
	//每次删除k/2个数， min为了处理数组长度小于k/2的情况
	i := min(n, k/2) - 1
	j := min(m, k/2) - 1
	// [1 2 3 4]
	// [1 2 3 4 5 6 7 8 9 10]
	// 总长度为7，left=8 right=9 k/2=4，  -1是为了从长度->索引
	// i=2 j=2，比较这两个位置大小就可以知道，应先删除小的那组，大的那组可能有超过中位的
	if nums1[i] > nums2[j] {
		//删nums2前部分，   k-j-1是计算差多少到中位，-1是从索引变成长度
		return getMid(nums1, nums2[j+1:], k-j-1)
	} else {
		return getMid(nums1[i+1:], nums2, k-i-1)
	}
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
