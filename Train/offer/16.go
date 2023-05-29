package offer

import "strings"

// 单词替换
func replaceWords(dictionary []string, sentence string) string {
	trie := InitTrie()
	for _, v := range dictionary {
		trie.Insert(v)
	}

	sent := strings.Split(sentence, " ")
	for i, word := range sent {
		index := trie.Search(word)
		if index == -1 {
			continue
		}

		sent[i] = string(word)[:index]
	}
	return strings.Join(sent, " ")
}

type TrieNode struct {
	ch    [26]*TrieNode
	isEnd bool
}

func InitTrie() TrieNode {
	return TrieNode{}
}

func (trie *TrieNode) Insert(word string) {
	p := trie
	for _, v := range word {
		if p.ch[v-'a'] == nil {
			p.ch[v-'a'] = &TrieNode{}
		}
		p = p.ch[v-'a']
	}
	p.isEnd = true
}

// Search 返回值int，代表走到哪里是END
func (trie *TrieNode) Search(word string) int {
	p := trie
	for i, v := range word {
		if p.isEnd {
			return i
		}
		if p.ch[v-'a'] == nil {
			return -1
		}
		p = p.ch[v-'a']
	}
	return -1
}
