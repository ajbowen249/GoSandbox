// The console package provides basic console control 
// for positioning the cursor and non-blocking input.
package console

// #include <stdio.h>
// #include <windows.h>
// #include <conio.h>
//
// void MoveTo(SHORT row, SHORT column)
// {
//     COORD Cord;
//     Cord.X = row;
//     Cord.Y = column;
//     SetConsoleCursorPosition(GetStdHandle(STD_OUTPUT_HANDLE), Cord);
// }
//
// char GetKey()
// {
//     char character = 0x00;
//     if(kbhit())
//     {
//        character = getch();
//     }
//
//     return character;
// }
//
// void SetCursorProperties(int fill, int visible)
// {
//     CONSOLE_CURSOR_INFO cursorInfo;
//     cursorInfo.dwSize = (DWORD)fill;
//     cursorInfo.bVisible = (BOOL)visible;
//     
//     SetConsoleCursorInfo(GetStdHandle(STD_OUTPUT_HANDLE), &cursorInfo);
// }
import "C"
import "fmt"

// MoveTo sets the console cursor postition
func MoveTo(column int, row int){
	C.MoveTo(C.SHORT(column), C.SHORT(row))
}

// ClearScreen blanks out the console window 
// and returns the cursor to {0, 0}
func ClearScreen(numCols int, numRows int){
	numChars := numCols * numRows
	MoveTo(0, 0)
	
	for i :=0; i < numChars; i++{
		fmt.Print(" ")
	}
	
	MoveTo(0, 0)
}

// GetKey returns a bool indicating whether a key
// was pressed on the keyboard and a string for
// which character was pressed. It does not block
// for input.
func GetKey() (bool, string){
	character := ""
	isHit := false
	fromC := C.GetKey()
	
	if byte(fromC) != 0x00{
		character = C.GoStringN(&fromC, 1)
		isHit = true
	}
	
	return isHit, character
}

// SetCursorProperties sets the visibility of the cursor. 
// The first argument is the percentage of the cell filled
// by the cursor, from 1 to 100. The second argument sets 
// whether the cursor is visible.
func SetCursorProperties(fillPercent int, isVisible bool){
	visible := 0
	if isVisible{ visible = 1}
	
	if fillPercent < 1{fillPercent = 1}
	if fillPercent > 100{fillPercent = 100}
	
	C.SetCursorProperties(C.int(fillPercent), C.int(visible))
}
