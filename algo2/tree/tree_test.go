package tree_test

import (
	"reflect"
	"testing"

	"github.com/sKudryashov/algo/tree"
)

func TestTreeTraversal(t *testing.T) {
	///      4
	/// 2         6
	//1   3     5     8
	// output 1234568
	unsortedKey := tree.KeySet{6, 5, 8, 3, 2, 1, 4, 9, 10}
	h := tree.NewBSTHandler()
	node := h.BuildTree(unsortedKey)
	h.InOrderPrinter(node)
	h.PreOrderPrinter(node)
	h.PostOrderPrinter(node, 0)
	lvlOrder := h.LevelOrderTraversal()
	h.RangeTraversal()
	if reflect.DeepEqual(lvlOrder, []int{5, 2, 8, 1, 3, 6, 9, 4, 10}) {
		t.Fatalf("expect 5, 2, 8, 1, 3, 6, 9, 4, 10 sequence")
	}

}
