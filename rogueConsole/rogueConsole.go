package rogueConsole

import (
	"fmt"

	"github.com/ajbowen249/GoSandbox/console"
)

//RogueConsole wraps up interaction with the system console to draw
//its given sprites, foregrounds, and backgrounds.
type RogueConsole struct {
	EnvWidth, EnvHeight, CameraWidth, CameraHeight, CameraX, CameraY int

	bgLayers, fgLayers [][][]rune
	sprites            [][]*Sprite
}

//NewRogueConsole sets up and returns a new RogueConsole instance.
func NewRogueConsole(envWidth int, envHeight int, cameraWidth int, cameraHeight int) *RogueConsole {
	con := new(RogueConsole)

	con.EnvWidth = envWidth
	con.EnvHeight = envHeight
	con.CameraWidth = cameraWidth
	con.CameraHeight = cameraHeight

	return con
}

//AddBackgroundS takes a flat string to expand to a 2D background layer.
//backgrounds should be added from the bottom up.
func (con *RogueConsole) AddBackgroundS(layer string) {
	con.bgLayers = append(con.bgLayers, stringToArray(con.EnvWidth, con.EnvHeight, layer))
}

//AddForegroundS takes a flat string to expand to a 2D foreground layer.
//Foreground should be added from the bottom up.
func (con *RogueConsole) AddForegroundS(layer string) {
	con.fgLayers = append(con.fgLayers, stringToArray(con.EnvWidth, con.EnvHeight, layer))
}

//RegisterSprite takes a pointer to a sprite to be included in the scene
//and a layer index to order the drawing.
func (con *RogueConsole) RegisterSprite(sp *Sprite, layer int) {
	con.expandSpriteLayers(layer)
	con.sprites[layer] = append(con.sprites[layer], sp)
}

//GetFrameArray returns a final buffer with all foregrounds, backgrounds,
//and sprites drawn.
func (con *RogueConsole) GetFrameArray() [][]rune {
	frame := fillArray(con.CameraWidth, con.CameraHeight, ' ')

	for i := 0; i < len(con.bgLayers); i++ {
		grabWindow(con.CameraX, con.CameraY, con.CameraWidth, con.CameraHeight, &con.bgLayers[i], &frame)
	}

	spriteLayer := fillArray(con.EnvWidth, con.EnvHeight, ' ')
	for layer := len(con.sprites) - 1; layer >= 0; layer-- {
		for sprite := 0; sprite < len(con.sprites[layer]); sprite++ {
			drawSprite(con.sprites[layer][sprite].X, con.sprites[layer][sprite].Y, con.sprites[layer][sprite].GetArray(), &spriteLayer)
		}
	}

	grabWindow(con.CameraX, con.CameraY, con.CameraWidth, con.CameraHeight, &spriteLayer, &frame)

	for i := 0; i < len(con.fgLayers); i++ {
		grabWindow(con.CameraX, con.CameraY, con.CameraWidth, con.CameraHeight, &con.fgLayers[i], &frame)
	}

	return frame
}

//GetFrameString returns a final flattened string with all foregrounds,
//backgrounds, and sprites drawn.
func (con *RogueConsole) GetFrameString() string {
	return arrayToString(con.GetFrameArray())
}

//Draw outputs the frame buffer to the console.
func (con *RogueConsole) Draw() {
	frame := con.GetFrameArray()

	for row := 0; row < len(frame); row++ {
		for col := 0; col < len(frame[row]); col++ {
			console.MoveTo(col, row)
			fmt.Print(string(frame[row][col]))
		}
	}
}

func grabWindow(x int, y int, width int, height int, source *[][]rune, destination *[][]rune) {
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			character := (*source)[row+y][col+x]

			if character != ' ' {
				(*destination)[row][col] = character
			}
		}
	}
}

func drawSprite(x int, y int, source [][]rune, destination *[][]rune) {
	for row := 0; row < len(source); row++ {
		for col := 0; col < len(source[0]); col++ {
			character := source[row][col]
			charX := col + x
			charY := row + y

			if character != ' ' &&
				charX >= 0 &&
				charX < len((*destination)[0]) &&
				charY >= 0 &&
				charY < 9 {
				(*destination)[charY][charX] = character
			}
		}
	}
}

func (con *RogueConsole) expandSpriteLayers(maxIndex int) {
	if len(con.sprites) < maxIndex+1 {
		buffer := make([][]*Sprite, maxIndex+1)
		for i := 0; i < len(con.sprites); i++ {
			buffer[i] = con.sprites[i]
		}

		con.sprites = buffer
	}
}
