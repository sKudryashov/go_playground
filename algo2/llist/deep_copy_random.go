package llist

type RNode struct {
	Next   *RNode
	Random *RNode
	Value  int
}

type LListManager struct {
	visited map[*RNode]*RNode
}

func NewLListManager() *LListManager {
	return &LListManager{
		visited: make(map[*RNode]*RNode),
	}
}

func (m *LListManager) MakeDeepCopyRandom(head *RNode) *RNode {
	if head == nil {
		return nil
	}
	oldNode := head
	newNode := &RNode{
		Value:  head.Value,
		Random: head.Random,
		Next:   oldNode.Next,
	}
	m.visited[oldNode] = newNode

	for oldNode != nil {
		newNode.Random = m.getClonedNode(newNode.Random)
		newNode.Next = m.getClonedNode(newNode.Next)
		oldNode = oldNode.Next
		newNode = newNode.Next
	}
	return m.visited[head]
}

func (m *LListManager) getClonedNode(head *RNode) *RNode {
	if head == nil {
		return nil
	}
	if node, ok := m.visited[head]; ok {
		return node
	}
	m.visited[head] = &RNode{
		Value: head.Value,
	}
	return m.visited[head]
}

/*
// Definition for a Node.
class Node {
    public int val;
    public Node next;
    public Node random;

    public Node() {}

    public Node(int _val,Node _next,Node _random) {
        val = _val;
        next = _next;
        random = _random;
    }
};
*/
// public class Solution {
// 	// Visited dictionary to hold old node reference as "key" and new node reference as the "value"
// 	HashMap<Node, Node> visited = new HashMap<Node, Node>();

// 	public Node getClonedNode(Node node) {
// 	  // If the node exists then
// 	  if (node != null) {
// 		// Check if the node is in the visited dictionary
// 		if (this.visited.containsKey(node)) {
// 		  // If its in the visited dictionary then return the new node reference from the dictionary
// 		  return this.visited.get(node);
// 		} else {
// 		  // Otherwise create a new node, add to the dictionary and return it
// 		  this.visited.put(node, new Node(node.val, null, null));
// 		  return this.visited.get(node);
// 		}
// 	  }
// 	  return null;
// 	}

// 	public Node copyRandomList(Node head) {

// 	  if (head == null) {
// 		return null;
// 	  }

// 	  Node oldNode = head;

// 	  // Creating the new head node.
// 	  Node newNode = new Node(oldNode.val);
// 	  this.visited.put(oldNode, newNode);

// 	  // Iterate on the linked list until all nodes are cloned.
// 	  while (oldNode != null) {
// 		// Get the clones of the nodes referenced by random and next pointers.
// 		newNode.random = this.getClonedNode(oldNode.random);
// 		newNode.next = this.getClonedNode(oldNode.next);

// 		// Move one step ahead in the linked list.
// 		oldNode = oldNode.next;
// 		newNode = newNode.next;
// 	  }
// 	  return this.visited.get(head);
// 	}
//   }
