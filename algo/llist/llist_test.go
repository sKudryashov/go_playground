package llist

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestMergeList(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 8,
				},
			},
		},
	}
	MergeTwoLists(l1, l2)
}

func TestReverseList(t *testing.T) {
	ll := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	p := ReverseListTrue(ll)
	spew.Dump(p)
}
