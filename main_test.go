package main

import "testing"

func TestTableMultiplyFractions(t *testing.T) {
	var tests = []struct {
		inputFracs []string
		expected   string
	}{
		//input with 2 fractions
		{[]string{"3/4", "5/6"}, "5/8"},
		//input with 12 fractions
		{[]string{"3/5", "7/8", "2/6", "7/66", "12/13", "4/12", "39/60", "23/46", "6/13", "12/67", "2/9", "11/13"}, "49/1698450"},
		//input with denominator 0
		{[]string{"3/4", "5/0"}, undefinedString},
		//input with numerator 0
		{[]string{"3/4", "0/6"}, infiniteString},
		//input with space
		{[]string{"  3/4", "5/6"}, "5/8"},
		//input with large number
		{[]string{"3/445", "3/445", "3/445"}, "27/88121125"},
	}

	for _, test := range tests {
		if output := multiplyFractions(test.inputFracs...); output != test.expected {
			t.Error("Test failed: {} inputted, {} expected, received: {}", test.inputFracs, test.expected, output)
		}
	}
}
