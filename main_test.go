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
	}

	for _, test := range tests {
		if output := multiplyFractions(test.inputFracs...); output != test.expected {
			t.Error("Test failed: {} inputted, {} expected, received: {}", test.inputFracs, test.expected, output)
		}
	}
}
