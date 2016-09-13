package main

import (
	"github.com/ajbowen249/GoSandbox/console"
	"github.com/ajbowen249/GoSandbox/forms"
)

func main() {
	gotInfo, initScreenInfo := console.GetScreenBufferInfo()
	if gotInfo {
		defer console.SetScreenBufferInfo(initScreenInfo)
	}

	form := forms.NewForm(80, 25, console.ChBgDarkGrey)

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

	form.AddControl(button1, true)
	form.AddControl(button2, true)
	form.InitiVisual()
	form.FocusNext()

	for true {
		form.Process()
	}
}
