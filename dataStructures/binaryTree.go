package dataStructures

type IntBinaryNode struct{
	Value int
	Left, Right *IntBinaryNode
}

type IntBinaryTree struct{
	head *IntBinaryNode
}

func (tree *IntBinaryTree) Add(val int){
	newNode := new(IntBinaryNode)
	newNode.Value = val
	
	if tree.head == nil{
		tree.head = newNode
		return
	}
	
	point := tree.head
	
	for {
		if newNode.Value < point.Value{
			if point.Left == nil{
				point.Left = newNode
				break
			}else{
				point = point.Left
			}
		}else{
			if point.Right == nil{
				point.Right = newNode
				break
			}else{
				point = point.Right
			}
		}
	}
}

func (tree *IntBinaryTree) VisitAscending(visit func(node *IntBinaryNode)){
	visitAscending_rec(tree.head, visit)
}

func visitAscending_rec(node *IntBinaryNode, visit func(node *IntBinaryNode)){
	if node == nil{
		return
	}
	
	visitAscending_rec(node.Left, visit)
	visit(node)
	visitAscending_rec(node.Right, visit)
}