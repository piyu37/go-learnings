package main

import "fmt"

type Trie struct {
	Links [26]*Trie
	Flag  bool
}

func Constructor() Trie {
	return Trie{
		Links: [26]*Trie{},
		Flag:  false,
	}
}

func containsKey(this *Trie, ch byte) bool {
	val := this.Links[ch-'a']
	return val != nil
}

func get(this *Trie, ch byte) *Trie {
	return this.Links[ch-'a']
}

func put(this *Trie, link *Trie, ch byte) {
	this.Links[ch-'a'] = link
}

func setEnd(this *Trie) {
	this.Flag = true
}

func (t *Trie) Insert(word string) {
	if t == nil {
		return
	}

	trie := t
	for i := range word {
		char := word[i]

		if !containsKey(trie, char) {
			put(trie, &Trie{}, char)
		}

		trie = get(trie, char)
	}

	setEnd(trie)
}

func (t *Trie) Search(word string) bool {
	trie := t

	for i := range word {
		ch := word[i]

		if !containsKey(trie, ch) {
			return false
		}

		trie = get(trie, ch)
	}

	return trie.Flag
}

func (t *Trie) WildSearch(word string) bool {
	trie := t

	for i := range word {
		ch := word[i]

		if ch == '.' || !containsKey(trie, ch) {
			if ch == '.' {
				for _, link := range trie.Links {
					if link != nil {
						newWord := word[i+1:]
						if link.WildSearch(newWord) {
							return true
						}
					}
				}
			}

			return false
		} else {
			trie = get(trie, ch)
		}
	}

	return trie.Flag
}

func (t *Trie) StartsWith(prefix string) bool {
	trie := t

	for i := range prefix {
		ch := prefix[i]

		if !containsKey(trie, ch) {
			return false
		}

		trie = get(trie, ch)
	}

	return true
}

// https://leetcode.com/problems/implement-trie-prefix-tree/description/?envType=study-plan-v2&envId=top-interview-150
// https://leetcode.com/problems/design-add-and-search-words-data-structure/description/?envType=study-plan-v2&envId=top-interview-150
func trie() {
	t := Constructor()

	t.Insert("apple")
	t.Insert("apps")
	t.Insert("a")
	fmt.Println(t.Search("apps"))
	fmt.Println(t.Search("abcd"))
	fmt.Println(t.Search("app"))
	fmt.Println(t.WildSearch("a."))
	fmt.Println(t.WildSearch("a.ps"))
	fmt.Println(t.WildSearch("..ps"))
	fmt.Println(t.WildSearch(".ps"))
	fmt.Println(t.StartsWith("app"))
	fmt.Println(t.StartsWith("ab"))
}
