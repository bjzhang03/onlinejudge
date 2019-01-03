package main

import "fmt"

type Trie struct {
	Value uint8
	Next  []*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return *(&Trie{0, []*Trie{}})
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	tmp := this
	for i := 0; i < len(word); i++ {
		// 查找元素，如果可以找到的话则直接添加
		flag := true
		for _, item := range tmp.Next {
			if item.Value == word[i] {
				flag = false
				tmp = item
				break
			}
		}
		// 找不到元素的话，则直接新建一个
		if flag {
			tmpitem := &Trie{word[i], []*Trie{}}
			tmp.Next = append(tmp.Next, tmpitem)
			tmp = tmpitem
		}
	}
	// 塞一个结束符号进去
	tmp.Next = append(tmp.Next, &Trie{'.', []*Trie{}})
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	tmp := this
	for i := 0; i < len(word); i++ {
		flag := false
		for _, item := range tmp.Next {
			if item.Value == word[i] {
				tmp = item
				flag = true
			}
		}
		if !flag {
			return false
		}
	}
	fmt.Println("value = ", tmp.Value, " ,next = ", tmp.Next)

	for _, item := range tmp.Next {
		if item.Value == '.' {
			return true
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	tmp := this
	for i := 0; i < len(prefix); i++ {
		flag := false
		for _, item := range tmp.Next {
			if item.Value == prefix[i] {
				tmp = item
				flag = true
			}
		}
		if !flag {
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
