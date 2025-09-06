package math_test

import (
	"testing"

	"github.com/fatjyc/hello-go/pkg/math"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "positive numbers",
			a:        1,
			b:        2,
			expected: 3,
		},
		{
			name:     "negative numbers",
			a:        -1,
			b:        -2,
			expected: -3,
		},
		{
			name:     "zero values",
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			name:     "positive and negative",
			a:        5,
			b:        -3,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := math.Sum(tt.a, tt.b)
			assert.Equal(t, tt.expected, result)
		})
	}
}
