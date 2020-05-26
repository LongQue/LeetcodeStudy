package main

import "sort"

/*
[49]字母异位词分组
//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//
// 说明：
//
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。
//
// Related Topics 哈希表 字符串
*/

//leetcode submit region begin(Prohibit modification and deletion)
type bytes []byte

func (b bytes) Len() int {
	return len(b)
}
func (b bytes) Less(i, j int) bool {
	return b[i] < b[j]
}
func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func groupAnagrams(strs []string) [][]string {
	var ret [][]string
	m := make(map[string]int)
	for _, v := range strs {
		by := bytes(v)
		sort.Sort(by)
		s := string(by)
		if idx, ok := m[s]; !ok {
			m[s] = len(ret)
			ret = append(ret, []string{v})
		} else {
			ret[idx] = append(ret[idx], v)
		}

	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
