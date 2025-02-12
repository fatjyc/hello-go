package tools_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/fatjyc/hello-go/tools"
)

func TestIsOdd(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, false},
		{0, false},
	}

	for _, test := range tests {
		if result := tools.IsOdd(test.input); result != test.expected {
			assert.Equal(t, test.expected, result, "IsOdd(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{1, false},
		{2, true},
		{3, false},
		{4, true},
		{-2, true},
		{0, true},
	}

	for _, test := range tests {
		if result := tools.IsEven(test.input); result != test.expected {
			assert.Equal(t, test.expected, result, "IsEven(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{9, false},
		{13, true},
		{0, false},
		{-3, false},
	}

	for _, test := range tests {
		if result := tools.IsPrime(test.input); result != test.expected {
			assert.Equal(t, test.expected, result, "IsPrime(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}
