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
	draw(rCon, vCon)
	
	for{
		isHit, char := vCon.GetKey()
		if isHit{
			
			char = strings.ToLower(char)
			switch char{
				case "w":
					rCon.CameraY--
				case "a":
					rCon.CameraX--
				case "s":
					rCon.CameraY++
				case "d":
					rCon.CameraX++
			}
			
			draw(rCon, vCon)
			vCon.MoveTo(vCon.NumCols - 1, vCon.NumRows - 1)
		}
	}
}

func setup() *rc.RogueConsole{
	con := rc.NewRogueConsole(25, 9, 5, 5)

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
	
	sprite := new(rc.Sprite)
	sprite.Width = 2
	sprite.Height = 2
	sprite.SetString( 
		"SS" +
		"S ")
		
	con.AddSprite(4, 3, sprite)
	
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