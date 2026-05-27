package main

import "fmt"

type Node struct {
	key        string
	value      int
	prev, next *Node
}

type LRUCache struct {
	maxSize    int
	head, tail *Node
	llMap      map[string]*Node
}

func NewLRUCache(size int) *LRUCache {
	if size <= 0 {
		panic("LRUCache size must be > 0")
	}
	return &LRUCache{
		maxSize: size,
		llMap:   make(map[string]*Node, size),
	}
}

// --- doubly-linked list helpers ---

func (c *LRUCache) addToFront(n *Node) {
	n.prev = nil
	n.next = c.head

	if c.head != nil {
		c.head.prev = n
	}
	c.head = n

	if c.tail == nil {
		// first element
		c.tail = n
	}
}

func (c *LRUCache) removeNode(n *Node) {
	if n == nil {
		return
	}

	if n.prev != nil {
		n.prev.next = n.next
	} else {
		// removing head
		c.head = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	} else {
		// removing tail
		c.tail = n.prev
	}

	n.prev = nil
	n.next = nil
}

func (c *LRUCache) moveToFront(n *Node) {
	if n == nil || n == c.head {
		return
	}
	c.removeNode(n)
	c.addToFront(n)
}

// --- required API methods ---

func (c *LRUCache) InsertKeyValuePair(key string, value int) {
	if node, ok := c.llMap[key]; ok {
		// update existing value and move to front
		node.value = value
		c.moveToFront(node)
		return
	}

	// new key
	if len(c.llMap) == c.maxSize {
		// evict least recently used (tail)
		lru := c.tail
		delete(c.llMap, lru.key)
		c.removeNode(lru)
	}

	newNode := &Node{key: key, value: value}
	c.addToFront(newNode)
	c.llMap[key] = newNode
}

// The second return value indicates whether or not the key was found
// in the cache.
func (c *LRUCache) GetValueFromKey(key string) (int, bool) {
	n, ok := c.llMap[key]
	if !ok {
		return -1, false
	}
	// mark as most recently used
	c.moveToFront(n)
	return n.value, true
}

// The second return value is false if the cache is empty.
func (c *LRUCache) GetMostRecentKey() (string, bool) {
	if c.head == nil {
		return "", false
	}
	return c.head.key, true
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
