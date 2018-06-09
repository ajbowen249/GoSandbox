// Package console provides basic console control
// for positioning the cursor and non-blocking input.
package console

// IMPROVE: There is a lot of hacking in this file
//          to get it working on posix systems.
//          Most of it still holds on to Windows
//          ideas. This will need heavy refactoring
//          to make it all cleanly cross-platform.

// #ifdef __WINDOWS__
// #include <stdio.h>
// #include <windows.h>
// #include <conio.h>
//
// const int IS_WINDOWS = 1;
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
// void SetCursorProperties(int visible)
// {
//     CONSOLE_CURSOR_INFO cursorInfo;
//     cursorInfo.dwSize = 1;
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
// int GetConAttributes(CONSOLE_SCREEN_BUFFER_INFO* infoBuffer)
// {
//     return GetConsoleAttributes(GetStdHandle(STD_OUTPUT_HANDLE), infoBuffer);
// }
// #else
// #include <stdio.h>
// #include <termios.h>
// #include <unistd.h>
// #include <fcntl.h>
// const int IS_WINDOWS = 0;
//
// // Stole this from here: https://cboard.cprogramming.com/c-programming/63166-kbhit-linux.html
// int kbhit()
// {
//     struct termios oldt, newt;
//     int ch;
//     int oldf;
//
//     tcgetattr(STDIN_FILENO, &oldt);
//     newt = oldt;
//     newt.c_lflag &= ~(ICANON | ECHO);
//     tcsetattr(STDIN_FILENO, TCSANOW, &newt);
//     oldf = fcntl(STDIN_FILENO, F_GETFL, 0);
//     fcntl(STDIN_FILENO, F_SETFL, oldf | O_NONBLOCK);
//
//     ch = getchar();
//
//     tcsetattr(STDIN_FILENO, TCSANOW, &oldt);
//     fcntl(STDIN_FILENO, F_SETFL, oldf);
//
//     if(ch != EOF)
//     {
//         ungetc(ch, stdin);
//         return 1;
//     }
//
//     return 0;
// }
//
// void MoveTo(int row, int column)
// {
//     // terminal codes are 1-indexed.
//     row++;
//     column++;
//     printf("\e[%i;%iH", column, row);
//     fflush(stdout);
// }
//
// char GetKey()
// {
//     char character = 0x00;
//     if(kbhit())
//     {
//        character = getchar();
//        if ( character == 0x1b) {
//            // Make this work like Windows
//            if(kbhit())
//            {
//                getchar(); // skip the '['
//                return 0xE0; // return the windows special character identifier
//            }
//        }
//     }
//
//     return character;
// }
//
//
// void SetCursorProperties(int visible)
// {
//     printf("\e[?25");
//     printf(visible ? "h" : "l");
//     fflush(stdout);
// }
// /*
//     Hack here to translate from the Windows interface.
//     Integer structure is:
//     xxxxxxxxxxxxxxxxxxxxxxx|        x|           x|           x|                 xxx|                 xxx|
//           reserved (blink?)|underline|bg intensity|fg intensity|background color BGR|foreground color BGR|
// */
// const int FOREGROUND_RED = 0x00000001;
// const int FOREGROUND_GREEN = 0x00000002;
// const int FOREGROUND_BLUE = 0x00000004;
// const int BACKGROUND_RED = 0x00000008;
// const int BACKGROUND_GREEN = 0x00000010;
// const int BACKGROUND_BLUE = 0x00000020;
// const int FOREGROUND_INTENSITY = 0x00000040;
// const int BACKGROUND_INTENSITY = 0x00000080;
// const int COMMON_LVB_UNDERSCORE = 0x00000100;
//
// void SetCharacterProperties(int properties)
// {
//     int foreground = (properties & 0x00000007) + 30;
//     if ( properties & FOREGROUND_INTENSITY ) {
//         foreground += 60;
//     }
//
//     int background = ((properties & 0x00000038) >> 3) + 40;
//     if ( properties & BACKGROUND_INTENSITY ) {
//         background += 60;
//     }
//     printf("\e[%i;%im", foreground, background);
//     fflush(stdout);
// }
//
// #endif // __WINDOWS__
import "C"
import (
	"fmt"
)

// Attributes contains data about the current console window.
type Attributes struct {
	CharacterColor int
	ShowCursor     bool
	Echo           bool
}

// KeyboardInputInfo contains data about a single keyboard event.
// It could represent either a character key or a special code.
// The codes are wrapped up in the Sc* constants in thie package.
type KeyboardInputInfo struct {
	KeyDown, IsSpecial bool
	Char               rune
	SpecialChar        byte
}

// MoveTo sets the console cursor postition
func MoveTo(column int, row int) {
	C.MoveTo(C.int(column), C.int(row))
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
	value := KeyboardInputInfo{false, false, ' ', 0x00}

	fromC := C.GetKey()
	fromB := byte(fromC)

	if fromB == 0x00 {
		return false, value
	}

	value.KeyDown = true

	if fromB == ScEsc {
		value.IsSpecial = true
		value.SpecialChar = ScEsc
		return true, value
	}

	if fromB == ScIdentifier {
		value.IsSpecial = true
		value.SpecialChar = byte(C.GetKey())
		// HACK: translate from posix
		if C.IS_WINDOWS == 0 {
			switch value.SpecialChar {
			case ScArrowUpPosix:
				value.SpecialChar = ScArrowUp
			case ScArrowLeftPosix:
				value.SpecialChar = ScArrowLeft
			case ScArrowRightPosix:
				value.SpecialChar = ScArrowRight
			case ScArrowDownPosix:
				value.SpecialChar = ScArrowDown
			}
		}

		return true, value
	}

	value.Char = rune(fromC)

	return true, value
}

// SetCursorProperties sets the visibility of the cursor.
// The first argument is the percentage of the cell filled
// by the cursor, from 1 to 100. The second argument sets
// whether the cursor is visible.
func SetCursorProperties(isVisible bool) {
	visible := 0
	if isVisible {
		visible = 1
	}

	C.SetCursorProperties(C.int(visible))
}

// SetCharacterProperties sets the various properties
// (foreground and background color, etc.) of the
// output to the console. Use by or-ing together
// the Ch* constants in this package.
func SetCharacterProperties(properties int) {
	C.SetCharacterProperties(C.int(properties))
}

// GetDefaultAttributes returns data abound the current console window.
func GetDefaultAttributes() (bool, Attributes) {
	// infoBuffer := new(C.CONSOLE_SCREEN_BUFFER_INFO)
	// if int(C.GetConAttributes(infoBuffer)) != 0 {
	// 	return true, Attributes{int(infoBuffer.wAttributes)}
	// }

	return true, Attributes{ChFgWhite | ChBgBlack, true, true}
}

// SetAttributes takes a Attributes and sets all the
// current window's properties to match.
func SetAttributes(info Attributes) {
	SetCharacterProperties(info.CharacterColor)
	SetCursorProperties(info.ShowCursor)
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

const (
	ScArrowUpPosix    = 0x41
	ScArrowLeftPosix  = 0x44
	ScArrowRightPosix = 0x43
	ScArrowDownPosix  = 0x42
)
