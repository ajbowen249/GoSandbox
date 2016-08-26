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
// int GetConScreenBufferInfo(CONSOLE_SCREEN_BUFFER_INFO* infoBuffer)
// {
//     return GetConsoleScreenBufferInfo(GetStdHandle(STD_OUTPUT_HANDLE), infoBuffer);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

// ScreenBufferInfo contains data about the current console window.
type ScreenBufferInfo struct {
	CharacterColor int
}

// KeyboardInputInfo contains data about a single keyboard event.
// It could represent either a character key or a special code.
// The codes are wrapped up in the Sc* constants in thie package.
type KeyboardInputInfo struct {
	IsSpecial   bool
	Char        rune
	SpecialChar byte
}

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

// GetKeyEX returns a bool indicating whether a key
// was pressed on the keyboard and a KeyboardInputInfo
// for what pressed. It does not block
// for input.
func GetKeyEX() (bool, KeyboardInputInfo) {
	value := KeyboardInputInfo{false, ' ', 0x00}

	fromC := C.GetKey()
	fromB := byte(fromC)

	if fromB == 0x00 {
		return false, value
	}

	if fromB == ScEsc {
		value.IsSpecial = true
		value.SpecialChar = ScEsc
		return true, value
	}

	if fromB == ScIdentifier {
		value.IsSpecial = true
		value.SpecialChar = byte(C.GetKey())
		return true, value
	}

	value.Char = rune(fromC)

	return true, value
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

// GetScreenBufferInfo returns data abound the current console window.
func GetScreenBufferInfo() (bool, ScreenBufferInfo) {
	infoBuffer := new(C.CONSOLE_SCREEN_BUFFER_INFO)
	if int(C.GetConScreenBufferInfo(infoBuffer)) != 0 {
		return true, ScreenBufferInfo{int(infoBuffer.wAttributes)}
	}

	return false, ScreenBufferInfo{0}
}

// SetScreenBufferInfo takes a ScreenBufferInfo and sets all the
// current window's properties to match.
func SetScreenBufferInfo(info ScreenBufferInfo) {
	SetCharacterProperties(info.CharacterColor)
}

// SetTitle sets the title of the console
func SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.SetConsoleTitle((*C.CHAR)(cTitle))
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

const (
	ScArrowUp    = 0x48
	ScArrowLeft  = 0x4B
	ScArrowRight = 0x4D
	ScArrowDown  = 0x50
	ScEsc        = 0x1B
	ScIdentifier = 0xE0
)
