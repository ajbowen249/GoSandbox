package main

import(
"fmt"
"github.com/ajbowen249/GoSandbox/console"
"bufio"
"os"
)

func main(){
	console.ClearScreen()
	//fmt.Println("Testing moveTo:")
	console.MoveTo(10, 10)
	fmt.Print("here")
	
	input := bufio.NewScanner(os.Stdin)
    input.Scan()
}