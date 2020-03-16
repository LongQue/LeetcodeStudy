package main

//1.两数之和
func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for index, num := range nums {
		if i, ok := m[target-num]; ok {
			return []int{i, index}
		}
		m[num] = index
	}
	return []int{}
}

//7. Reverse Integer
func reverse(x int) int {
	if x < 0 {
		return -reverse(-x)
	}
	var temp int
	for x != 0 {
		temp = temp*10 + x%10
		x = x / 10
	}
	if temp < 0x7fffffff {
		return temp
	}
	return 0
}

//9. Palindrome Number 回文数
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	num, temp := 0, x
	for x > 0 {
		num = num*10 + x%10
		x /= 10
	}
	return num == temp
}
