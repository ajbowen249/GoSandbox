package dataStructures

type IntLinkedList struct{
	head *IntNode
	tail *IntNode
}

type IntNode struct{
	value int
	previous *IntNode
	next *IntNode
}

func (ll IntLinkedList) Add(newNode *IntNode) {
	if ll.head == nil{
		ll.head = newNode
		ll.tail = newNode
		return
	}
	
	ll.tail.next = newNode
	newNode.previous = ll.tail
	ll.tail = newNode
}