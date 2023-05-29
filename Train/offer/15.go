package offer

// 实现前缀树
type Trie struct {
	ch    [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	p := this
	for _, v := range word {
		if p.ch[v-'a'] == nil {
			p.ch[v-'a'] = &Trie{}
		}
		p = p.ch[v-'a']
	}
	p.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	p := this
	for _, v := range word {
		if p.ch[v-'a'] == nil {
			return false
		}
		p = p.ch[v-'a']
	}
	return p.isEnd == true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	p := this
	for _, v := range prefix {
		if p.ch[v-'a'] == nil {
			return false
		}
		p = p.ch[v-'a']
	}
	return true
}
