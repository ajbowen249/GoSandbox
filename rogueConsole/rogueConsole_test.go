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
	sprite.SetString(
		"SS" +
			"S ")

	con.RegisterSprite(sprite, 0)

	sprite2 := new(Sprite)
	sprite2.Width = 4
	sprite2.Height = 1
	sprite2.X = 1
	sprite2.Y = 1
	sprite2.SetString("aaaa")
	con.RegisterSprite(sprite2, 3)

	sprite3 := new(Sprite)
	sprite3.Width = 3
	sprite3.Height = 1
	sprite3.X = 2
	sprite3.Y = 1
	sprite3.SetString("bbb")
	con.RegisterSprite(sprite3, 2)

	sprite4 := new(Sprite)
	sprite4.Width = 2
	sprite4.Height = 1
	sprite4.X = 3
	sprite4.Y = 1
	sprite4.SetString("cc")
	con.RegisterSprite(sprite4, 1)

	sprite5 := new(Sprite)
	sprite5.Width = 1
	sprite5.Height = 1
	sprite5.X = 4
	sprite5.Y = 1
	sprite5.SetString("d")
	con.RegisterSprite(sprite5, 0)

	con.CameraX = 0
	con.CameraY = 0

	actualBuffer, actualColors := con.GetFrameArray()

	expectedString :=
		"┌────" +
			"│abcd" +
			"│ cb\\" +
			"│ d─S" +
			"│   S"

	actualString := ArrayToString(actualBuffer)

	if actualString != expectedString {
		t.Errorf("Expected GetFrameString() == \n%v\n but was\n%v", expectedString, actualString)
	}

	expectedColors := [][]int{
		[]int{console.ChFgRed, console.ChFgRed, console.ChFgRed, console.ChFgRed, console.ChFgRed},
		[]int{console.ChFgRed, console.ChFgWhite, console.ChFgWhite, console.ChFgWhite, console.ChFgWhite},
		[]int{console.ChFgRed, console.ChFgWhite, console.ChFgYellow, console.ChFgBlue, console.ChFgGreen},
		[]int{console.ChFgRed, console.ChFgWhite, console.ChFgYellow, console.ChFgGreen, console.ChFgWhite},
		[]int{console.ChFgRed, console.ChFgWhite, console.ChFgWhite, console.ChFgWhite, console.ChFgWhite},
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if actualColors[i][j] != expectedColors[i][j] {
				t.Errorf("Expected colors[%v][%v] == %v, but was %v.", i, j, expectedColors[i][j], actualColors[i][j])
			}
		}
	}
}
