package main

import (
	"fmt"
	"strings"

	"github.com/ajbowen249/GoSandbox/console"
)

func main() {
	numCols := 80
	numRows := 25

	console.ClearScreen(numCols, numRows)
	x, y := 0, 0
	console.MoveTo(x, y)
	sprite := "O"
	fmt.Print(sprite)
	console.MoveTo(numCols-1, numRows-1)
	defaultConsoleProperties := console.GetCharacterProperties()

	for {
		isHit, char := console.GetKey()
		if isHit {
			console.MoveTo(x, y)
			fmt.Print(" ")

			char = strings.ToLower(char)
			switch char {
			case "w":
				y--
			case "a":
				x--
			case "s":
				y++
			case "d":
				x++
			}

			console.MoveTo(x, y)

			console.SetCharacterProperties(console.ChFgBlue | console.ChBgYellow | console.ChUnderline)
			fmt.Print(sprite)
			console.SetCharacterProperties(defaultConsoleProperties)

			console.MoveTo(numCols-1, numRows-1)
		}
	}
}
