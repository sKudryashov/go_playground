package llist

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestLListManager_MakeDeepCopyRandom(t *testing.T) {
	var head, one, two, three, four *RNode
	four = &RNode{
		Value: 5,
		Next:  nil,
		// Random: three,
	}
	three = &RNode{
		Value: 4,
		Next:  four,
		// Random: one,
	}
	two = &RNode{
		Value: 3,
		Next:  three,
		// Random: four,
	}
	one = &RNode{
		Value: 2,
		Next:  two,
		// Random: head,
	}
	head = &RNode{
		Value: 1,
		// Random: three,
		Next: one,
	}
	four.Random = three
	three.Random = one
	two.Random = four
	one.Random = head
	head.Random = three

	m := NewLListManager()
	spew.Dump(" before  ", head)
	got := m.MakeDeepCopyRandom(head)
	spew.Dump(" after  ", got)
}
