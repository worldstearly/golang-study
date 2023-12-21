// calc.go
package main

import "testing"

func add(num1 int, num2 int) int {
	return num1 + num2
}

func TestAdd(t *testing.T) {
	if ans := add(1, 2); ans != 3 {
		t.Error("add(1, 2) should be equal to 3")
	}
}
