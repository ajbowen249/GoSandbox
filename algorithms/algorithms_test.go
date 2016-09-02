package algorithms

import "testing"

type facTestCase struct {
	input, expected int64
}

func TestFac(t *testing.T) {
	cases := []facTestCase{
		{0, 1},
		{1, 1},
		{3, 6},
		{10, 3628800},
	}

	for _, c := range cases {
		result := Fac(c.input)
		if result != c.expected {
			t.Errorf("Expected Fac(%v) == %v, but was %v", c.input, c.expected, result)
		}
	}
}

type minIntTestCase struct {
	input                        []int
	expectedValue, expectedIndex int
}

func TestMinInt(t *testing.T) {
	cases := []minIntTestCase{
		{[]int{0, 1, 2, 3}, 0, 0},
		{[]int{255}, 255, 0},
		{[]int{10, 9, 8, 7, 6}, 6, 4},
		{[]int{10, 1, 9, 2, 8, 3, 7, 4, 6, 5}, 1, 1},
	}

	for _, c := range cases {
		val, pos := MinInt(c.input)
		if val != c.expectedValue {
			t.Errorf("Expected MinInt(%v) first return == %v, but was %v", c.input, c.expectedValue, val)
		}

		if pos != c.expectedIndex {
			t.Errorf("Expected MinInt(%v) second return == %v, but was %v", c.input, c.expectedIndex, pos)
		}
	}
}

type sortIntCase struct {
	input, expected []int
}

func TestSortInt(t *testing.T) {
	cases := []sortIntCase{
		{[]int{2, 1}, []int{1, 2}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{2, 1, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	for _, c := range cases {
		buf := SortInt(c.input)

		for i := 0; i < len(buf); i++ {
			if buf[i] != c.expected[i] {
				t.Errorf("Expected buf[%v] == %v, but was %v", i, c.expected[i], buf[i])
			}
		}
	}
}

type stringSearchCase struct {
	value         string
	expectedIndex int
}

func TestSearchSliceString(t *testing.T) {
	cases := []stringSearchCase{
		{"bob", 2},
		{"frank", 0},
		{"jim", 1},
		{"nobody", -1},
		{"fred", 3},
	}

	testArray := []string{
		"frank",
		"jim",
		"bob",
		"fred",
		"bob", //extra bob, should not return this index
	}

	for _, c := range cases {
		actualIndex := SearchSliceString(testArray, c.value)

		if actualIndex != c.expectedIndex {
			t.Errorf("expected index of %v == %v, but was %v.", c.value, c.expectedIndex, actualIndex)
		}
	}
}
