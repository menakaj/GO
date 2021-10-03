package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}
type Cache struct {
	Head        *Node
	Tail        *Node
	Size        int
	Counter     int
	lookupTable map[string]*Node
}

func main() {

	c := Cache{
		Head:        nil,
		Tail:        nil,
		Size:        5,
		Counter:     0,
		lookupTable: map[string]*Node{},
	}

	c.addToHead("one", 1)
	c.addToHead("two", 2)
	c.addToHead("three", 3)
	c.addToHead("four", 4)
	c.addToHead("Five", 5)
	c.addToHead("Six", 6)
	fmt.Println(c.Head)
	fmt.Println(c.lookupTable)
	c.print()
	fmt.Println("Get four >> ", c.get("four"))
	c.print()
	c.addToHead("three", 3)
	c.addToHead("four", 4)
	c.addToHead("Five", 5)
	c.addToHead("Six", 6)
	c.print()

}

func (c *Cache) evict() {
	fmt.Println("Evicting ", c.Tail.Value)
	newTail := c.Tail.Prev
	newTail.Next = nil
	c.Tail = newTail
}

func (c *Cache) addToHead(key string, value int) {
	head := c.Head
	fmt.Println("Adding to head", key, value)
	if head == nil {
		node := &Node{
			Value: value,
			Next:  nil,
			Prev:  nil}
		c.Head = node
		c.Tail = node
		c.lookupTable[key] = c.Head
		fmt.Println(c.Head)
		c.Counter = 1
		return
	}

	c.Counter++
	if c.Size < c.Counter {
		c.evict()
	}

	node := &Node{
		Value: value,
		Next:  head,
		Prev:  nil,
	}

	fmt.Println("Current head", head)
	head.Prev = node
	c.Head = node
	c.lookupTable[key] = node
	fmt.Println("C.Head ", c.Head)
}

func (c *Cache) len() int {
	return c.Size
}

func (c *Cache) print() {
	head := c.Head
	if head == nil {
		fmt.Println("Cache is empty...")
	}

	for head != nil {
		fmt.Println(head.Value)
		head = head.Next
	}
}

func (c *Cache) get(key string) int {
	// check in the map. If empty, return null
	if c.lookupTable[key] == nil {
		return -1
	}

	// get the node
	n := c.lookupTable[key]

	// Create a copy
	newN := n
	// Move the new node to head
	// n's previous's  next -> n.next
	// n's Next  -> n.previoue.next
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	c.Head.Prev = newN
	newN.Next = c.Head

	c.Head = newN
	return newN.Value
}
