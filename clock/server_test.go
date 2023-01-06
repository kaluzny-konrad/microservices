package main

import (
	"testing"
)

type clockTests struct {
	input    int
	expected string
}

var addClockTests = []clockTests{
	{0, "00"},
	{1, "01"},
	{9, "09"},
	{10, "10"},
	{11, "11"},
	{12, "12"},
}

func TestGetClockVariable(t *testing.T) {
	for _, test := range addClockTests {
		output := getClockVariable(test.input)
		if output != test.expected {
			t.Errorf("Output %s not equal to expected %s", output, test.expected)
		}
	}
}
