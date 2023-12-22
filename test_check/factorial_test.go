// factorial_test.go
package main

import "testing"

func TestFactorial(t *testing.T) {
	// Test cases
	testCases := []struct {
		input    int
		expected int
	}{
		{0, 1},        // 0! = 1
		{1, 1},        // 1! = 1
		{5, 120},      // 5! = 5*4*3*2*1 = 120
		{-1, -1},      // Error case: negative input
		{10, 3628800}, // 10! = 10*9*8*7*6*5*4*3*2*1 = 3628800
	}

	for _, testCase := range testCases {
		result := Factorial(testCase.input)
		if result != testCase.expected {
			t.Errorf("For input %d, expected %d, but got %d", testCase.input, testCase.expected, result)
		}
	}
}
