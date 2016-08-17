package rogueConsole

// StringToArray transforms a string into a 2D slice of rune.
// Its intended use it to translate between the syntactical
// ease of multiline string constants to the engine requirement
// of a 2S slice.
func StringToArray(width int, height int, str string) [][]rune {
	array := make([][]rune, height)

	runes := []rune(str)
	strIndex := 0

	for row := 0; row < height; row++ {
		array[row] = make([]rune, width)
		for col := 0; col < width; col++ {
			array[row][col] = runes[strIndex]
			strIndex++
		}
	}

	return array
}

// ArrayToString flattens out a 2D rune array into a string.
func ArrayToString(array [][]rune) string {
	str := ""

	for row := 0; row < len(array); row++ {
		for col := 0; col < len(array[0]); col++ {
			str += string(array[row][col])
		}
	}

	return str
}

// FillArrayR creates a 2D slice filled with the given rune.
func FillArrayR(width int, height int, character rune) [][]rune {
	array := make([][]rune, height)

	for row := 0; row < height; row++ {
		array[row] = make([]rune, width)
		for col := 0; col < width; col++ {
			array[row][col] = character
		}
	}

	return array
}

// FillArrayI creates a 2D slice filled with the given integer.
func FillArrayI(width int, height int, integer int) [][]int {
	array := make([][]int, height)

	for row := 0; row < height; row++ {
		array[row] = make([]int, width)
		for col := 0; col < width; col++ {
			array[row][col] = integer
		}
	}

	return array
}
