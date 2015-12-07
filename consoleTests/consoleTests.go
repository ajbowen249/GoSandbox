package main

import(
"fmt"
"github.com/ajbowen249/GoSandbox/console"
)

func main(){
	con := console.Default()

	con.ClearScreen()
	//fmt.Println("Testing moveTo:")
	con.MoveTo(10, 10)
	fmt.Print("here")
	
	for{
		isHit, hit := con.KetKey()
		
		if isHit{
			fmt.Print(hit)
		}
	}
}