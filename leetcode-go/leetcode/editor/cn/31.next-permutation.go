package main

/*
[31]下一个排列
//实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
//
// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//
// 必须原地修改，只允许使用额外常数空间。
//
// 以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
//1,2,3 → 1,3,2
//3,2,1 → 1,2,3
//1,1,5 → 1,5,1
// Related Topics 数组
*/

//leetcode submit region begin(Prohibit modification and deletion)
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i := len(nums) - 2
	//从后面寻找，招待第一个前小于后的情况，i为前的索引
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	//如果i=-1则整个数组称降序，直接做逆序交换
	if i >= 0 {
		//如果不等于-1，从后面找第一个大于nums[i]的情况，交换
		//交换后i之后的数组都是降序
		k := len(nums) - 1
		for i < k {
			if nums[k] > nums[i] {
				nums[k], nums[i] = nums[i], nums[k]
				break
			}
			k--
		}
	}
	//将i之后的数组逆序
	for a, b := i+1, len(nums)-1; a < b; a, b = a+1, b-1 {
		nums[a], nums[b] = nums[b], nums[a]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
