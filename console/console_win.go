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
import "C"
import "fmt"

type Console struct{
	NumCols, NumRows int
}

func Default() *Console{
	con := new(Console)
	con.NumCols = 80
	con.NumRows = 25
	
	return con
}

// MoveTo sets the console cursor postition
func (con *Console) MoveTo(row int, column int){
	C.MoveTo(C.short(row), C.short(column))
}

func (con *Console) ClearScreen(){
	numChars := con.NumCols * con.NumRows
	
	con.MoveTo(0, 0)
	
	for i :=0; i < numChars; i++{
		fmt.Print(" ")
	}
	
	con.MoveTo(0, 0)
}

func (con *Console) KetKey() (bool, int){
	isHit :=  C.kbhit() == 0
	
	var hit int
	if isHit{
		hit = C.GoInt(C.getch())
	} else {
		hit = 0
	}
	
	return isHit, hit
}