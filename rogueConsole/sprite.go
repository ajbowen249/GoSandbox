package rogueConsole

//Sprite is a movable collection of runes
//to represent a dynamic object on screen.
type Sprite struct {
	Width, Height, X, Y int
	runes               [][]rune
	colors              [][]int
}

// SetGraphics sets the characters and colors of the sprite.
func (sprite *Sprite) SetGraphics(runes [][]rune, colors [][]int) {
	sprite.runes = runes
	sprite.colors = colors
}

// GetGraphics returns the characters and colors of the sprite.
func (sprite *Sprite) GetGraphics() ([][]rune, [][]int) {
	return sprite.runes, sprite.colors
}
