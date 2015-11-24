package table

import "testing"

func TestTable1(t *testing.T){
	expectedTable := []string{
		"┌─────┬──────┐",
		"│XVal │YVal  │",
		"├─────┼──────┤",
		"│    1│     2│",
		"│    5│    10│",
		"│12345│     2│",
		"│    1│123456│",
		"└─────┴──────┘",
	}
	
	tb := New()
	tb.AddColumn("XVal")
	tb.AddColumn("YVal")
	
	tb.AddRow("1", "2")
	tb.AddRow("5", "10")
	tb.AddRow("12345", "2")
	tb.AddRow("1", "123456")
	
	resultBuffer := make([]string, len(expectedTable))
	resultRow := 0
	
	tb.Output(func (row string){
		t.Log(row)
		resultBuffer[resultRow] = row
		resultRow++
	})
	
	for i := 0; i < len(expectedTable); i++{
		if resultBuffer[i] != expectedTable[i]{
			t.Errorf("Expected row %v == %v, but was %v", i, expectedTable[i], resultBuffer[i])
		}
	}
}

func TestTable2(t *testing.T){
	expectedTable := []string{
		"┌────┬──────┬───────┐",
		"│XVal│YVal  │ZVal   │",
		"├────┼──────┼───────┤",
		"│   1│123456│1234567│",
		"└────┴──────┴───────┘",
	}
	
	tb := New()
	tb.AddColumn("XVal")
	tb.AddColumn("YVal")
	tb.AddColumn("ZVal")
	tb.AddRow("1", "123456", "1234567")
	
	resultBuffer := make([]string, len(expectedTable))
	resultRow := 0
	
	tb.Output(func (row string){
		t.Log(row)
		resultBuffer[resultRow] = row
		resultRow++
	})
	
	for i := 0; i < len(expectedTable); i++{
		if resultBuffer[i] != expectedTable[i]{
			t.Errorf("Expected row %v == %v, but was %v", i, expectedTable[i], resultBuffer[i])
		}
	}
}

func TestCenterRight1(t *testing.T){
	expectedTable := []string{
		"┌───────┐",
		"│XValues│",
		"├───────┤",
		"│  1111 │",
		"└───────┘",
	}
	
	tb := New()
	tb.CellAlign = CENTERRIGHTBIAS
	tb.AddColumn("XValues")
	tb.AddRow("1111")
	
	resultBuffer := make([]string, len(expectedTable))
	resultRow := 0
	
	tb.Output(func (row string){
		t.Log(row)
		resultBuffer[resultRow] = row
		resultRow++
	})
	
	for i := 0; i < len(expectedTable); i++{
		if resultBuffer[i] != expectedTable[i]{
			t.Errorf("Expected row %v == %v, but was %v", i, expectedTable[i], resultBuffer[i])
		}
	}
}

func TestCenterRight2(t *testing.T){
	expectedTable := []string{
		"┌─────────┐",
		"│X--Values│",
		"├─────────┤",
		"│   1111  │",
		"└─────────┘",
	}
	
	tb := New()
	tb.CellAlign = CENTERRIGHTBIAS
	tb.AddColumn("X--Values")
	tb.AddRow("1111")
	
	resultBuffer := make([]string, len(expectedTable))
	resultRow := 0
	
	tb.Output(func (row string){
		t.Log(row)
		resultBuffer[resultRow] = row
		resultRow++
	})
	
	for i := 0; i < len(expectedTable); i++{
		if resultBuffer[i] != expectedTable[i]{
			t.Errorf("Expected row %v == %v, but was %v", i, expectedTable[i], resultBuffer[i])
		}
	}
}

func TestCenterLeft1(t *testing.T){
	expectedTable := []string{
		"┌───────┐",
		"│XValues│",
		"├───────┤",
		"│ 1111  │",
		"└───────┘",
	}
	
	tb := New()
	tb.CellAlign = CENTERLEFTBIAS
	tb.AddColumn("XValues")
	tb.AddRow("1111")
	
	resultBuffer := make([]string, len(expectedTable))
	resultRow := 0
	
	tb.Output(func (row string){
		t.Log(row)
		resultBuffer[resultRow] = row
		resultRow++
	})
	
	for i := 0; i < len(expectedTable); i++{
		if resultBuffer[i] != expectedTable[i]{
			t.Errorf("Expected row %v == %v, but was %v", i, expectedTable[i], resultBuffer[i])
		}
	}
}

func TestCenterLeft2(t *testing.T){
	expectedTable := []string{
		"┌─────────┐",
		"│X--Values│",
		"├─────────┤",
		"│  1111   │",
		"└─────────┘",
	}
	
	tb := New()
	tb.CellAlign = CENTERLEFTBIAS
	tb.AddColumn("X--Values")
	tb.AddRow("1111")
	
	resultBuffer := make([]string, len(expectedTable))
	resultRow := 0
	
	tb.Output(func (row string){
		t.Log(row)
		resultBuffer[resultRow] = row
		resultRow++
	})
	
	for i := 0; i < len(expectedTable); i++{
		if resultBuffer[i] != expectedTable[i]{
			t.Errorf("Expected row %v == %v, but was %v", i, expectedTable[i], resultBuffer[i])
		}
	}
}

func TestCenterNoBias(t *testing.T){
	expectedTable := []string{
		"┌────────┐",
		"│X-Values│",
		"├────────┤",
		"│  1111  │",
		"└────────┘",
	}
	
	tb := New()
	tb.CellAlign = CENTERRIGHTBIAS
	tb.AddColumn("X-Values")
	tb.AddRow("1111")
	
	resultBuffer := make([]string, len(expectedTable))
	resultRow := 0
	
	tb.Output(func (row string){
		t.Log(row)
		resultBuffer[resultRow] = row
		resultRow++
	})
	
	for i := 0; i < len(expectedTable); i++{
		if resultBuffer[i] != expectedTable[i]{
			t.Errorf("Expected row %v == %v, but was %v", i, expectedTable[i], resultBuffer[i])
		}
	}
	
	tb.CellAlign = CENTERLEFTBIAS
	
	resultBuffer = make([]string, len(expectedTable))
	resultRow = 0
	
	tb.Output(func (row string){
		t.Log(row)
		resultBuffer[resultRow] = row
		resultRow++
	})
	
	for i := 0; i < len(expectedTable); i++{
		if resultBuffer[i] != expectedTable[i]{
			t.Errorf("Expected row %v == %v, but was %v", i, expectedTable[i], resultBuffer[i])
		}
	}
}