package main

import (
	"testing"
)

func testCalculate(t *testing.T) {
	if addTwo(2) != 4 {
		t.Error("Expected 2 + 2 equals to 4")
	}
}
