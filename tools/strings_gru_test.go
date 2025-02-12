package tools

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "multiple characters",
			input:    "hello",
			expected: "olleh",
		},
		{
			name:     "with spaces",
			input:    "hello world",
			expected: "dlrow olleh",
		},
		{
			name:     "with special characters",
			input:    "hello!@#$%",
			expected: "%$#@!olleh",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "single character",
			input:    "a",
			expected: 1,
		},
		{
			name:     "multiple characters",
			input:    "hello",
			expected: 5,
		},
		{
			name:     "with spaces",
			input:    "hello world",
			expected: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Length(tt.input)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "single character",
			input:    "a",
			expected: true,
		},
		{
			name:     "two same characters",
			input:    "aa",
			expected: true,
		},
		{
			name:     "palindrome odd length",
			input:    "radar",
			expected: true,
		},
		{
			name:     "palindrome even length",
			input:    "noon",
			expected: true,
		},
		{
			name:     "not palindrome",
			input:    "hello",
			expected: false,
		},
		{
			name:     "not palindrome with spaces",
			input:    "race a car",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "empty strings",
			s1:       "",
			s2:       "",
			expected: true,
		},
		{
			name:     "single character same",
			s1:       "a",
			s2:       "a",
			expected: true,
		},
		{
			name:     "single character different",
			s1:       "a",
			s2:       "b",
			expected: false,
		},
		{
			name:     "different lengths",
			s1:       "rat",
			s2:       "car",
			expected: false,
		},
		{
			name:     "valid anagram",
			s1:       "listen",
			s2:       "silent",
			expected: true,
		},
		{
			name:     "not anagram same length",
			s1:       "hello",
			s2:       "world",
			expected: false,
		},
		{
			name:     "with spaces",
			s1:       "debit card",
			s2:       "bad credit",
			expected: true,
		},
		{
			name:     "with special characters",
			s1:       "a!b@c#",
			s2:       "c@b#a!",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsAnagram(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
