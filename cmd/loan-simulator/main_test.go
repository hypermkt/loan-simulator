package main

import (
	"testing"
)

func TestCalcMonths(t *testing.T) {
	patterns := []struct {
		y        int
		expected int
	}{
		{1, 12},
		{2, 24},
	}

	for idx, pattern := range patterns {
		actual := calcMonths(pattern.y)
		if actual != pattern.expected {
			t.Errorf("pattern %d: want %d, actual %d", idx, pattern.expected, actual)
		}
	}
}
