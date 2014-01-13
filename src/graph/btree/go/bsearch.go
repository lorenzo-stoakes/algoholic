package algoholic

// Non-self-balancing implementation.

// Binary search tree property: for each node, all nodes in its left subtree must have values
// of less than its value, all nodes in its right subtree must have values greater than its
// value.

func NewBinarySearchTree(ns []int) *BinaryNode {
	if len(ns) == 0 {
		return nil
	}

	root := NewBinaryNode(ns[0])

	for _, n := range ns[1:] {
		root.Insert(n)
	}

	return root
}

func (node *BinaryNode) Insert(n int) {
	if n < node.Val {
		if node.Left == nil {
			node.Left = NewBinaryNode(n)
		} else {
			node.Left.Insert(n)
		}
	} else if n > node.Val {
		if node.Right == nil {
			node.Right = NewBinaryNode(n)
		} else {
			node.Right.Insert(n)
		}
	}
	// Duplicate nodes are ignored.
}

func (node *BinaryNode) Find(n int) *BinaryNode {
	if node == nil || node.Val == n {
		return node
	} else if n < node.Val {
		return node.Left.Find(n)
	} else {
		return node.Right.Find(n)
	}
}

// In-order traversal, walks values in sorted order. Hence the name :)
func (node *BinaryNode) Walk(visitor func(int)) {
	if node == nil {
		return
	}

	node.Left.Walk(visitor)
	visitor(node.Val)
	node.Right.Walk(visitor)
}

// O(n).
// TODO: Cache.
func (node *BinaryNode) Count() int {
	ret := 0
	node.Walk(func(_ int) {
		ret++
	})

	return ret
}
