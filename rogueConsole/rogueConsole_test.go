package rogueConsole

import (
	"testing"
)

func TestConsole1(t *testing.T) {
	con := NewRogueConsole(25, 9, 5, 5)

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

	con.AddBackgroundS(bg1)
	con.AddBackgroundS(bg2)
	con.AddForegroundS(fg1)
	con.AddForegroundS(fg2)

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

	expectedBuffer :=
		"┌────" +
			"│abcd" +
			"│ cb\\" +
			"│ d─S" +
			"│   S"

	actualBuffer := con.GetFrameString()
	if actualBuffer != expectedBuffer {
		t.Errorf("Expected GetFrameString() == \n%v\n but was\n%v", expectedBuffer, actualBuffer)
	}
}
