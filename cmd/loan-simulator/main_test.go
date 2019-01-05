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

func TestCalcInterest(t *testing.T) {
	patterns := []struct {
		c        int
		mir      float64
		expected int
	}{
		{30000000, 0.000833333333333334, 25000},
	}

	for idx, pattern := range patterns {
		p := Params{
			CurrentBalance:      pattern.c,
			MonthlyInterestRate: pattern.mir,
		}
		actual := calcInterest(p)
		if actual != pattern.expected {
			t.Errorf("pattern %d: want %d, actual %d", idx, pattern.expected, actual)
		}
	}
}
