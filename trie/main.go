package main

import "fmt"

// Trie is base Trie tree object
type Trie struct {
	root *Node
}

// Node is base Trie Node object
type Node struct {
	bros  *Node
	child *Node
	data  string
}

func (n *Node) setChild(x string) *Node {
	node := &Node{
		data:  x,
		child: n.child,
	}
	n.child = node
	return node
}

func (n *Node) getChild(x string) *Node {
	node := n.child
	for node != nil {
		if node.data == x {
			break
		}
		node = node.bros
	}
	return node
}

func (n *Node) deleteChild(x string) bool {
	node := n.child
	if node == nil {
		return false
	}

	if node.data == x {
		n.child = n.child.bros
		return true
	}

	for node.bros != nil {
		if node.bros.data == x {
			node.bros = node.bros.bros
			return true
		}
		node = node.bros
	}

	return false
}

func main() {
	fmt.Println("vim-go")
}
