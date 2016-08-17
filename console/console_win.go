// Package console provides basic console control
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
//
// void SetCharacterProperties(int properties)
// {
//     SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE), (WORD)properties);
// }
//
// int GetCharacterProperties()
// {
//     //TODO: This api call can fail. Not sure what to do about that yet.
//     CONSOLE_SCREEN_BUFFER_INFO infoBuffer;
//     GetConsoleScreenBufferInfo(GetStdHandle(STD_OUTPUT_HANDLE), &infoBuffer);
//     return infoBuffer.wAttributes;
// }
import "C"
import "fmt"

// MoveTo sets the console cursor postition
func MoveTo(column int, row int) {
	C.MoveTo(C.SHORT(column), C.SHORT(row))
}

// ClearScreen blanks out the console window
// and returns the cursor to {0, 0}
func ClearScreen(numCols int, numRows int) {
	numChars := numCols * numRows
	MoveTo(0, 0)

	for i := 0; i < numChars; i++ {
		fmt.Print(" ")
	}

	MoveTo(0, 0)
}

// GetKey returns a bool indicating whether a key
// was pressed on the keyboard and a string for
// which character was pressed. It does not block
// for input.
func GetKey() (bool, string) {
	character := ""
	isHit := false
	fromC := C.GetKey()

	if byte(fromC) != 0x00 {
		character = C.GoStringN(&fromC, 1)
		isHit = true
	}

	return isHit, character
}

// SetCursorProperties sets the visibility of the cursor.
// The first argument is the percentage of the cell filled
// by the cursor, from 1 to 100. The second argument sets
// whether the cursor is visible.
func SetCursorProperties(fillPercent int, isVisible bool) {
	visible := 0
	if isVisible {
		visible = 1
	}

	if fillPercent < 1 {
		fillPercent = 1
	}
	if fillPercent > 100 {
		fillPercent = 100
	}

	C.SetCursorProperties(C.int(fillPercent), C.int(visible))
}

// SetCharacterProperties sets the various properties
// (foreground and background color, etc.) of the
// output to the console. Use by or-ing together
// the Ch* constants in this package.
func SetCharacterProperties(properties int) {
	C.SetCharacterProperties(C.int(properties))
}

// GetCharacterProperties returns the flags currently
// set for the console character properties. It can be
// used to save the state of the console before
// altering it.
func GetCharacterProperties() int {
	return int(C.GetCharacterProperties())
}

const (
	ChFgBlack       = 0
	ChFgDarkBlue    = C.FOREGROUND_BLUE
	ChFgDarkGreen   = C.FOREGROUND_GREEN
	ChFgDarkCyan    = C.FOREGROUND_GREEN | C.FOREGROUND_BLUE
	ChFgDarkRed     = C.FOREGROUND_RED
	ChFgDarkMagenta = C.FOREGROUND_RED | C.FOREGROUND_BLUE
	ChFgDarkYellow  = C.FOREGROUND_RED | C.FOREGROUND_GREEN
	ChFgDarkGrey    = C.FOREGROUND_RED | C.FOREGROUND_GREEN | C.FOREGROUND_BLUE
	ChFgGrey        = C.FOREGROUND_INTENSITY
	ChFgBlue        = C.FOREGROUND_INTENSITY | C.FOREGROUND_BLUE
	ChFgGreen       = C.FOREGROUND_INTENSITY | C.FOREGROUND_GREEN
	ChFgCyan        = C.FOREGROUND_INTENSITY | C.FOREGROUND_GREEN | C.FOREGROUND_BLUE
	ChFgRed         = C.FOREGROUND_INTENSITY | C.FOREGROUND_RED
	ChFgMagenta     = C.FOREGROUND_INTENSITY | C.FOREGROUND_RED | C.FOREGROUND_BLUE
	ChFgYellow      = C.FOREGROUND_INTENSITY | C.FOREGROUND_RED | C.FOREGROUND_GREEN
	ChFgWhite       = C.FOREGROUND_INTENSITY | C.FOREGROUND_RED | C.FOREGROUND_GREEN | C.FOREGROUND_BLUE

	ChBgBlack       = 0
	ChBgDarkBlue    = C.BACKGROUND_BLUE
	ChBgDarkGreen   = C.BACKGROUND_GREEN
	ChBgDarkCyan    = C.BACKGROUND_GREEN | C.BACKGROUND_BLUE
	ChBgDarkRed     = C.BACKGROUND_RED
	ChBgDarkMagenta = C.BACKGROUND_RED | C.BACKGROUND_BLUE
	ChBgDarkYellow  = C.BACKGROUND_RED | C.BACKGROUND_GREEN
	ChBgDarkGrey    = C.BACKGROUND_RED | C.BACKGROUND_GREEN | C.BACKGROUND_BLUE
	ChBgGrey        = C.BACKGROUND_INTENSITY
	ChBgBlue        = C.BACKGROUND_INTENSITY | C.BACKGROUND_BLUE
	ChBgGreen       = C.BACKGROUND_INTENSITY | C.BACKGROUND_GREEN
	ChBgCyan        = C.BACKGROUND_INTENSITY | C.BACKGROUND_GREEN | C.BACKGROUND_BLUE
	ChBgRed         = C.BACKGROUND_INTENSITY | C.BACKGROUND_RED
	ChBgMagenta     = C.BACKGROUND_INTENSITY | C.BACKGROUND_RED | C.BACKGROUND_BLUE
	ChBgYellow      = C.BACKGROUND_INTENSITY | C.BACKGROUND_RED | C.BACKGROUND_GREEN
	ChBgWhite       = C.BACKGROUND_INTENSITY | C.BACKGROUND_RED | C.BACKGROUND_GREEN | C.BACKGROUND_BLUE

	ChUnderline = C.COMMON_LVB_UNDERSCORE
)
