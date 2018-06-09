package forms

import (
	"github.com/ajbowen249/GoSandbox/console"
	rc "github.com/ajbowen249/GoSandbox/rogueConsole"
)

// TextBox is a control that performs a specific action when activated.
type TextBox struct {
	name, text                                               string
	x, y, layer, contentWidth                                int
	enabledBorderColor, focusedBorderColor, enabledTextColor int
	disabledBorderColor, disabledTextColor                   int
	hasFocus, isEnabled                                      bool
	sprite                                                   *rc.Sprite
	owner                                                    *Form
}

// NewTextBox creates a new TextBox and sets
// all of its properties to their defaults
func NewTextBox(name string) *TextBox {
	tb := new(TextBox)
	tb.name = name
	tb.x = 0
	tb.y = 0
	tb.layer = 0
	tb.contentWidth = 10
	tb.hasFocus = false
	tb.isEnabled = true

	tb.enabledBorderColor = console.ChFgBlack | console.ChBgGrey
	tb.enabledTextColor = console.ChFgBlack | console.ChBgWhite
	tb.focusedBorderColor = console.ChFgBlue | console.ChBgGrey
	tb.disabledBorderColor = console.ChFgWhite | console.ChBgDarkGrey
	tb.disabledTextColor = console.ChFgBlack | console.ChBgDarkGrey

	tb.sprite = new(rc.Sprite)

	return tb
}

// SetText sets the text
func (tb *TextBox) SetText(newText string) {
	tb.text = newText
}

// SetX sets the x value
func (tb *TextBox) SetX(x int) {
	tb.x = x
}

// SetY sets the y value
func (tb *TextBox) SetY(y int) {
	tb.y = y
}

// SetContentWidth sets the width of the display area
func (tb *TextBox) SetContentWidth(contentWidth int) {
	tb.contentWidth = contentWidth
}

// GetName returns the name of the TextBox.
func (tb *TextBox) GetName() string {
	return tb.name
}

// Focus sets the focus to the TextBox if it is enabled.
func (tb *TextBox) Focus() bool {
	if tb.isEnabled {
		tb.hasFocus = true
		tb.draw()
	}

	return tb.hasFocus
}

// Unfocus removes focus from the TextBox.
func (tb *TextBox) Unfocus() {
	tb.hasFocus = false
	tb.draw()
}

// Process gives the TextBox time to handle user interaction.
func (tb *TextBox) Process(frameInfo *FrameInfo) {
	if tb.hasFocus {
		if frameInfo.KeyInfo.KeyDown {
			if !frameInfo.KeyInfo.IsSpecial {
				if frameInfo.KeyInfo.Char == '\t' {
					tb.owner.FlagFocusNext()
				} else if byte(frameInfo.KeyInfo.Char) == console.ScDelete {
					if len(tb.text) > 0 {
						tb.text = string([]rune(tb.text)[0:len(tb.text) - 1])
						tb.draw()
					}
				} else {
					tb.text += string(frameInfo.KeyInfo.Char)
					tb.draw()
				} // TODO: backspace and arrow keys
			}
		}
	}
}

// InitVisual allows the TextBox to register with its
// container's visual engine.
func (tb *TextBox) InitVisual(rCon *rc.RogueConsole) {
	rCon.RegisterSprite(tb.sprite, tb.layer)
	tb.draw()
}

// SetOwner sets the owner form
func (tb *TextBox) SetOwner(form *Form) {
	tb.owner = form
}

func (tb *TextBox) draw() {
	tb.sprite.X = tb.x
	tb.sprite.Y = tb.y

	// +2 for the border
	rectWidth := tb.contentWidth + 2
	// line of text plus top and bottom border
	rectHeight := 3

	borderColor := tb.enabledBorderColor
	textColor := tb.enabledTextColor
	if tb.hasFocus {
		borderColor = tb.focusedBorderColor
	}

	if !tb.isEnabled {
		borderColor = tb.disabledBorderColor
		textColor = tb.disabledTextColor
	}

	rBuffer := rc.FillArrayR(rectWidth, rectHeight, ' ')
	cBuffer := rc.FillArrayI(rectWidth, rectHeight, borderColor)

	// draw the corners of the border
	rBuffer[0][0] = '┌'
	rBuffer[0][rectWidth-1] = '┐'
	rBuffer[rectHeight-1][0] = '└'
	rBuffer[rectHeight-1][rectWidth-1] = '┘'

	// fill in top and bottom of the border
	for i := 1; i < rectWidth-1; i++ {
		rBuffer[0][i] = '─'
		rBuffer[rectHeight-1][i] = '─'
	}

	// fill in left and right border
	for i := 1; i < rectHeight-1; i++ {
		rBuffer[i][0] = '│'
		rBuffer[i][rectWidth-1] = '│'
	}

	// fill in text
	runeSlice := []rune(tb.text)
	if tb.hasFocus && len(tb.text) > tb.contentWidth {
		runeSlice = runeSlice[len(tb.text) - tb.contentWidth : len(tb.text)]
	}

	for i := 0; i < rectWidth-2; i++ {
		cBuffer[1][i+1] = textColor
		if i < len(runeSlice) {
			rBuffer[1][i+1] = runeSlice[i]
		}
	}

	tb.sprite.SetGraphics(rBuffer, cBuffer)
	tb.owner.InvalidateVisual()
}
