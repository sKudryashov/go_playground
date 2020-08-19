package tree

import (
	"sort"
	"strconv"

	"github.com/sKudryashov/algo/queue"
)

type Node struct {
	left  *Node
	right *Node
	val   int
}

type KeySet []int

func (s KeySet) Len() int {
	return len(s)
}

func (s KeySet) Less(i, j int) bool {
	if s[i] < s[j] {
		return true
	}
	return false
}

func (s KeySet) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//BSTHandler represents a BST handling object
type BSTHandler struct {
	inOrderSet, postOrderSet, rootOrderSet, levelOrderSet []int
	keySet                                                []int
	node                                                  *Node
	queue                                                 *queue.Queue
}

// NewBSTHandler constructor
func NewBSTHandler() *BSTHandler {
	return &BSTHandler{
		inOrderSet:   make([]int, 0, 100),
		postOrderSet: make([]int, 0, 100),
		rootOrderSet: make([]int, 0, 100),
	}
}

// func (h *BSTHandler) buildTreePostOrder(left, right int) *Node {
// 	root := (left + right) / 2
// 	tree := &Node{
// 		val: h.keySet[root],
// 	}
// 	tree.left = h.buildTreePostOrder(root-1, left)
// 	tree.right = h.buildTreePostOrder(root+1, right)
// 	return tree
// 	///      4
// 	/// 2         6
// 	//1   3     5     8
// 	// output 1234568
// }

func (h *BSTHandler) buildTreePreOrder(left, right int) *Node {
	if left > right {
		//endof traversal (we could add nil here)
		return nil
	}
	root := (left + right) / 2
	tree := &Node{
		val: h.keySet[root],
	}
	tree.left = h.buildTreePreOrder(left, root-1)
	tree.right = h.buildTreePreOrder(root+1, right)
	return tree
	///      4
	/// 2         6
	//1   3     5     8
	// output 1234568
}

// BuildTree b
func (h *BSTHandler) BuildTree(keySet KeySet) *Node {
	sort.Sort(keySet)
	h.keySet = keySet
	node := h.buildTreePreOrder(0, keySet.Len()-1)
	h.node = node
	return node
}

// PreOrderPrinter pr
func (h *BSTHandler) PreOrderPrinter(nd *Node) {
	// println("val ", nd.val)
	h.rootOrderSet = append(h.rootOrderSet, nd.val)
	if nd.left != nil {
		h.PreOrderPrinter(nd.left)
	}
	if nd.right != nil {
		h.PreOrderPrinter(nd.right)
	}
	///      4
	/// 2         6
	//1   3     5     8
	// output 4213658
}

// InOrderPrinter ino pr
func (h *BSTHandler) InOrderPrinter(nd *Node) {
	///      4
	/// 2         6
	//1   3     5     8
	// output 1234568
	if nd.left != nil {
		// add nil node here to be able to restore
		h.InOrderPrinter(nd.left)
	}
	// println("val ", nd.val)
	h.inOrderSet = append(h.inOrderSet, nd.val)
	if nd.right != nil {
		h.InOrderPrinter(nd.right)
	}
}

// PostOrderPrinter postorder prnt
func (h *BSTHandler) PostOrderPrinter(nd *Node, index int) {
	// rem := index % 2
	// if rem == 0 {
	if nd.left != nil {
		index++
		// println("traverse ", nd.val)
		h.PostOrderPrinter(nd.left, index)
	}
	// } else {
	if nd.right != nil {
		index++
		// println("traverse ", nd.val)
		h.PostOrderPrinter(nd.right, index)
	}
	// }
	// println("append ", nd.val)
	h.postOrderSet = append(h.postOrderSet, nd.val)
	///      4
	/// 2         6
	//1   3     5     8
	// expect 1325864
	// actual 1325864
}

// excellent recursion example
func (h *BSTHandler) height(node *Node) int {
	if node == nil {
		return 0
	}
	lheight := h.height(node.left)
	rheight := h.height(node.right)
	if lheight > rheight {
		return lheight + 1
	} else {
		return rheight + 1
	}
}

// GetLeverOrdered returns level ordered tree
func (h *BSTHandler) getLeverOrdered(n *Node, lvl int) {
	if n == nil {
		return
	}
	if lvl == 1 {
		println(" level data  1 ", n.val)
		h.levelOrderSet = append(h.levelOrderSet, n.val)
	}
	if lvl > 1 {
		h.getLeverOrdered(n.left, lvl-1)
		// println("lvl ", lvl)
		// h.levelOrderSet = append(h.levelOrderSet, n.val)
		h.getLeverOrdered(n.right, lvl-1)
		// h.levelOrderSet = append(h.levelOrderSet, n.val)
	}
}

// LevelOrderTraversal performs level-order traversal by given tree
func (h *BSTHandler) LevelOrderTraversal() []int {
	he := h.height(h.node)
	h.levelOrderSet = make([]int, 0, he*he)
	for i := 1; i <= he; i++ {
		h.getLeverOrdered(h.node, i)
	}
	return h.levelOrderSet
}

// BFSTraversal is the same as LevelOrderTraversal but uses queue instead of callstack
func (h *BSTHandler) BFSTraversal() []int {
	h.queue.Enqeue(h.node)
	// hh := h.height(h.node)
	for !h.queue.Empty() {
		for h.queue.Len() > 0 {
			node := h.queue.Deqeue().(*Node)
			if node.left != nil {
				h.queue.Enqeue(node.left)
			}
			if h.node.right != nil {
				h.queue.Enqeue(node.right)
			}
		}
	}

	return h.levelOrderSet
}

//CalculateTreeWidth calculates tree width
func (h *BSTHandler) CalculateTreeWidth() {

}

// RangeTraversal Traverse the tree in the inorder traversal.
// If the Binary search tree is traversed in inorder traversal the keys are traversed in increasing order.
// So while traversing the keys in the inorder traversal. If the key lies in the range print the key
// else skip the key.
func (h *BSTHandler) RangeTraversal() {
	strconv.Atoi("12")
}

// KSmallestTraversal The Inorder Traversal of a BST traverses the nodes in increasing order.
// So the idea is to traverse the tree in Inorder. While traversing, keep track of the count of the
// nodes visited. If the count becomes k, print the node.
func (h *BSTHandler) KSmallestTraversal() {

}

//FloorTraversal - just go inorder  and find less or equal
func (h *BSTHandler) FloorTraversal() {

}

// CeilTraversal - just go inorder  and find less or equal
func (h *BSTHandler) CeilTraversal() {

}
