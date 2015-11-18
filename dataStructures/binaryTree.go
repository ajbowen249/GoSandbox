package dataStructures

type IntBinaryNode struct{
	value int
	left, right *IntBinaryNode
}

type IntBinaryTree struct{
	head *IntBinaryNode
}

func (tree *IntBinaryTree) Add(val int){
	newNode := new(IntBinaryNode)
	newNode.value = val
	
	if tree.head == nil{
		tree.head = newNode
		return
	}
	
	point := tree.head
	
	for {
		if newNode.value < point.value{
			if point.left == nil{
				point.left = newNode
				break
			}else{
				point = point.left
			}
		}else{
			if point.right == nil{
				point.right = newNode
				break
			}else{
				point = point.right
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
	
	visitAscending_rec(node.left, visit)
	visit(node)
	visitAscending_rec(node.right, visit)
}