package kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		name     string
		input1   *point
		input2 	 *point
		expected int
	}{
		{
			name:     "Simple Positive Case",
			input1:    newPoint(1,1),
			input2:    newPoint(2, 2),
			expected: 2,
		},
		{
			name:     "Negative Coordinates",
			input1:   newPoint(-1, 14),
			input2:   newPoint(10, -23),
			expected: 48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ManhattanDistance(tt.input1, tt.input2))
		})
	}
}
