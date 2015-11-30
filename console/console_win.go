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

// MoveTo sets the console cursor postition
func MoveTo(row int, column int){
	C.MoveTo(C.short(row), C.short(column))
}

func ClearScreen(){
	MoveTo(0, 0)
	
	for i :=0; i < 2000; i++{
		fmt.Print(" ")
	}
	
	MoveTo(0, 0)
}