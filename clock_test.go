package main

import (
	"testing"
)

var testHours = []struct {
	given24HrsH    int
	expected12HrsH int
	cycle          cycle // AM or PM
	name           string
}{
	{
		given24HrsH:    0,
		expected12HrsH: 12,
		cycle:          am,
		name:           "12_in_the_morning",
	},
	{
		given24HrsH:    1,
		expected12HrsH: 1,
		cycle:          am,
		name:           "1_AM_in_the_morning",
	},
	{
		given24HrsH:    11,
		expected12HrsH: 11,
		cycle:          am,
		name:           "11_AM_in_the_morning",
	},
	{
		given24HrsH:    12,
		expected12HrsH: 12,
		cycle:          pm,
		name:           "12_PM_in_the_afternoon",
	},
	{
		given24HrsH:    13,
		expected12HrsH: 1,
		cycle:          pm,
		name:           "1_PM_in_the_afternoon",
	},
	{
		given24HrsH:    23,
		expected12HrsH: 11,
		cycle:          pm,
		name:           "11_PM_in_the_night",
	},
}

func TestNormalize12Hour(t *testing.T) {
	for _, tc := range testHours {
		t.Run(tc.name, func(t *testing.T) {
			actual12H, cycle, _ := normalizeTo12Hour(tc.given24HrsH)
			if tc.expected12HrsH != actual12H || tc.cycle != cycle {
				t.Errorf("expected: %d, cycle: %s actual: %d, cycle: %s", tc.expected12HrsH, tc.cycle, actual12H, cycle)
			}
		})
	}
}
