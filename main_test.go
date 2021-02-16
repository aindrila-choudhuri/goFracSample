package main

import "testing"

func TestTableMultiplyFractions(t *testing.T) {
	var tests = []struct {
		inputFracs []string
		expected   string
	}{
		{[]string{"3/4", "5/6"}, "5/8"},
		{[]string{"56/78", "23/67"}, "644/2613"},
		{[]string{"2/3", "5/8", "11/32"}, "55/384"},
		{[]string{"2/3", "3/4", "4/5", "5/6", "6/7", "7/8", "8/9", "9/10", "10/11", "2/5"}, "4/55"},
		{[]string{"3/5", "7/8", "2/6", "7/66", "12/13", "4/12", "39/60", "23/46", "6/13", "12/67"}, "147/958100"},
		{[]string{"3/5", "7/8", "2/6", "7/66", "12/13", "4/12", "39/60", "23/46", "6/13", "12/67", "2/9", "11/13"}, "49/1698450"},
		{[]string{"3/4", "5/0"}, undefinedString},
		{[]string{"3/4", "0/6"}, infiniteString},
		{[]string{"  3/4", "5/6"}, "5/8"},
	}
	//{[]string{"2/3", "3/4", "4/5", "5/6", "6/7", "7/8", "8/9", "9/10", "10/11", "2/5"}, "4/55"},
	for _, test := range tests {
		if output := multiplyFractions(test.inputFracs...); output != test.expected {
			t.Error("Test failed: {} inputted, {} expected, received: {}", test.inputFracs, test.expected, output)
		}
	}
}
