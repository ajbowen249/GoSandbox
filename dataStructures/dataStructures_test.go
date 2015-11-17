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

