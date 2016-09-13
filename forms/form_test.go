package forms

import (
	"fmt"
	"testing"

	"github.com/ajbowen249/GoSandbox/algorithms"
	rc "github.com/ajbowen249/GoSandbox/rogueConsole"
)

type TestControl struct {
	Name      string
	HasFocus  bool
	Focusable bool

	ProcessCallback func(string)
}

func NewTestControl(name string) *TestControl {
	return &TestControl{name, false, true, func(string) {}}
}

func (tc *TestControl) GetName() string {
	return tc.Name
}

func (tc *TestControl) Focus() bool {
	tc.HasFocus = tc.Focusable
	return tc.Focusable
}

func (tc *TestControl) Unfocus() {
	tc.HasFocus = false
}

func (tc *TestControl) Process(frameInfo *FrameInfo) {
	if tc.ProcessCallback != nil {
		tc.ProcessCallback(tc.Name)
	}
}

func (tc *TestControl) InitVisual(*rc.RogueConsole) {

}

func (tc *TestControl) SetOwner(*Form) {

}

func TestAutoFocusOrder(t *testing.T) {
	testForm := NewForm(80, 25)
	numTestControls := 10
	testControls := make([]*TestControl, numTestControls)

	for i := 0; i < numTestControls; i++ {
		testControls[i] = NewTestControl(fmt.Sprintf("tc%v", i))
		testForm.AddControl(testControls[i], true)
	}

	for i := 0; i < numTestControls; i++ {
		testForm.FocusNext()

		for j := 0; j < numTestControls; j++ {
			expectFocus := j == i
			if testControls[j].HasFocus != expectFocus {
				t.Errorf("expcted %v HasFocus == %v, but was %v.", testControls[i].GetName(), expectFocus, testControls[i].HasFocus)
			}
		}
	}
}

func TestSpecificFocus(t *testing.T) {
	testForm := NewForm(80, 25)

	tc1 := NewTestControl("tc1")
	tc2 := NewTestControl("tc2")

	testForm.AddControl(tc1, true)
	testForm.AddControl(tc2, true)

	testForm.FocusSpecific("tc2")

	if !tc2.HasFocus {
		t.Errorf("Expected that tc2 would have focus, and it did not.")
	}

	testForm.FocusSpecific("tc1")

	if !tc1.HasFocus {
		t.Errorf("Expected that tc1 would have focus, and it did not.")
	}
}

func TestUnfocusableItems(t *testing.T) {
	testForm := NewForm(80, 25)

	tc1 := NewTestControl("tc1")
	tc2 := NewTestControl("tc2")
	tc3 := NewTestControl("tc3")
	tc4 := NewTestControl("tc4")

	tc2.Focusable = false
	tc3.Focusable = false

	testForm.AddControl(tc1, true)
	testForm.AddControl(tc2, true)
	testForm.AddControl(tc3, true)
	testForm.AddControl(tc4, true)

	testForm.FocusNext()

	failMessage := "Expected that %v would%v have focus."

	if !tc1.HasFocus {
		t.Errorf(failMessage, "tc1", "")
	}

	testForm.FocusNext()

	if tc2.HasFocus {
		t.Errorf(failMessage, "tc2", "n't")
	}

	if tc3.HasFocus {
		t.Errorf(failMessage, "tc3", "n't")
	}

	if !tc4.HasFocus {
		t.Errorf(failMessage, "tc4", "")
	}
}

func TestProces(t *testing.T) {
	testForm := NewForm(80, 25)

	numControls := 10
	var processedControls []string

	for i := 0; i < numControls; i++ {
		testControl := NewTestControl(fmt.Sprintf("tc%v", i))
		testControl.ProcessCallback = func(name string) {
			processedControls = append(processedControls, name)
		}

		testForm.AddControl(testControl, true)
	}

	testForm.Process()

	//Note: process order is not defined
	for i := 0; i < numControls; i++ {
		name := fmt.Sprintf("tc%v", i)
		if algorithms.SearchSliceString(processedControls, name) < 0 {
			t.Errorf("Did not process %v.", name)
		}
	}
}
