package main

/*
[139]单词拆分
//给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
//
// 说明：
//
//
// 拆分时可以重复使用字典中的单词。
// 你可以假设字典中没有重复的单词。
//
//
// 示例 1：
//
// 输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
//
//
// 示例 2：
//
// 输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
//     注意你可以重复使用字典中的单词。
//
//
// 示例 3：
//
// 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false
//
// Related Topics 动态规划
*/

//leetcode submit region begin(Prohibit modification and deletion)
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 || s == "" {
		return false
	}
	//构造Set集合
	mSet := make(map[string]int)
	for _, s := range wordDict {
		mSet[s] = 1
	}
	//广度优先遍历，从0下标开始找 s的匹配
	queue := []int{0}
	//标记查找过得位置
	startFlag := make([]int, len(s))
	for len(queue) != 0 {
		//取出第一个
		start := queue[0]
		queue = queue[1:]
		//如果该位置未被查找过
		if startFlag[start] == 0 {
			for end := start + 1; end <= len(s); end++ {
				//切割字符串，判断是否在Set中，由于切割的特殊性 [start,end)不包含end，所以end<=len(s)
				if _, ok := mSet[s[start:end]]; ok {
					//如果有则把end作为下次广度查找的起点下标
					queue = append(queue, end)
					//如果刚好end达到末尾则说明全部存在
					if end == len(s) {
						return true
					}
				}
			}
		}
		startFlag[start] = 1
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
