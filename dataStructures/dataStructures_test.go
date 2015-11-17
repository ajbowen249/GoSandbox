package dataStructures

import "testing"

type intLLGetCase struct{
	inputs []int
}

func TestIntLLGet(t *testing.T){
	cases := []intLLGetCase{
		{[]int{3, 7, 1, 4}},
	}
	
	for _, c := range cases{
		ll := new(IntLinkedList)
		
		for n := range c.inputs{
			node := new(IntNode)
			node.value = n
			ll.Add(node)
		}
	}
}

