package table

type Table struct{
	HeaderAlign Align
	CellAlign Align
	
	columns []string
	data map[string]map[int]string
	numRows int
}

type Align int
const(
	RIGHT Align = iota
	LEFT
	CENTERLEFTBIAS
	CENTERRIGHTBIAS
)

func New() *Table{
	tb := new(Table)
	tb.data = make(map[string]map[int]string)
	tb.numRows = 0
	tb.HeaderAlign = LEFT
	tb.CellAlign = RIGHT
	
	return tb
}

func (tb *Table) AddColumn(name string){
	tb.columns = append(tb.columns, name)
	tb.data[name] = make(map[int]string)
}


func (tb *Table) AddRow(values ...string){
	for i := 0; i < len(values); i++{
		tb.data[tb.columns[i]][tb.numRows] = values[i]
	}
	
	tb.numRows++
}

func (tb *Table) Output(outFunc func(string)){
	outData := make([][]string, len(tb.columns))
	
	for i := 0; i < len(tb.columns); i++{
		name := tb.columns[i]
		width := tb.columnWidth(name)
		seperator := tb.horizontalSeperator(width)
		
		outData[i] = make([]string, tb.numRows + 4)
		
		outData[i][0] = seperator
		outData[i][1] = fillCell(name, width, tb.HeaderAlign)
		outData[i][2] = seperator
		
		for row := 0; row < tb.numRows; row++{
			outData[i][row + 3] = fillCell(tb.data[name][row], width, tb.CellAlign)
		}
		
		outData[i][tb.numRows + 3] = seperator
	}
	
	outFunc(makeRow("┌", "┬", "┐", outData, 0))
	outFunc(makeRow("│", "│", "│", outData, 1))
	outFunc(makeRow("├", "┼", "┤", outData, 2))
	
	for i := 0; i < tb.numRows; i++{
		outFunc(makeRow("│", "│", "│", outData, i + 3))
	}
	
	outFunc(makeRow("└", "┴", "┘", outData, tb.numRows + 3))
}

func (tb *Table) columnWidth(columnName string) int{
	width := len(columnName) //minimum width is the column header
	for i := 0; i < tb.numRows; i++{
		itemWidth := len(tb.data[columnName][i])
		if itemWidth > width{
			width = itemWidth
		}
	}
	
	return width
}

func (tb *Table) horizontalSeperator(width int) string{
	return strExtend(width, "─")
}

func strExtend(width int, char string) string{
	str := ""
	for i := 0; i < width; i++{
		str += char
	}
	
	return str
}

func fillCell(content string, width int, align Align) string{
	fullFillWidth, smallFill, bigFill := fillWidths(len(content), width)
	
	
	var cell string

	switch(align){
		case RIGHT:
			cell = strExtend(fullFillWidth, " ") + content
		case LEFT:
			cell = content + strExtend(fullFillWidth, " ")
		case CENTERLEFTBIAS:
			cell = strExtend(smallFill, " ") + content + strExtend(bigFill, " ")
		case CENTERRIGHTBIAS:
			cell = strExtend(bigFill, " ") + content + strExtend(smallFill, " ")
	}
	
	return cell
}

func makeRow(left string, seperator string, right string, data [][]string, row int) string{
	rowString := left
	for i := 0; i < len(data); i++{
		rowString += data[i][row]
		if i < len(data) - 1{
			rowString += seperator
		}
	}
	
	rowString += right
	return rowString	
}

func fillWidths(taken int, available int) (full, small, big int){
	full = available - taken
	half := full / 2
	
	if full % 2 == 0{
		big = half
		small = half
	} else {
		//relying on integer division here
		small = half
		big = half + 1
	}
	
	return
}