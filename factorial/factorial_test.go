package Factorial

import (
	"math/big"
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
	testCases := []struct {
		input    int
		expected *big.Int
	}{
		{0, big.NewInt(1)},
		{1, big.NewInt(1)},
		{5, big.NewInt(120)},
		{10, big.NewInt(3628800)},
	}

	for _, testCase := range testCases {
		result := factorial(testCase.input)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("factorial(%d) = %v, expected %v", testCase.input, result, testCase.expected)
		}
	}
}

func TestCalculateFactorials(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected map[string]*big.Int
	}{
		{0, 1, map[string]*big.Int{"a_factorial": big.NewInt(1), "b_factorial": big.NewInt(1)}},
		{3, 4, map[string]*big.Int{"a_factorial": big.NewInt(6), "b_factorial": big.NewInt(24)}},
		{5, 5, map[string]*big.Int{"a_factorial": big.NewInt(120), "b_factorial": big.NewInt(120)}},
		// Add more test cases as needed
	}

	for _, testCase := range testCases {
		result := CalculateFactorials(testCase.a, testCase.b)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("CalculateFactorials(%d, %d) = %v, expected %v", testCase.a, testCase.b, result, testCase.expected)
		}
	}
}
