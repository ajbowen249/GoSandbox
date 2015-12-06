package rogueConsole

type RogueConsole struct{
	EnvWidth, EnvHeight, CameraWidth, CameraHeight, CameraX, CameraY int
	
	bgLayers, fgLayers [][][]rune
	sprites []*Sprite
}

func NewRogueConsole(envWidth int, envHeight int, cameraWidth int, cameraHeight int) *RogueConsole{
	con := new(RogueConsole)
	
	con.EnvWidth = envWidth
	con.EnvHeight = envHeight
	con.CameraWidth = cameraWidth
	con.CameraHeight = cameraHeight
	
	return con
}

func (con *RogueConsole)AddBackgroundS(layer string){
	con.bgLayers = append(con.bgLayers, stringToArray(con.EnvWidth, con.EnvHeight, layer))
}

func (con *RogueConsole)AddForegroundS(layer string){
	con.fgLayers = append(con.fgLayers, stringToArray(con.EnvWidth, con.EnvHeight, layer))
}

func (con *RogueConsole)RegisterSprite(sp *Sprite){
	con.sprites = append(con.sprites, sp)
}

func (con *RogueConsole)GetFrameArray() [][]rune{
	frame := fillArray(con.CameraWidth, con.CameraHeight, ' ')
	
	for i := 0; i < len(con.bgLayers); i++{
		grabWindow(con.CameraX, con.CameraY, con.CameraWidth, con.CameraHeight, &con.bgLayers[i], &frame)
	}
	
	spriteLayer := fillArray(con.EnvWidth, con.EnvHeight, ' ')
	for i := 0; i < len(con.sprites); i++{
		mergeArrays(con.sprites[i].X, con.sprites[i].Y, con.sprites[i].GetArray(), &spriteLayer)
	}
	
	grabWindow(con.CameraX, con.CameraY, con.CameraWidth, con.CameraHeight, &spriteLayer, &frame)
	
	for i := 0; i < len(con.fgLayers); i++{
		grabWindow(con.CameraX, con.CameraY, con.CameraWidth, con.CameraHeight, &con.fgLayers[i], &frame)
	}
	
	return frame
}

func (con *RogueConsole)GetFrameString() string{
	return arrayToString(con.GetFrameArray())
}

func grabWindow(x int, y int, width int, height int, source *[][]rune, destination *[][]rune){
	for row := 0; row < height; row++{
		for col := 0; col < width; col++{
			character := (*source)[row + y][col + x]
			
			if character != ' '{
				(*destination)[row][col] = character
			}
		}
	}
}

func mergeArrays(x int, y int, source [][]rune, destination *[][]rune){
	for row := 0; row < len(source); row++{
		for col := 0; col < len(source[0]); col++{
			character := source[row][col]
			if character != ' '{
				(*destination)[row + y][col + x] = character
			}
		}
	}
}