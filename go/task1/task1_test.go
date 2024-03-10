package main

import "testing"

type ConvToIntTest struct {
	arg1  string
	teamX int
	teamY int
}

var AddTests = []ConvToIntTest{
	{"1:0", 1, 0},
	{"3:2", 3, 2},
}

func TestConvToInt(t *testing.T) {

	for _, test := range AddTests {
		teamX, teamY, err := convToInt(test.arg1)

		if err != nil {
			t.Errorf("Error to convert number")
		}

		if teamX != test.teamX || teamY != test.teamY {
			t.Errorf("Output %q %q not equal to expected %q %q", teamX, teamY, test.teamX, test.teamY)
		}
	}
}
