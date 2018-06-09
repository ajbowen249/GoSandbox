package rogueConsole

import (
	"testing"

	"github.com/ajbowen249/GoSandbox/console"
)

func TestConsole1(t *testing.T) {
	con := NewRogueConsole(25, 9, 5, 5)

	redArray := FillArrayI(25, 9, console.ChFgRed)
	greenArray := FillArrayI(25, 9, console.ChFgGreen)
	blueArray := FillArrayI(25, 9, console.ChFgBlue)
	yellowArray := FillArrayI(25, 9, console.ChFgYellow)

	bg1 := StringToArray(25, 9,
		"┌───────────────────────┐"+
			"│                       │"+
			"│                       │"+
			"│                       │"+
			"│                       │"+
			"│                       │"+
			"│                       │"+
			"│                       │"+
			"└───────────────────────┘")

	bg2 := StringToArray(25, 9,
		"                         "+
			"                         "+
			"  /─\\                    "+
			"  \\─/                    "+
			"                         "+
			"          1234           "+
			"                         "+
			"                         "+
			"                         ")

	fg1 := StringToArray(25, 9,
		"                         "+
			"                         "+
			"  ab                     "+
			"                         "+
			"                         "+
			"                         "+
			"                         "+
			"                         "+
			"                         ")

	fg2 := StringToArray(25, 9,
		"                         "+
			"                         "+
			"  c                      "+
			"  d                      "+
			"                         "+
			"                         "+
			"                         "+
			"                         "+
			"                         ")

	Replace(&bg1, ' ', TransparancyChar)
	Replace(&bg2, ' ', TransparancyChar)
	Replace(&fg1, ' ', TransparancyChar)
	Replace(&fg2, ' ', TransparancyChar)

	con.AddBackground(bg1, redArray)
	con.AddBackground(bg2, greenArray)
	con.AddForeground(fg1, blueArray)
	con.AddForeground(fg2, yellowArray)

	sprite := new(Sprite)
	sprite.X = 4
	sprite.Y = 3

	spriteChars := StringToArray(2, 2, "SSS ")
	Replace(&spriteChars, ' ', TransparancyChar)

	spriteColors := FillArrayI(2, 2, console.ChFgWhite)
	sprite.SetGraphics(spriteChars, spriteColors)

	con.RegisterSprite(sprite, 0)

	sprite2 := new(Sprite)
	sprite2.X = 1
	sprite2.Y = 1
	sprite2.SetGraphics([][]rune{{'a', 'a', 'a', 'a'}}, [][]int{{console.ChFgDarkBlue, console.ChFgWhite, console.ChFgWhite, console.ChFgWhite}})
	con.RegisterSprite(sprite2, 3)

	sprite3 := new(Sprite)
	sprite3.X = 2
	sprite3.Y = 1
	sprite3.SetGraphics([][]rune{{'b', 'b', 'b'}}, [][]int{{console.ChFgDarkCyan, console.ChFgWhite, console.ChFgWhite}})
	con.RegisterSprite(sprite3, 2)

	sprite4 := new(Sprite)
	sprite4.X = 3
	sprite4.Y = 1
	sprite4.SetGraphics([][]rune{{'c', 'c'}}, [][]int{{console.ChFgDarkGreen, console.ChFgWhite}})
	con.RegisterSprite(sprite4, 1)

	sprite5 := new(Sprite)
	sprite5.X = 4
	sprite5.Y = 1
	sprite5.SetGraphics([][]rune{{'d'}}, [][]int{{console.ChFgGrey}})
	con.RegisterSprite(sprite5, 0)

	con.CameraX = 0
	con.CameraY = 0

	expectedString :=
		"┌────" +
			"│abcd" +
			"│ cb\\" +
			"│ d─S" +
			"│   S"

	expectedArray := StringToArray(5, 5, expectedString)

	expectedColors := [][]int{
		{console.ChFgRed, console.ChFgRed, console.ChFgRed, console.ChFgRed, console.ChFgRed},
		{console.ChFgRed, console.ChFgDarkBlue, console.ChFgDarkCyan, console.ChFgDarkGreen, console.ChFgGrey},
		{console.ChFgRed, defaultTransparencyColor, console.ChFgYellow, console.ChFgBlue, console.ChFgGreen},
		{console.ChFgRed, defaultTransparencyColor, console.ChFgYellow, console.ChFgGreen, console.ChFgWhite},
		{console.ChFgRed, defaultTransparencyColor, defaultTransparencyColor, defaultTransparencyColor, console.ChFgWhite},
	}

	con.Visit(func(r rune, i int, row int, col int) {
		if r != expectedArray[row][col] {
			t.Errorf("Expected runes[%v][%v] == %v, but was %v.", row, col, expectedArray[row][col], r)
		}

		if i != expectedColors[row][col] {
			t.Errorf("Expected colors[%v][%v] == %v, but was %v.", row, col, expectedColors[row][col], i)
		}
	})
}
