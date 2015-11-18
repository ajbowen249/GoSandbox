package main

import(
"fmt"
"math/rand"
"time"
ds "github.com/ajbowen249/GoSandbox/dataStructures"
al "github.com/ajbowen249/GoSandbox/algorithms"
)

func main(){
	for i := 0; i < 200000; i += 1000{
		test(i)
	}
}

func test(length int) (time.Duration, time.Duration){
	input := make([]int, length)
	
	for i := range input{
		input[i] = rand.Int()
	}
	
	fmt.Println("Starting selection sort...")
	selectSortStart := time.Now()
	
	al.SortInt(input)
	selectSortDuration := time.Since(selectSortStart)
	
	fmt.Println("select sort took", selectSortDuration)
	
	fmt.Println("Starting binary tree sort...")
	treeSortStart := time.Now()
	resultBuffer := make([]int, length)
	index := 0
	
	tree := new(ds.IntBinaryTree)
	
	for i := range input{
		tree.Add(input[i])
	}
	
	tree.VisitAscending(func(node *ds.IntBinaryNode){
		resultBuffer[index] = node.Value
		index++
	})
	
	treeSortDuration := time.Since(treeSortStart)
	
	fmt.Println("tree sort took", treeSortDuration)
	return selectSortDuration, treeSortDuration
}