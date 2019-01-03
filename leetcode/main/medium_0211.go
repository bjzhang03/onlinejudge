package main

import "fmt"

// 未解决
type WordDictionary struct {
	Value uint8
	Next  []*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return *(&WordDictionary{0, []*WordDictionary{}})
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	tmp := this
	for i := 0; i < len(word); i++ {
		// 查找元素
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
			tmpitem := &WordDictionary{word[i], []*WordDictionary{}}
			tmp.Next = append(tmp.Next, tmpitem)
			tmp = tmpitem
		}
	}
	// 塞一个结束符号进去
	tmp.Next = append(tmp.Next, &WordDictionary{'#', []*WordDictionary{}})
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	tmp := this
	for i := 0; i < len(word); i++ {
		if word[i] == '.' {
			if i == len(word)-1 {
				break
			} else {
				result := false
				for _, item := range tmp.Next {
					result = result || item.Search(word[i+1:])
				}
			}

		} else {
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
	}
	fmt.Println("value = ", tmp.Value, " ,next = ", tmp.Next)
	for _, item := range tmp.Next {
		if item.Value == '#' {
			return true
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
