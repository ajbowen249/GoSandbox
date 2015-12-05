package rogueConsole

import 
(
"testing"
)

func TestConsole1(t *testing.T){
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
	sprite.SetString( 
		"SS" +
		"S ")
		
	con.AddSprite(4, 3, sprite)
	
	con.CameraX = 0
	con.CameraY = 0
	
	expectedBuffer :=
		"┌────" +
		"│    " +
		"│ cb\\" +
		"│ d─S" +
		"│   S"
		
	actualBuffer := con.GetFrameString()
	if actualBuffer != expectedBuffer{
		t.Errorf("Expected GetFrameString() == \n%v\n but was\n%v", expectedBuffer, actualBuffer)
	}
}