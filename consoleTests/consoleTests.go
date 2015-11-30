package main

import(
"fmt"
"github.com/ajbowen249/GoSandbox/console"
"bufio"
"os"
)

func main(){
	con := console.Default()

	con.ClearScreen()
	//fmt.Println("Testing moveTo:")
	con.MoveTo(10, 10)
	fmt.Print("here")
	
	input := bufio.NewScanner(os.Stdin)
    input.Scan()
}