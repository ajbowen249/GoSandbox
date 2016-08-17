package rogueConsole

//Sprite is a movable collection of runes
//to represent a dynamic object on screen.
type Sprite struct {
	Width, Height, X, Y int
	str                 string

	array    [][]rune
	arraySet bool
}

//SetString takes a flat string containing
//the runes of the sprite. It will automatically
//be expanded out to the 2D slice format.
func (sp *Sprite) SetString(str string) {
	sp.str = str
	sp.arraySet = false
}

//GetString returnes the flat string containing
//the runes of the sprite.
func (sp *Sprite) GetString() string {
	return sp.str
}

//GetArray returns the expanded 2D slice containing
//the runes of the sprite.
func (sp *Sprite) GetArray() [][]rune {
	if !sp.arraySet {
		sp.array = StringToArray(sp.Width, sp.Height, sp.str)
		sp.arraySet = true
	}

	return sp.array
}
