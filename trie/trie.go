package trie

// Trie is base Trie tree object
type Trie struct {
	root *Node
}

// Node is base Trie Node object
type Node struct {
	bros  *Node
	child *Node
	data  rune
	item  interface{}
	last  bool
}

// NewTrie returns Trie object
func NewTrie() *Trie {
	return &Trie{
		root: new(Node),
	}
}

func (n *Node) setChild(r rune, item interface{}) *Node {
	node := &Node{
		data: r,
		bros: n.child,
		item: item,
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

// Insert inserts seq in the trie tree
func (t *Trie) Insert(seq string, item interface{}) {
	node := t.root

	for i, r := range seq {
		child := node.getChild(r)
		if child == nil {
			child = node.setChild(r, item)
			node = child
			continue
		}

		if i == len(seq)-1 {
			child.item = item
			child.last = true
		}

		node = child
	}
}

// Search searches item of given seq in the tre tree
func (t *Trie) Search(seq string) (bool, interface{}) {
	node := t.root

	for _, r := range seq {
		node = node.getChild(r)
		if node == nil {
			return false, nil
		}
	}

	return true, node.item
}

// Delete deletes node of given seq in the trie truee
func (t *Trie) Delete(seq string) bool {
	node := t.root

	for i, r := range seq {
		child := node.child
		if child == nil {
			return false
		}

		if i == len(seq)-1 && child.data == r {
			node.child = child.bros
			return true
		}

		for child.bros != nil {
			if i == len(seq)-1 && child.bros.data == r {
				child.bros = child.bros.bros
				return true
			}
			child = child.bros
		}

		node = node.child
	}

	return false
}

func (n *Node) traverse() []string {
	child := n.child
	if child == nil {
		return nil
	}

	list := make([]string, 0)
	for child != nil {
		li := child.traverse()
		if li == nil {
			list = append(list, string(child.data))
		} else {
			for _, v := range li {
				list = append(list, string(child.data)+v)
			}
		}

		val := string(child.data)
		ok := exist(list, val)
		if !ok && child.last {
			list = append(list, val)
		}

		child = child.bros
	}

	return list
}

// Traverse Searches all seq
func (t *Trie) Traverse() []string {
	node := t.root
	return node.traverse()
}

func exist(lists []string, val string) bool {
	for _, v := range lists {
		if v == val {
			return true
		}
	}
	return false
}
