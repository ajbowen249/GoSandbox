package main

import (
	"strings"

	"github.com/ajbowen249/GoSandbox/console"
	rc "github.com/ajbowen249/GoSandbox/rogueConsole"
)

func main() {
	rCon := setup()
	console.SetCursorProperties(1, false)
	console.ClearScreen(80, 25)

	sprite := new(rc.Sprite)
	sprite.Width = 2
	sprite.Height = 2
	sprite.X = 4
	sprite.Y = 3
	sprite.SetString(
		"SS" +
			"S ")

	rCon.RegisterSprite(sprite, 0)

	rCon.Draw()

	for {
		isHit, char := console.GetKey()
		if isHit {

			char = strings.ToLower(char)
			switch char {
			case "w":
				if sprite.Y > 0 {
					sprite.Y--
				}
			case "a":
				if sprite.X > 0 {
					sprite.X--
				}
			case "s":
				if sprite.Y < rCon.EnvHeight {
					sprite.Y++
				}
			case "d":
				if sprite.X < rCon.EnvWidth {
					sprite.X++
				}
			case "q":
				return
			}

			rCon.Draw()
		}
	}
}

func setup() *rc.RogueConsole {
	con := rc.NewRogueConsole(25, 9, 25, 9)

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

	con.CameraX = 0
	con.CameraY = 0

	return con
}
