package main

import(
"fmt"
"strings"
rc "github.com/ajbowen249/GoSandbox/rogueConsole"
vc "github.com/ajbowen249/GoSandbox/console"
)

func main(){
	rCon := setup()
	vCon := vc.Default()
	vCon.ClearScreen()
	vCon.SetCursorProperties(1, false)
	draw(rCon, vCon)
	
	sprite := new(rc.Sprite)
	sprite.Width = 2
	sprite.Height = 2
	sprite.X = 4
	sprite.Y = 3
	sprite.SetString( 
		"SS" +
		"S ")
		
	rCon.RegisterSprite(sprite, 0)
	
	for{
		isHit, char := vCon.GetKey()
		if isHit{
			
			char = strings.ToLower(char)
			switch char{
				case "w":
					if sprite.Y > 0{
						sprite.Y--
					}
				case "a":
					if sprite.X > 0{
						sprite.X--
					}
				case "s":
					if sprite.Y < rCon.EnvHeight{
						sprite.Y++
					}
				case "d":
					if sprite.X < rCon.EnvWidth{
						sprite.X++
					}
			}
			
			draw(rCon, vCon)
		}
	}
}

func setup() *rc.RogueConsole{
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

func draw(rCon *rc.RogueConsole, vCon * vc.Console){
	frame := rCon.GetFrameArray()
	
	for row := 0; row < len(frame); row++{
		for col:= 0; col < len(frame[row]); col++{
			vCon.MoveTo(col, row)
			fmt.Print(string(frame[row][col]))
		}
	}
}
