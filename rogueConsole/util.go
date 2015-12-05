package rogueConsole

func stringToArray(width int, height int, str string) [][]rune{
	array := make([][]rune, width)
	
	runes := []rune(str)
	strIndex := 0
	
	for row := 0; row < height; row++{
		array[row] = make([]rune, width)
		for col := 0; col < width; col++{
			array[row][col] = runes[strIndex]
			strIndex++
		}
	}
	
	return array
}

func arrayToString(array [][]rune) string{
	str := ""

	for row := 0; row < len(array); row++{
		for col := 0; col < len(array[0]); col++{
			str += string(array[row][col])
		}
	}
	
	return str
}

func fillArray(width int, height int, character rune) [][]rune{
	array := make([][]rune, width)
	
	for row := 0; row < height; row++{
		array[row] = make([]rune, width)
		for col := 0; col < width; col++{
			array[row][col] = character
		}
	}
	
	return array
}