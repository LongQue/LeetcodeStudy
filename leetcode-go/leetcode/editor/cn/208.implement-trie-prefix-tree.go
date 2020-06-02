package main

/*
[208]实现 Trie (前缀树)
//实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
//
// 示例:
//
// Trie trie = new Trie();
//
//trie.insert("apple");
//trie.search("apple");   // 返回 true
//trie.search("app");     // 返回 false
//trie.startsWith("app"); // 返回 true
//trie.insert("app");
//trie.search("app");     // 返回 true
//
// 说明:
//
//
// 你可以假设所有的输入都是由小写字母 a-z 构成的。
// 保证所有输入均为非空字符串。
//
// Related Topics 设计 字典树
*/

//leetcode submit region begin(Prohibit modification and deletion)
type Trie struct {
	B   byte
	End bool
	T   map[byte]*Trie
}

//可尝试用数组解决
//type Trie struct {
//
//	End bool
//	T [26]*Trie
//}
/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		B:   'a',
		End: false,
		T:   make(map[byte]*Trie, 0),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	temp := this
	for i := range word {
		//若存在则跳过插入
		if _, ok := temp.T[word[i]]; !ok {
			temp.T[word[i]] = &Trie{
				B:   word[i],
				End: false,
				T:   make(map[byte]*Trie, 0),
			}
		}
		temp = temp.T[word[i]]
	}
	//最后一次插入 需要标记位记录结尾
	temp.End = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	temp := this
	for i := range word {
		//ok 防止为nil
		value, ok := temp.T[word[i]]
		if ok && value.B == word[i] {
			temp = value
		} else {
			return false
		}
	}
	return temp.End
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	temp := this
	for i := range prefix {
		//ok防止为kong
		value, ok := temp.T[prefix[i]]
		if ok && value.B == prefix[i] {
			temp = value
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)
