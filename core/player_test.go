package core

import (
	"testing"
)

func TestDice(t *testing.T) {

	testCases := []struct {
		description string
		p           player
		check       func(int) bool
	}{{
		description: "rolling dice should return value between 1-6",
		p:           player{},
		check: func(value int) bool {
			return value >= 1 && value <= 6
		},
	}}

	for _, test := range testCases {

		value := test.p.rollDice()

		if !test.check(value) {
			t.Fatal("dice roll test case failed")
		}
	}
}
