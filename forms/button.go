package forms

import (
	"github.com/ajbowen249/GoSandbox/console"
	rc "github.com/ajbowen249/GoSandbox/rogueConsole"
)

// Button is a control that performs a specific action when activated.
type Button struct {
	name, text                                       string
	x, y, layer, xPadding, yPadding                  int
	enabledBgColor, focusedBGColor, enabledTextColor int
	disabledBgColor, disabledTextColor, shadowColor  int
	hasFocus, isEnabled, isPressed                   bool
	execute                                          func()
	sprite                                           *rc.Sprite
	owner                                            *Form
}

// NewButton creates a new button and sets
// all of its properties to their defaults
func NewButton(name string) *Button {
	btn := new(Button)
	btn.name = name
	btn.x = 0
	btn.y = 0
	btn.layer = 0
	btn.xPadding = 0
	btn.yPadding = 0
	btn.hasFocus = false
	btn.isEnabled = true
	btn.isPressed = false
	btn.execute = func() {}

	btn.enabledBgColor = console.ChBgBlue
	btn.enabledTextColor = console.ChFgBlack
	btn.focusedBGColor = console.ChBgRed
	btn.disabledBgColor = console.ChBgGrey
	btn.disabledTextColor = console.ChFgWhite
	btn.shadowColor = console.ChBgBlack

	btn.sprite = new(rc.Sprite)

	return btn
}

// SetText sets the text
func (btn *Button) SetText(newText string) {
	btn.text = newText
}

// SetXPadding sets the x padding
func (btn *Button) SetXPadding(newPadding int) {
	btn.xPadding = newPadding
}

// SetYPadding sets the y padding
func (btn *Button) SetYPadding(newPadding int) {
	btn.yPadding = newPadding
}

// SetX sets the x value
func (btn *Button) SetX(x int) {
	btn.x = x
}

// SetY sets the y value
func (btn *Button) SetY(y int) {
	btn.y = y
}

// GetName returns the name of the button.
func (btn *Button) GetName() string {
	return btn.name
}

// Focus sets the focus to the button if it is enabled.
func (btn *Button) Focus() bool {
	if btn.isEnabled {
		btn.hasFocus = true
		btn.draw()
	}

	return btn.hasFocus
}

// Unfocus removes focus from the button.
func (btn *Button) Unfocus() {
	btn.hasFocus = false
	btn.draw()
}

// Process gives the button time to handle user interaction.
func (btn *Button) Process(frameInfo *FrameInfo) {
	if btn.hasFocus {
		if frameInfo.KeyInfo.KeyDown {
			if !frameInfo.KeyInfo.IsSpecial {
				if !btn.isPressed && (frameInfo.KeyInfo.Char == '\r' || frameInfo.KeyInfo.Char == ' ') {
					btn.press()
				} else if frameInfo.KeyInfo.Char == '\t' {
					btn.unpress()
					btn.owner.FocusNext()
				}
			}
		} else {
			if btn.isPressed {
				btn.unpress()
				btn.execute()
			}
		}
	}
}

// InitVisual allows the button to register with its
// container's visual engine.
func (btn *Button) InitVisual(rCon *rc.RogueConsole) {
	rCon.RegisterSprite(btn.sprite, btn.layer)
	btn.draw()
}

// SetOwner sets the owner form
func (btn *Button) SetOwner(form *Form) {
	btn.owner = form
}

func (btn *Button) press() {
	btn.isPressed = true
	btn.draw()
}

func (btn *Button) unpress() {
	btn.isPressed = false
	btn.draw()
}

func (btn *Button) draw() {
	btn.sprite.X = btn.x
	btn.sprite.Y = btn.y

	rectWidth := len(btn.text) + (btn.xPadding * 2)
	rectHeight := (btn.yPadding * 2) + 1

	//+1 for the shadow
	rBuffer := rc.FillArrayR(rectWidth+1, rectHeight+1, rc.TransparancyChar)
	cBuffer := rc.FillArrayI(rectWidth+1, rectHeight+1, 0)

	//draw the shadow
	if !btn.isPressed {
		for i := 1; i <= rectWidth; i++ {
			rBuffer[rectHeight][i] = ' '
			cBuffer[rectHeight][i] = btn.shadowColor
		}

		for i := 1; i <= rectHeight; i++ {
			rBuffer[i][rectWidth] = ' '
			cBuffer[i][rectWidth] = btn.shadowColor
		}
	}

	//draw the box
	rectOffset := 0
	if btn.isPressed {
		rectOffset = 1
	}

	fgColor := btn.enabledTextColor
	bgColor := btn.enabledBgColor

	if btn.hasFocus {
		bgColor = btn.focusedBGColor
	}

	if !btn.isEnabled {
		fgColor = btn.disabledTextColor
		bgColor = btn.disabledBgColor
	}

	rectColor := fgColor | bgColor

	for row := 0; row < rectHeight; row++ {
		for col := 0; col < rectWidth; col++ {
			rBuffer[row+rectOffset][col+rectOffset] = ' '
			cBuffer[row+rectOffset][col+rectOffset] = rectColor
		}
	}

	textRow := btn.yPadding + rectOffset
	textIndent := btn.xPadding + rectOffset
	textSlice := []rune(btn.text)

	for i := 0; i < len(textSlice); i++ {
		rBuffer[textRow][textIndent+i] = textSlice[i]
	}

	btn.sprite.SetGraphics(rBuffer, cBuffer)
	btn.owner.InvalidateVisual()
}
