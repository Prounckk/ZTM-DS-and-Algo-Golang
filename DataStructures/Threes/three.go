package threes

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (t *BinaryTree) insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &Node{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *Node) insert(data int) {
	if n == nil {
		return
	} else if data < n.data {
		if n.left == nil {
			n.left = &Node{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func (t *BinaryTree) search(data int) bool {
	if t.root == nil {
		return false
	} else {
		return t.root.search(data)
	}
}

//  6
// 4 10
func (n *Node) search(data int) bool {
	if n == nil {
		return false
	} else if n.data == data {
		return true
	} else if data < n.data {
		if n.left == nil {
			return false
		} else {
			return n.left.search(data)

		}
	} else {
		if n.right == nil {
			return false
		} else {
			return n.right.search(data)
		}
	}

}
