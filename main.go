package main

import "fmt"

type Node struct {
	Val   string
	Right *Node
	Left  *Node
}

type Queue struct {
	Capacity int
	Head     *Node
	Tail     *Node
	Size     int
}

type Hash map[string]*Node

type Cache struct {
	queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{
		queue: NewQueue(),
		Hash:  Hash{},
	}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head
	return Queue{Head: head, Tail: tail, Size: 0, Capacity: 5}
}

func (cache *Cache) Check(str string) {
	node := &Node{}

	if val, ok := cache.Hash[str]; ok {
		node = cache.Remove(val)
	} else {
		node = &Node{Val: str}
	}
	cache.Add(node)
	cache.Hash[str] = node
}

func (cache *Cache) Remove(node *Node) *Node {
	fmt.Printf("Remove: %s\n", node.Val)

	node.Left.Right = node.Right
	node.Right.Left = node.Left
	cache.queue.Size -= 1
	delete(cache.Hash, node.Val)
	return node
}

func (cache *Cache) Add(node *Node) {
	fmt.Printf("Adding Node : %s \n", node.Val)

	// link right node of head to the adding node
	node.Right = cache.queue.Head.Right
	node.Right.Left = node

	// link head with adding node
	cache.queue.Head.Right = node
	node.Left = cache.queue.Head

	cache.queue.Size++

	if cache.queue.Size > cache.queue.Capacity {
		cache.Remove(cache.queue.Tail.Left)
	}
}

func (queue *Queue) Display() {
	node := queue.Head.Right
	fmt.Printf("%d -[", queue.Size)
	for i := 0; i < queue.Size; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < queue.Size-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func (cache *Cache) Display() {
	cache.queue.Display()
}

func main() {
	fmt.Println("STARTING CACHE .....")
	cache := NewCache()
	for _, word := range []string{"parrot", "avocado", "dragon", "tree", "potato", "tomato", "tree"} {
		cache.Check(word)
		cache.Display()
	}
}
