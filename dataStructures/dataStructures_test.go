package dataStructures

import 
(
"testing"
)

type IntLLGetCase struct{
	inputs []int
}

func TestIntLLGet(t *testing.T){
	cases := []IntLLGetCase{
		{[]int{3, 7, 1, 4}},
		{[]int{3}},
		{[]int{3, 7, 1, 4, 9, 4, 0}},
	}
	
	for _, c := range cases{
		ll := new(IntLinkedList)
		
		for i := range c.inputs{
			ll.Add(c.inputs[i])
		}
		
		for i := 0; i < len(c.inputs); i++{
			expected := c.inputs[i]
			result := ll.Get(i)
			
			if expected != result{
				t.Errorf("Expected Get(%v) == %v, but was %v", i, expected, result)
			}
		}
	}
}

type IntBinaryTreeCase struct{
	inputs, order []int
}

func TestIntBinaryTreeVisit(t *testing.T){
	cases := []IntBinaryTreeCase{
		{[]int{3, 7, 1, 4}, []int{1, 3, 4, 7}},
		{[]int{2}, []int{2}},
		{[]int{50, 0, -3, 1, -10, 19}, []int{-10, -3, 0, 1, 19, 50}},
	}
	
	for _, c := range cases{
		tree := new(IntBinaryTree)
		
		for i := range c.inputs{
			tree.Add(c.inputs[i])
		}
		
		index := 0
		
		tree.VisitAscending(func(node *IntBinaryNode){
			if node.value != c.order[index]{
				t.Errorf("Expected tree[%v] == %v, but was %v", index, c.order[index], node.value)
			}
			
			index++
		})
	}
}