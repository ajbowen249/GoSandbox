package main

import (
	"github.com/ajbowen249/GoSandbox/console"
	rc "github.com/ajbowen249/GoSandbox/rogueConsole"
)

func main() {
	console.SaveInitialScreenState()
	defer console.RestoreInitialScreenState()
	console.SetNoEcho()

	rCon := setup()
	console.SetCursorProperties(false)
	console.ClearScreen(80, 25)

	sprite := new(rc.Sprite)
	sprite.X = 4
	sprite.Y = 3

	spriteChars := rc.StringToArray(2, 2, "SSS ")
	rc.Replace(&spriteChars, ' ', rc.TransparancyChar)
	spriteColors := rc.FillArrayI(2, 2, console.ChFgCyan)
	spriteColors[0][1] = console.ChFgGreen | console.ChBgDarkMagenta
	sprite.SetGraphics(spriteChars, spriteColors)

	rCon.RegisterSprite(sprite, 0)

	rCon.Draw()

	goUp := func() {
		if sprite.Y > 0 {
			sprite.Y--
		}
	}

	goDown := func() {
		if sprite.Y < rCon.EnvHeight {
			sprite.Y++
		}
	}

	goLeft := func() {
		if sprite.X > 0 {
			sprite.X--
		}
	}

	goRight := func() {
		if sprite.X < rCon.EnvWidth {
			sprite.X++
		}
	}

	for {
		isHit, keyInfo := console.GetKey()
			if isHit {
				if !keyInfo.IsSpecial {
				switch keyInfo.Char {
				case 'w': goUp()
				case 'a': goLeft()
				case 's': goDown()
				case 'd': goRight()
				case 'q':
					return
				}

			} else {
				switch keyInfo.SpecialChar {
				case console.ScArrowUp: goUp()
				case console.ScArrowLeft: goLeft()
				case console.ScArrowDown: goDown()
				case console.ScArrowRight: goRight()
				case console.ScEsc:
					return
				}
			}

			rCon.Draw()
		}
	}
}

func setup() *rc.RogueConsole {
	con := rc.NewRogueConsole(25, 9, 25, 9)

	redArray := rc.FillArrayI(25, 9, console.ChFgRed)
	greenArray := rc.FillArrayI(25, 9, console.ChFgGreen)
	blueArray := rc.FillArrayI(25, 9, console.ChFgBlue)
	yellowArray := rc.FillArrayI(25, 9, console.ChFgYellow)

	bg1 := rc.StringToArray(25, 9,
		"┌───────────────────────┐"+
			"│                       │"+
			"│                       │"+
			"│                       │"+
			"│                       │"+
			"│                       │"+
			"│                       │"+
			"│                       │"+
			"└───────────────────────┘")

	bg2 := rc.StringToArray(25, 9,
		"                         "+
			"                         "+
			"  /─\\                    "+
			"  \\─/                    "+
			"                         "+
			"          1234           "+
			"                         "+
			"                         "+
			"                         ")

	fg1 := rc.StringToArray(25, 9,
		"                         "+
			"                         "+
			"  ab                     "+
			"                         "+
			"                         "+
			"                         "+
			"                         "+
			"                         "+
			"                         ")

	fg2 := rc.StringToArray(25, 9,
		"                         "+
			"                         "+
			"  c                      "+
			"  d                      "+
			"                         "+
			"                         "+
			"                         "+
			"                         "+
			"                         ")

	rc.Replace(&bg1, ' ', rc.TransparancyChar)
	rc.Replace(&bg2, ' ', rc.TransparancyChar)
	rc.Replace(&fg1, ' ', rc.TransparancyChar)
	rc.Replace(&fg2, ' ', rc.TransparancyChar)

	con.AddBackground(bg1, redArray)
	con.AddBackground(bg2, greenArray)
	con.AddForeground(fg1, blueArray)
	con.AddForeground(fg2, yellowArray)

	con.CameraX = 0
	con.CameraY = 0

	return con
}
