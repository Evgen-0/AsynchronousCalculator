package main

import (
	"bytes"
	"fmt"
	"testing"
)

// Unit tests for the code:
// 1. Test user input string with only one value.
// 2. Test user input string with more than three values.
// 3. Test user input string with invalid operator.
// 4. Test user input string with invalid number format.
// 5. Test user input string with correct format and values.
// 6. Test for the output of the "operate" function with different operators and values.

func Test_divide(t *testing.T) {
	testCases := []struct {
		num1     float64
		num2     float64
		expected float64
		err      error
	}{
		{10, 2, 5, nil},
		{-12, 4, -3, nil},
		{0, 4, 0, nil},
		{6, -2, -3, nil},
	}
	// Test cases for invalid input
	invalidTestCases := []struct {
		num1 float64
		num2 float64
		err  error
	}{
		{10, 0, fmt.Errorf("помилка: Ділення на нуль неможливе")},
		{-5, 0, fmt.Errorf("помилка: Ділення на нуль неможливе")},
	}
	// Test cases for valid input
	for _, tc := range testCases {
		result, err := divide(tc.num1, tc.num2)
		if err != tc.err {
			t.Errorf("Expected error: %v but got: %v", tc.err, err)
		}

		if result != tc.expected {
			t.Errorf("Expected result: %v but got: %v", tc.expected, result)
		}
	}
	// Test cases for invalid input
	for _, tc := range invalidTestCases {
		_, err := divide(tc.num1, tc.num2)
		if err.Error() != tc.err.Error() {
			t.Errorf("Expected error: %v but got: %v", tc.err, err)
		}
	}
}

func Test_operate(t *testing.T) {
	// Test cases for valid input
	testCases := []struct {
		operator string
		num1     float64
		num2     float64
		expected string
	}{
		{"+", 10, 2, "10.00 + 2.00 = 12.00\n"},
		{"-", -12, 4, "-12.00 - 4.00 = -16.00\n"},
		{"*", 0, 4, "0.00 * 4.00 = 0.00\n"},
		{"/", 6, -2, "6.00 / -2.00 = -3.00\n"},
	}
	// Test cases for invalid input
	invalidTestCases := []struct {
		operator string
		num1     float64
		num2     float64
		expected string
	}{
		{"%", 10, 2, "Помилка: Невірна операція\n"},
		{"/", 10, 0, "помилка: Ділення на нуль неможливе\n"},
	}

	// Test cases for valid input
	for _, tc := range testCases {
		var buf bytes.Buffer
		fmt.SetOutput(&buf)
		operate(tc.operator, tc.num1, tc.num2)

		if buf.String() != tc.expected {
			t.Errorf("Expected output: %q but got: %q", tc.expected, buf.String())
		}
	}

	// Test cases for invalid input
	for _, tc := range invalidTestCases {
		var buf bytes.Buffer
		fmt.SetOutput(&buf)

		operate(tc.operator, tc.num1, tc.num2)

		if buf.String() != tc.expected {
			t.Errorf("Expected output: %q but got: %q", tc.expected, buf.String())
		}
	}
}
