// Package forms allows creation of console-oriented
// gui pages.
package forms

import (
	"fmt"

	"github.com/ajbowen249/GoSandbox/algorithms"
	rc "github.com/ajbowen249/GoSandbox/rogueConsole"
)

//Control is any visual element shown in a form.
type Control interface {
	GetName() string
	Focus() bool
	Unfocus()
	Process()
	InitVisual(*rc.RogueConsole)
}

//Form is a "screen" of an application that contains
//and manages Controls
type Form struct {
	Controls map[string]Control
	TabOrder []string

	currentTabIndex int
	visual          *rc.RogueConsole
}

//NewForm creates and initializes a form.
func NewForm(width int, height int) *Form {
	form := new(Form)
	form.Controls = make(map[string]Control)

	form.currentTabIndex = -1
	form.visual = rc.NewRogueConsole(width, height, width, height)

	return form
}

// AddControl adds a new control to the form.
// The new control must have a unique name.
func (form *Form) AddControl(control Control, autoAddTabOrder bool) {
	_, exits := form.Controls[control.GetName()]

	if exits {
		panic(fmt.Sprintf("Attempted to add duplicate control (%v)", control.GetName()))
	}

	form.Controls[control.GetName()] = control
	form.TabOrder = append(form.TabOrder, control.GetName())
}

// FocusNext moves focus from the current control
// to the next control specified in the default
// tab order.
func (form *Form) FocusNext() {
	form.unfocusCurrentControl()

	newIndex := form.currentTabIndex

	for true {
		newIndex = (newIndex + 1) % len(form.TabOrder)

		//we've looped all the way around and found no focusable controls
		if newIndex == form.currentTabIndex {
			form.currentTabIndex = -1
			return
		}

		if form.Controls[form.TabOrder[newIndex]].Focus() {
			form.currentTabIndex = newIndex
			return
		}
	}
}

// FocusSpecific moves focus to the control with the given name
// and sets the tab index to the index of that control.
func (form *Form) FocusSpecific(controlName string) {
	form.unfocusCurrentControl()

	if form.Controls[controlName].Focus() {
		form.currentTabIndex = algorithms.SearchSliceString(form.TabOrder, controlName)
	} else {
		form.currentTabIndex = -1
	}
}

// Process calls the Process method on all controls.
func (form *Form) Process() {
	form.forAllControls(func(control Control) {
		control.Process()
	})
}

// InitiVisual passes the form's visual context to
// the InitiVisual method of all controls.
func (form *Form) InitiVisual() {
	form.forAllControls(func(control Control) {
		control.InitVisual(form.visual)
	})
}

func (form *Form) unfocusCurrentControl() {
	if form.currentTabIndex != -1 {
		form.Controls[form.TabOrder[form.currentTabIndex]].Unfocus()
	}
}

func (form *Form) forAllControls(action func(Control)) {
	for _, control := range form.Controls {
		action(control)
	}
}
