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

	bg1 :=
		"┌───────────────────────┐" +
			"│                       │" +
			"│                       │" +
			"│                       │" +
			"│                       │" +
			"│                       │" +
			"│                       │" +
			"│                       │" +
			"└───────────────────────┘"

	bg2 :=
		"                         " +
			"                         " +
			"  /─\\                    " +
			"  \\─/                    " +
			"                         " +
			"          1234           " +
			"                         " +
			"                         " +
			"                         "

	fg1 :=
		"                         " +
			"                         " +
			"  ab                     " +
			"                         " +
			"                         " +
			"                         " +
			"                         " +
			"                         " +
			"                         "

	fg2 :=
		"                         " +
			"                         " +
			"  c                      " +
			"  d                      " +
			"                         " +
			"                         " +
			"                         " +
			"                         " +
			"                         "

	con.AddBackground(StringToArray(25, 9, bg1), redArray)
	con.AddBackground(StringToArray(25, 9, bg2), greenArray)
	con.AddForeground(StringToArray(25, 9, fg1), blueArray)
	con.AddForeground(StringToArray(25, 9, fg2), yellowArray)

	sprite := new(Sprite)
	sprite.Width = 2
	sprite.Height = 2
	sprite.X = 4
	sprite.Y = 3

	spriteChars := StringToArray(sprite.Width, sprite.Height, "SSS ")
	spriteColors := FillArrayI(sprite.Width, sprite.Height, defaultCharacterAttributes)
	sprite.SetGraphics(spriteChars, spriteColors)

	con.RegisterSprite(sprite, 0)

	sprite2 := new(Sprite)
	sprite2.Width = 4
	sprite2.Height = 1
	sprite2.X = 1
	sprite2.Y = 1
	sprite2.SetGraphics([][]rune{[]rune{'a', 'a', 'a', 'a'}}, [][]int{[]int{console.ChFgDarkBlue, defaultCharacterAttributes, defaultCharacterAttributes, defaultCharacterAttributes}})
	con.RegisterSprite(sprite2, 3)

	sprite3 := new(Sprite)
	sprite3.Width = 3
	sprite3.Height = 1
	sprite3.X = 2
	sprite3.Y = 1
	sprite3.SetGraphics([][]rune{[]rune{'b', 'b', 'b'}}, [][]int{[]int{console.ChFgDarkCyan, defaultCharacterAttributes, defaultCharacterAttributes}})
	con.RegisterSprite(sprite3, 2)

	sprite4 := new(Sprite)
	sprite4.Width = 2
	sprite4.Height = 1
	sprite4.X = 3
	sprite4.Y = 1
	sprite4.SetGraphics([][]rune{[]rune{'c', 'c'}}, [][]int{[]int{console.ChFgDarkGreen, defaultCharacterAttributes}})
	con.RegisterSprite(sprite4, 1)

	sprite5 := new(Sprite)
	sprite5.Width = 1
	sprite5.Height = 1
	sprite5.X = 4
	sprite5.Y = 1
	sprite5.SetGraphics([][]rune{[]rune{'d'}}, [][]int{[]int{console.ChFgDarkGrey}})
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
		[]int{console.ChFgRed, console.ChFgRed, console.ChFgRed, console.ChFgRed, console.ChFgRed},
		[]int{console.ChFgRed, console.ChFgDarkBlue, console.ChFgDarkCyan, console.ChFgDarkGreen, console.ChFgDarkGrey},
		[]int{console.ChFgRed, console.ChFgWhite, console.ChFgYellow, console.ChFgBlue, console.ChFgGreen},
		[]int{console.ChFgRed, console.ChFgWhite, console.ChFgYellow, console.ChFgGreen, console.ChFgWhite},
		[]int{console.ChFgRed, console.ChFgWhite, console.ChFgWhite, console.ChFgWhite, console.ChFgWhite},
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
