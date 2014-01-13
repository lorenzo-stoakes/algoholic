package algoholic

type BinaryNode struct {
	Val         int
	Left, Right *BinaryNode
}

func NewBinaryNode(n int) *BinaryNode {
	return &BinaryNode{n, nil, nil}
}
