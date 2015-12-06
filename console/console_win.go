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
import "C"
import "fmt"

type Console struct{
	NumCols, NumRows int
}

// Default returns a Console struct where
// NumCols == 80 and NumRows == 25
func Default() *Console{
	con := Console{80, 25}
	
	return &con
}

// MoveTo sets the console cursor postition
func (con *Console) MoveTo(column int, row int){
	C.MoveTo(C.short(column), C.short(row))
}

// ClearScreen blanks out the console window 
// and returns the cursor to {0, 0}
func (con *Console) ClearScreen(){
	numChars := con.NumCols * con.NumRows
	con.MoveTo(0, 0)
	
	for i :=0; i < numChars; i++{
		fmt.Print(" ")
	}
	
	con.MoveTo(0, 0)
}

// GetKey returns a bool indicating whether a key
// was pressed on the keyboard and a string for
// which character was pressed. It does not block
// for input.
func (con *Console) GetKey() (bool, string){
	character := ""
	isHit := false
	fromC := C.GetKey()
	
	if byte(fromC) != 0x00{
		character = C.GoStringN(&fromC, 1)
		isHit = true
	}
	
	return isHit, character
}