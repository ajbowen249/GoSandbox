// Package console provides basic console control
// for positioning the cursor and non-blocking input.
package console
// #include "console.h"
import "C"
import (
	"fmt"
)

// KeyboardInputInfo contains data about a single keyboard event.
// It could represent either a character key or a special code.
// The codes are wrapped up in the Sc* constants in thie package.
type KeyboardInputInfo struct {
	KeyDown, IsSpecial bool
	Char               rune
	SpecialChar        byte
}

// SaveInitialScreenState saves initial properties of the console
// so they can be restored on shutdown.
func SaveInitialScreenState() {
	C.SaveInitialScreenState()
}

// RestoreInitialScreenState restores the initial console
// properties for cleanup.
func RestoreInitialScreenState() {
	C.RestoreInitialScreenState()
}

// SetNoEcho disables console echo (if necessary on
// platform).
func SetNoEcho() {
	C.SetNoEcho()
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
// was pressed on the keyboard and a KeyboardInputInfo
// for what pressed. It does not block
// for input.
// It is recommended that SetNoEcho()
// is called during program initialization if
// you plan to use this.
func GetKey() (bool, KeyboardInputInfo) {
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
