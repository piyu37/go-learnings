package main

import "fmt"

// Do not edit the class below except for the insertKeyValuePair,
// getValueFromKey, and getMostRecentKey methods. Feel free
// to add new properties and methods to the class.
type Node struct {
	key        string
	value      int
	prev, next *Node
}

type LRUCache struct {
	maxSize    int
	head, tail *Node
	llMap      map[string]*Node
	// Add fields here.
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		maxSize: size,
	}
}

func (cache *LRUCache) deleteNode(node *Node, key string) {
	if cache.head == cache.tail {
		delete(cache.llMap, key)
		cache.head = nil
		cache.tail = nil
		node = nil
		return
	}

	if node == cache.head {
		node.next.prev = nil
		cache.head = node.next
		delete(cache.llMap, key)
		node = nil
		return
	}

	if node == cache.tail {
		cache.tail = node.prev
		node.prev.next = nil
		delete(cache.llMap, key)
		node = nil
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	delete(cache.llMap, key)
	node = nil
}

func (cache *LRUCache) InsertNode(key string, value int) {
	if len(cache.llMap) == 0 {
		node := &Node{
			key:   key,
			value: value,
		}

		cache.llMap = map[string]*Node{}

		cache.llMap[key] = node
		cache.head = node
		cache.tail = node

		return
	}

	node := &Node{
		key:   key,
		value: value,
	}

	node.next = cache.head
	cache.head.prev = node
	cache.head = node
	cache.llMap[key] = node
}

func (cache *LRUCache) InsertKeyValuePair(key string, value int) {
	if n, exist := cache.llMap[key]; exist {
		// delete node and add at start
		// update map
		cache.deleteNode(n, key)
		cache.InsertNode(key, value)

	} else {
		if len(cache.llMap) == cache.maxSize {
			// delete node and add at start
			// update map
			cache.deleteNode(cache.tail, cache.tail.key)
			cache.InsertNode(key, value)
		} else {
			// add node at start
			// update map
			cache.InsertNode(key, value)
		}
	}
}

// The second return value indicates whether or not the key was found
// in the cache.
func (cache *LRUCache) GetValueFromKey(key string) (int, bool) {
	v, ok := cache.llMap[key]
	if !ok {
		return -1, false
	}

	cache.InsertKeyValuePair(key, v.value)

	return v.value, true
}

// The second return value is false if the cache is empty.
func (cache *LRUCache) GetMostRecentKey() (string, bool) {
	if len(cache.llMap) == 0 {
		return "", false
	}

	return cache.head.key, true
}

func lruCache() {
	c := NewLRUCache(3)
	c.InsertKeyValuePair("b", 2)
	c.InsertKeyValuePair("a", 1)
	c.InsertKeyValuePair("c", 3)
	fmt.Println(c.GetMostRecentKey())
	fmt.Println(c.GetValueFromKey("a"))
	fmt.Println(c.GetMostRecentKey())
	c.InsertKeyValuePair("d", 4)
	fmt.Println(c.GetValueFromKey("b"))
	c.InsertKeyValuePair("a", 5)
	fmt.Println(c.GetValueFromKey("a"))
}
