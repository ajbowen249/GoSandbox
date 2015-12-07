package main

import(
"fmt"
"strings"
"github.com/ajbowen249/GoSandbox/console"
)

func main(){
	con := console.Default()

	con.ClearScreen()
	x, y := 0, 0
	con.MoveTo(x, y)
	sprite := "O"
	fmt.Print(sprite)
	con.MoveTo(con.NumCols - 1, con.NumRows - 1)
	
	for{
		isHit, char := con.GetKey()
		if isHit{
			con.MoveTo(x, y)
			fmt.Print(" ")
			
			char = strings.ToLower(char)
			switch char{
				case "w":
					y--
				case "a":
					x--
				case "s":
					y++
				case "d":
					x++
			}
			
			con.MoveTo(x, y)
			fmt.Print(sprite)
			
			con.MoveTo(con.NumCols - 1, con.NumRows - 1)
		}
	}
}