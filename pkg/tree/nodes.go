package tree

import "fmt"
type Node struct {
	Data int64
	Next *Node
}

func NewNode(data int64) Node {
	return Node{Data: data}

}

func (n *Node) Iterate(head Node) {
	node := head;
	fmt.Println(node.Data)

	for node.Next != nil {
		node = *node.Next
		fmt.Println(node.Data);
	}

}

func (n *Node) AppendTail(data int64) {
	endNode := NewNode(data)

	node := n
	for node.Next != nil {
		node = node.Next
	}
	node.Next = &endNode
}

func (n *Node) DeleteNode(head Node, data int64) *Node {
	if head.Data == data {
		return head.Next
	}

	node := head
	for node.Next != nil {
		if node.Next.Data == data {
			node.Next = node.Next.Next
			return &node
		}
		node = *node.Next
	}
	return &node
}