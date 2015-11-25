package main

import(
"fmt"
"math/rand"
"time"
"os"
ds "github.com/ajbowen249/GoSandbox/dataStructures"
al "github.com/ajbowen249/GoSandbox/algorithms"
"github.com/ajbowen249/GoSandbox/table"
)

func main(){
	writeHeader()
	tb := startTable()
	
	for i := 0; i <= 18000; i += 1000{
		selection, binary := test(i)
		updateTable(tb, i, selection, binary)
		updateFile(i, selection, binary)
	}
	
	tb.Output(func (row string) { fmt.Println(row)})
}

func startTable() *table.Table{
	tb := table.New()
	tb.AddColumn("NumItems")
	tb.AddColumn("Selection Sort Time (ms)")
	tb.AddColumn("Binary Sort Time (ms)")
	
	return tb
}

func test(length int) (time.Duration, time.Duration){
	input := make([]int, length)
	
	rand.Seed(time.Now().UnixNano())
	for i := range input{
		input[i] = rand.Int()
	}
	
	selectSortStart := time.Now()
	
	al.SortInt(input)
	selectSortDuration := time.Since(selectSortStart)
	
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

func updateTable(tb *table.Table, numItems int, selection time.Duration, binary time.Duration){
	tb.AddRow(fmt.Sprintf("%v", numItems), fmt.Sprintf("%.3f", selection.Seconds() * 1000), fmt.Sprintf("%.3f", binary.Seconds() * 1000))
}