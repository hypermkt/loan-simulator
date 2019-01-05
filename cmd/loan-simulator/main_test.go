package main

import (
	"testing"
)

func TestCalcRepaidAmount(t *testing.T) {
	patterns := []struct {
		p        Params
		expected int
	}{
		{Params{
			Year:                35,
			Months:              420,
			InterestRate:        1,
			MonthlyInterestRate: 0.0008333333333333334,
			AmountMan:           3000,
			Amount:              30000000,
			CurrentBalance:      30000000,
		}, 84685},
	}

	for idx, pattern := range patterns {
		actual := calcRepaidAmount(pattern.p)
		if actual != pattern.expected {
			t.Errorf("pattern %d: want %d, actual %d", idx, pattern.expected, actual)
		}
	}
}

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

func TestToParams(t *testing.T) {
	patterns := []struct {
		y        int
		ir       float64
		a        int
		expected Params
	}{
		{35, 1, 3500, Params{
			Year:                35,
			Months:              420,
			InterestRate:        1,
			MonthlyInterestRate: 0.0008333333333333334,
			AmountMan:           3500,
			Amount:              35000000,
			CurrentBalance:      35000000,
		}},
	}

	for idx, pattern := range patterns {
		actual := toParams(pattern.y, pattern.ir, pattern.a)
		if actual != pattern.expected {
			t.Errorf("pattern %d: want %v, actual %v", idx, pattern.expected, actual)
		}
	}
}
