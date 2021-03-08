package tree

import "fmt"

// Node presents a normal tree node.
type Node struct {
	Element interface{}
	FirstChild *Node
	NextSibling *Node
}

func (n Node) ListAll(depth int) {
	fmt.Printf("")
}