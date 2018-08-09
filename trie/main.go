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
	data  rune
}

// NewNode returns Trie object
func NewNode() *Trie {
	return &Trie{
		root: new(Node),
	}
}

func (n *Node) setChild(r rune) *Node {
	node := &Node{
		data: r,
		bros: n.child,
	}
	n.child = node
	return node
}

func (n *Node) getChild(r rune) *Node {
	node := n.child
	for node != nil {
		if node.data == r {
			break
		}
		node = node.bros
	}
	return node
}

func (n *Node) deleteChild(r rune) bool {
	node := n.child
	if node == nil {
		return false
	}

	if node.data == r {
		n.child = n.child.bros
		return true
	}

	for node.bros != nil {
		if node.bros.data == r {
			node.bros = node.bros.bros
			return true
		}
		node = node.bros
	}

	return false
}

func (t *Trie) insert(seq string) {
	node := t.root

	for _, r := range seq {
		child := node.getChild(r)
		if child == nil {
			child = node.setChild(r)
		}
		node = child
	}
}

func (t *Trie) search(seq string) bool {
	node := t.root

	for _, r := range seq {
		node = node.getChild(r)
		if node == nil {
			return false
		}
	}

	return true
}

func main() {
	t := NewNode()
	t.insert("abc")
	fmt.Println(t.root)
	fmt.Println(t.root.child.data)
	fmt.Println(t.root.child.child.data)
	fmt.Println(t.root.child.child.child.data)

	fmt.Println(t.search("abc"))
	fmt.Println(t.search("cba"))
}
