package main

import (
	"fmt"

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
	_, info := console.GetDefaultAttributes()

	for {
		isHit, event := console.GetKeyEX()
		if isHit && event.IsSpecial {
			console.MoveTo(x, y)
			fmt.Print(" ")

			switch event.SpecialChar {
			case console.ScArrowUp:
				y--
			case console.ScArrowLeft:
				x--
			case console.ScArrowDown:
				y++
			case console.ScArrowRight:
				x++
			case console.ScEsc:
				return
			}

			console.MoveTo(x, y)

			console.SetCharacterProperties(console.ChFgBlue | console.ChBgYellow | console.ChUnderline)
			fmt.Print(sprite)
			console.SetCharacterProperties(info.CharacterColor)

			console.MoveTo(numCols-1, numRows-1)
		}
	}
}
