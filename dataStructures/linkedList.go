package dataStructures

import(
"fmt"
)

type IntLinkedList struct{
	Head *IntLLNode
	Tail *IntLLNode
}

type IntLLNode struct{
	Value int
	Previous *IntLLNode
	Next *IntLLNode
}

func (ll *IntLinkedList) Add(newValue int) {
	newNode := new(IntLLNode)
	newNode.Value = newValue
	
	if ll.Head == nil{
		ll.Head = newNode
		ll.Tail = newNode
		return
	}
	
	ll.Tail.Next = newNode
	newNode.Previous = ll.Tail
	ll.Tail = newNode
}

func (ll *IntLinkedList) Get(index int) int{
	current := ll.Head
	
	for index > 0 {
		if current == nil{
			return 0
		}
		
		current = current.Next
		index--
	}
	
	return current.Value
}

func (ll *IntLinkedList) Print(){
	current := ll.Head
	for current != nil{
		fmt.Println(current.Value)
		current = current.Next
	}
}