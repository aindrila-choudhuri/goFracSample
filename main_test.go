package main

import "testing"

func TestTableMultiplyFractions(t *testing.T) {
	var tests = []struct {
		inputFracOne string
		inputFracTwo string
		expected     string
	}{
		{"3/4", "5/6", "5/8"},
		{"56/78", "23/67", "644/2613"},
	}

	for _, test := range tests {
		if output := multiplyFractions(test.inputFracOne, test.inputFracTwo); output != test.expected {
			t.Error("Test failed: {} {} inputted, {} expected, received: {}", test.inputFracOne, test.inputFracTwo, test.expected, output)
		}
	}
}
