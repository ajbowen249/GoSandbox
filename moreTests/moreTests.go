package main

import(
"fmt"
"math/rand"
"time"
"os"
"strconv"
ds "github.com/ajbowen249/GoSandbox/dataStructures"
al "github.com/ajbowen249/GoSandbox/algorithms"
"github.com/ajbowen249/GoSandbox/table"
)

func main(){
	/*
	writeHeader()

	for i := 0; i <= 200000; i += 1000{
		fmt.Println("sorting", i, "integers")
		selection, binary := test(i)
		updateFile(i, selection, binary)
	}
	*/
	
	tableTest()
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

func writeHeader(){
	f, _ := os.Create("output.csv")
	defer f.Close()
	
	f.WriteString("NumItems, SelectionSort, TreeSort\n")
}

func updateFile(numItems int, selection time.Duration, binary time.Duration){
	f, _ := os.OpenFile("output.csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	defer f.Close()
	
	output := fmt.Sprintf("%v, %v, %v\n", numItems, selection.Seconds() * 1000, binary.Seconds() * 1000)
	f.WriteString(output)
}

func tableTest(){
	tb := table.New()
	tb.HeaderAlign = table.CENTERLEFTBIAS
	tb.CellAlign = table.RIGHT
	tb.AddColumn("X")
	tb.AddColumn("Y")
	tb.AddColumn("X + Y")
	
	for i := 0; i < 15; i++{
		x := rand.Int()
		y := rand.Int()
		z := x + y
		
		tb.AddRow(strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(z))
	}
	
	tb.Output(func (row string) { fmt.Println(row)})
}