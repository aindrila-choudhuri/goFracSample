package main

import "testing"

func TestMultiplyFractions(t *testing.T) {
	if multiplyFractions([]int{2, 3, 4}, []int{6, 7, 9}) != "4/63" {
		t.Error("Expected 2/6 * 3/7 * 4/9 to equal 4/63")
	}
}
