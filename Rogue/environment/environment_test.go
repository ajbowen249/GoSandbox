package environment

import(
"testing"
"github.com/ajbowen249/GoSandbox/Rogue/gameStructures"
)

type AddRoomTest struct{
	topLeft, bottomRight gameStructures.Point
	expectedBG string
}

func TestAddRoom(t *testing.T){
	cases := []AddRoomTest{
		{gameStructures.Point{5, 5}, gameStructures.Point{10, 10},
"                                                                                " +
"                                                                                " +
"                                                                                " +
"                                                                                " +
"                                                                                " +
"     ┌────┐                                                                     " +
"     │....│                                                                     " +
"     │....│                                                                     " +
"     │....│                                                                     " +
"     │....│                                                                     " +
"     └────┘                                                                     " +
"                                                                                " +
"                                                                                " +
"                                                                                " +
"                                                                                " +
"                                                                                " +
"                                                                                " +
"                                                                                " +
"                                                                                " +
"                                                                                " +
"                                                                                " +
"                                                                                " +
"                                                                                " +
"                                                                                " +
"                                                                                " },
	}
	
	

}