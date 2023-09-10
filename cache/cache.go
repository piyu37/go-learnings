package main

import "fmt"

type node struct {
	value      string
	next, prev *node
}

type cache struct {
	maxSize    int
	head, tail *node
	llMap      map[string]*node
}

func (c *cache) deleteNode(v string) {
	if c.head == c.tail {
		c.head = nil
		c.tail = nil
		delete(c.llMap, v)
		return
	}

	if v == c.head.value {
		c.head = c.head.next
		c.head.prev = nil
		delete(c.llMap, v)
		return
	}

	if v == c.tail.value {
		c.tail = c.tail.prev
		c.tail.next = nil
		delete(c.llMap, v)
		return
	}

	n := c.llMap[v]
	n.prev.next = n.next
	n.next.prev = n.prev
	delete(c.llMap, v)
}

func (c *cache) insertNode(v string) {
	n := &node{
		value: v,
	}

	if c.head == nil {
		c.head = n
		c.tail = n
		c.llMap[v] = n
		return
	}

	head := c.head
	c.head = n
	n.next = head
	head.prev = n
	c.llMap[v] = n
}

func (c *cache) insertValue(v string) {
	if _, ok := c.llMap[v]; ok {
		c.deleteNode(v)
		c.insertNode(v)
		return
	}

	if len(c.llMap) == c.maxSize {
		c.deleteNode(c.tail.value)
		c.insertNode(v)
	} else {
		c.insertNode(v)
	}
}

func (c *cache) ArrayChallenge(strArr []string) string {
	for i := range strArr {
		value := strArr[i]
		c.insertValue(value)
	}

	output := ""
	curr := c.head

	for curr != nil {
		output += fmt.Sprintf("%s-", curr.value)
		curr = curr.next
	}

	// code goes here
	return output
}

// func main() {
// 	arr := []string{"A", "B", "A", "C", "A", "B"}
// 	c := cache{
// 		maxSize: 5,
// 		llMap:   make(map[string]*node),
// 	}

// 	fmt.Println(c.ArrayChallenge(arr))
// }
