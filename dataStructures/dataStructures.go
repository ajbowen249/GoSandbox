package dataStructures

import(
"fmt"
)

type IntLinkedList struct{
	head *IntNode
	tail *IntNode
}

type IntNode struct{
	value int
	previous *IntNode
	next *IntNode
}

func (ll *IntLinkedList) Add(newValue int) {
	newNode := new(IntNode)
	newNode.value = newValue
	
	if ll.head == nil{
		ll.head = newNode
		ll.tail = newNode
		return
	}
	
	ll.tail.next = newNode
	newNode.previous = ll.tail
	ll.tail = newNode
}

func (ll *IntLinkedList) Get(index int) int{
	current := ll.head
	
	for index > 0 {
		if current == nil{
			return 0
		}
		
		current = current.next
		index--
	}
	
	return current.value
}

func (ll *IntLinkedList) Print(){
	current := ll.head
	fmt.Println("traverse")
	for current != nil{
		fmt.Println(current.value)
		current = current.next
	}
}