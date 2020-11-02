// https://leetcode-cn.com/problems/implement-trie-prefix-tree/

package main

type Trie struct {
	next  [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		idx := v - 'a'
		if this.next[idx] == nil {
			this.next[idx] = new(Trie)
		}
		this = this.next[idx]
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		idx := v - 'a'
		if this.next[idx] == nil {
			return false
		}
		this = this.next[idx]
	}
	return this.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		idx := v - 'a'
		if this.next[idx] == nil {
			return false
		}
		this = this.next[idx]
	}
	return true
}

func main() {

}
