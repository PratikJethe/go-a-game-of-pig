package cmd

import (
	"log"
	"os"
	"strconv"
	"strings"
	"github.com/pratikjethe/pig/core"
)

// Initiate function is responsible to parse input arguments and call StartSimulation function
// It performs basic validation checks

func Initiate() {
	if len(os.Args) < 3 {
		panic("Invalid Input")
	}

	p1LowerHoldLimit, p1UpperHoldLimit := getUpperAndLowerHoldLimit(os.Args[1])
	p2LowerHoldLimit, p2UpperHoldLimit := getUpperAndLowerHoldLimit(os.Args[2])

	core.StartSimulation(p1UpperHoldLimit, p1LowerHoldLimit, p2UpperHoldLimit, p2LowerHoldLimit)

}
// getUpperAndLowerHoldLimit is responsible to parse cmd line arguments and return upper and lower limit for an strategy
// incase of single srategy it returns same value for both 
func getUpperAndLowerHoldLimit(input string) (int, int) {

	if strings.Contains(input, "-") {

		splittedInput := strings.Split(input, "-")

		lowerHoldLimit, err := strconv.Atoi(splittedInput[0])
		if err != nil {
			log.Fatal(err.Error())
		}
		upperHoldLimit, err := strconv.Atoi(splittedInput[1])
		if err != nil {
			log.Fatal(err.Error())
		}
		if upperHoldLimit < lowerHoldLimit {
			log.Fatal("Upper range should be greater than lower range")
		}
		return lowerHoldLimit, upperHoldLimit
	} else {
		holdLimit, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err.Error())
		}
		return holdLimit, holdLimit

	}
}
