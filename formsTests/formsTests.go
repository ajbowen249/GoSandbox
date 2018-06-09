package main

import (
	"github.com/ajbowen249/GoSandbox/console"
	"github.com/ajbowen249/GoSandbox/forms"
)

func main() {
	standardBuffer := console.GetStandardScreenBufferHandle()
	defer console.SetActiveScreenBuffer(standardBuffer)

	appBuffer := console.CreateNewScreenBuffer()
	console.SetActiveScreenBuffer(appBuffer)
	console.SetCursorPropertiesForBuffer(1, false, appBuffer)

	form := forms.NewForm(80, 25, console.ChBgDarkGrey)
	form.SetAlternateScreenBuffer(appBuffer)

	button1 := forms.NewButton("btn1")
	button1.SetText("Button")
	button1.SetXPadding(2)
	button1.SetYPadding(2)
	button1.SetX(1)
	button1.SetY(2)

	button2 := forms.NewButton("btn2")
	button2.SetText("Another Button")
	button2.SetXPadding(1)
	button2.SetYPadding(1)
	button2.SetX(20)
	button2.SetY(3)

	quit := false

	quitButton := forms.NewButton("btnQuit")
	quitButton.SetText("Exit")
	quitButton.SetXPadding(1)
	quitButton.SetYPadding(1)
	quitButton.SetX(72)
	quitButton.SetY(20)
	quitButton.SetExecute(func() {
		quit = true
	})

	form.AddControl(button1, true)
	form.AddControl(button2, true)
	form.AddControl(quitButton, true)
	form.InitiVisual()
	form.FlagFocusNext()

	for true {
		form.Process()
		if quit {
			break
		}
	}
}
