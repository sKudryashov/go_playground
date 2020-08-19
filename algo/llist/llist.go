package llist

type ListNode struct {
	Next *ListNode
	Val  int
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = MergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoLists(l1, l2.Next)
		return l2
	}
}

func ReverseListTrue(pList *ListNode) *ListNode {
	pCurr := pList
	var prev *ListNode = nil
	for {
		if pCurr == nil {
			break
		}
		pTemp := pCurr.Next
		pCurr.Next = prev
		prev = pCurr
		pCurr = pTemp
	}
	return prev
}

func reverseListKItems(head *ListNode, k int) *ListNode {
	var newHead *ListNode = nil
	var ptr *ListNode = head
	for k > 0 {
		nextNode := ptr.Next
		ptr.Next = newHead
		newHead = ptr
		ptr = nextNode
		k--
	}
	return newHead
}

func ReverseListK(head *ListNode, k int) *ListNode {
	ptr := head
	var ktail *ListNode = nil
	var newHead *ListNode = nil
	var count int
	for ptr != nil {
		for count < k && ptr != nil {
			ptr = ptr.Next
			count++
		}
		ptr = head
		if count == k {
			revHead := reverseListKItems(head, k)
			if newHead == nil {
				newHead = revHead
			}
			if ktail != nil {
				ktail.Next = revHead
			}
			ktail = head
			head = ptr
		}
	}
	if ktail != nil {
		ktail.Next = head
	}
	if newHead == nil {
		return head
	}
	return newHead
}

func MergeKLists() {
	//from Queue import PriorityQueue

	// class Solution(object):
	// def mergeKLists(self, lists):
	// 	"""
	// 	:type lists: List[ListNode]
	// 	:rtype: ListNode
	// 	"""
	// 	head = point = ListNode(0)
	// 	q = PriorityQueue()
	// 	for l in lists:
	// 		if l:
	// 			q.put((l.val, l))
	// 	while not q.empty():
	// 		val, node = q.get()
	// 		point.next = ListNode(val)
	// 		point = point.next
	// 		node = node.next
	// 		if node:
	// 			q.put((node.val, node))
	// 	return head.next
}

// class Solution {

//     public ListNode reverseLinkedList(ListNode head, int k) {

//         // Reverse k nodes of the given linked list.
//         // This function assumes that the list contains
//         // atleast k nodes.
//         ListNode new_head = null;
//         ListNode ptr = head;

//         while (k > 0) {

//             // Keep track of the next node to process in the
//             // original list
//             ListNode next_node = ptr.next;

//             // Insert the node pointed to by "ptr"
//             // at the beginning of the reversed list
//             ptr.next = new_head;
//             new_head = ptr;

//             // Move on to the next node
//             ptr = next_node;

//             // Decrement the count of nodes to be reversed by 1
//             k--;
//         }

//         // Return the head of the reversed list
//         return new_head;

//     }

//     public ListNode reverseKGroup(ListNode head, int k) {

//         int count = 0;
//         ListNode ptr = head;

//         // First, see if there are atleast k nodes
//         // left in the linked list.
//         while (count < k && ptr != null) {
//             ptr = ptr.next;
//             count++;
//         }

//         // If we have k nodes, then we reverse them
//         if (count == k) {

//             // Reverse the first k nodes of the list and
//             // get the reversed list's head.
//             ListNode reversedHead = this.reverseLinkedList(head, k);

//             // Now recurse on the remaining linked list. Since
//             // our recursion returns the head of the overall processed
//             // list, we use that and the "original" head of the "k" nodes
//             // to re-wire the connections.
//             head.next = this.reverseKGroup(ptr, k);
//             return reversedHead;
//         }

//         return head;
//     }
// }
