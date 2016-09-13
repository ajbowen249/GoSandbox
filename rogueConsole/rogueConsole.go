package rogueConsole

import (
	"fmt"

	"github.com/ajbowen249/GoSandbox/console"
)

const defaultTransparencyColor = console.ChBgBlack

//TransparancyChar is the rune that flags a transparency
const TransparancyChar = rune(0x00)

//RogueConsole wraps up interaction with the system console to draw
//its given sprites, foregrounds, and backgrounds.
type RogueConsole struct {
	EnvWidth, EnvHeight, CameraWidth, CameraHeight, CameraX, CameraY int
	TransparencyColor                                                int

	bgLayers, fgLayers [][][]rune
	bgColors, fgColors [][][]int

	sprites [][]*Sprite
}

//NewRogueConsole sets up and returns a new RogueConsole instance.
func NewRogueConsole(envWidth int, envHeight int, cameraWidth int, cameraHeight int) *RogueConsole {
	con := new(RogueConsole)

	con.EnvWidth = envWidth
	con.EnvHeight = envHeight
	con.CameraWidth = cameraWidth
	con.CameraHeight = cameraHeight
	con.TransparencyColor = defaultTransparencyColor

	return con
}

//AddBackground adds a new background layer.
//backgrounds should be added from the bottom up.
func (con *RogueConsole) AddBackground(characters [][]rune, colors [][]int) {
	con.bgLayers = append(con.bgLayers, characters)
	con.bgColors = append(con.bgColors, colors)
}

//AddForeground adds a new foreground.
//Foreground should be added from the bottom up.
func (con *RogueConsole) AddForeground(characters [][]rune, colors [][]int) {
	con.fgLayers = append(con.fgLayers, characters)
	con.fgColors = append(con.fgColors, colors)
}

//RegisterSprite takes a pointer to a sprite to be included in the scene
//and a layer index to order the drawing.
func (con *RogueConsole) RegisterSprite(sp *Sprite, layer int) {
	con.expandSpriteLayers(layer)
	con.sprites[layer] = append(con.sprites[layer], sp)
}

//GetFrameArray returns a final buffer with all foregrounds, backgrounds,
//and sprites drawn.
func (con *RogueConsole) GetFrameArray() ([][]rune, [][]int) {
	frame := FillArrayR(con.CameraWidth, con.CameraHeight, TransparancyChar)
	frameColors := FillArrayI(con.CameraWidth, con.CameraHeight, con.TransparencyColor)

	for i := 0; i < len(con.bgLayers); i++ {
		grabWindow(con.CameraX, con.CameraY, con.CameraWidth, con.CameraHeight, &con.bgLayers[i], &con.bgColors[i], &frame, &frameColors)
	}

	spriteLayer := FillArrayR(con.EnvWidth, con.EnvHeight, TransparancyChar)
	spriteColors := FillArrayI(con.EnvWidth, con.EnvHeight, con.TransparencyColor)

	for layer := len(con.sprites) - 1; layer >= 0; layer-- {
		for sprite := 0; sprite < len(con.sprites[layer]); sprite++ {
			drawSprite(con.sprites[layer][sprite], &spriteLayer, &spriteColors)
		}
	}

	grabWindow(con.CameraX, con.CameraY, con.CameraWidth, con.CameraHeight, &spriteLayer, &spriteColors, &frame, &frameColors)

	for i := 0; i < len(con.fgLayers); i++ {
		grabWindow(con.CameraX, con.CameraY, con.CameraWidth, con.CameraHeight, &con.fgLayers[i], &con.fgColors[i], &frame, &frameColors)
	}

	Replace(&frame, TransparancyChar, ' ')

	return frame, frameColors
}

//Draw outputs the frame buffer to the console.
func (con *RogueConsole) Draw() {
	con.Visit(func(r rune, i int, row int, col int) {
		console.MoveTo(col, row)
		console.SetCharacterProperties(i)
		fmt.Print(string(r))
	})
}

// Visit passes each rune in the frame array through the
// given visitor function
func (con *RogueConsole) Visit(visit func(rune, int, int, int)) {
	frame, frameColors := con.GetFrameArray()

	for row := 0; row < len(frame); row++ {
		for col := 0; col < len(frame[row]); col++ {
			visit(frame[row][col], frameColors[row][col], row, col)
		}
	}
}

func grabWindow(x int, y int, width int, height int, charSource *[][]rune, colSource *[][]int, charDestination *[][]rune, colDestination *[][]int) {
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			character := (*charSource)[row+y][col+x]

			if character != TransparancyChar {
				(*charDestination)[row][col] = character
				(*colDestination)[row][col] = (*colSource)[row+y][col+x]
			}
		}
	}
}

func drawSprite(sprite *Sprite, charDestination *[][]rune, colDestination *[][]int) {
	charSource, colSource := sprite.GetGraphics()
	for row := 0; row < len(charSource); row++ {
		for col := 0; col < len(charSource[0]); col++ {
			character := charSource[row][col]
			charX := col + sprite.X
			charY := row + sprite.Y

			if character != TransparancyChar &&
				charX >= 0 &&
				charX < len((*charDestination)[0]) &&
				charY >= 0 &&
				charY < len(*charDestination) {
				(*charDestination)[charY][charX] = character
				(*colDestination)[charY][charX] = colSource[row][col]
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
