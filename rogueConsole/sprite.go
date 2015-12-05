package rogueConsole

type Sprite struct{
	Width, Height int
	str string
	
	array [][]rune
	arraySet bool
}

func (sp *Sprite)SetString(str string){
	sp.str = str
	sp.arraySet = false
}

func (sp *Sprite)GetString() string{
	return sp.str
}

func (sp *Sprite) GetArray() [][]rune{
	if !sp.arraySet{
		sp.array = stringToArray(sp.Width, sp.Height, sp.str)
		sp.arraySet = true
	}
	
	return sp.array
}