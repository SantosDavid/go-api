package shared_test

import (
	"testing"

	"github.com/santosdavid/go-api-v2/domain/shared"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	tests := []struct {
		elements []int
		needed   int
		expected bool
	}{
		{
			elements: []int{1, 2, 3, 4},
			needed:   8,
			expected: false,
		},
		{
			elements: []int{1, 2, 3, 4},
			needed:   4,
			expected: true,
		},
		{
			elements: []int{},
			needed:   0,
			expected: false,
		},
	}

	for _, tt := range tests {
		result := shared.Contains(tt.elements, tt.needed)

		assert.Equal(t, tt.expected, result)
	}
}
