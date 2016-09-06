package forms

import (
	"fmt"
	"testing"
)

type TestControl struct {
	Name     string
	HasFocus bool
}

func (tc *TestControl) GetName() string {
	return tc.Name
}

func (tc *TestControl) Focus() bool {
	tc.HasFocus = true
	return true
}

func (tc *TestControl) Unfocus() {
	tc.HasFocus = false
}

func TestAutoFocusOrder(t *testing.T) {
	testForm := NewForm()
	numTestControls := 10
	testControls := make([]*TestControl, numTestControls)

	for i := 0; i < numTestControls; i++ {
		testControls[i] = &TestControl{fmt.Sprintf("tc%v", i), false}
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
